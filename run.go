package main

import (
	"encoding/hex"
	"sort"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	log "github.com/sirupsen/logrus"
)

type Run struct {
	StartedAt time.Time
	EndedAt   time.Time
	LocalID   peer.ID
	Spans     []Span
	Involved  map[peer.ID]bool
}

func (r *Run) Data(content *Content) RunData {
	peerInfos := r.GatherPeerInfos(content)
	return RunData{
		StartedAt: r.StartedAt,
		EndedAt:   r.EndedAt,
		LocalID:   r.LocalID,
		Distance:  hex.EncodeToString(content.DistanceTo(r.LocalID)),
		Spans:     r.SpanData(),
		PeerInfos: peerInfos,
		PeerOrder: r.PeerOrder(peerInfos),
	}
}

func (r *Run) SpanData() []SpanData {
	spanData := make([]SpanData, len(r.Spans))
	for i, span := range r.Spans {
		spanData[i] = SpanData{
			RelStart:  span.StartedAt().Sub(r.StartedAt).Seconds(),
			DurationS: span.Duration().Seconds(),
			Start:     span.StartedAt(),
			End:       span.StartedAt().Add(span.Duration()),
			PeerID:    span.RemoteID(),
			Operation: span.Operation(),
			Type:      span.Type(),
		}
		if span.Error() != nil {
			spanData[i].Error = span.Error().Error()
		}
	}
	return spanData
}

func (r *Run) GatherPeerInfos(content *Content) map[string]PeerInfo {
	peerInfos := map[string]PeerInfo{}

	// Create a span slice that's sorted by the ending time stamp.
	// Doing it this way we can easily find the peers that have discovered
	// other peers first as only the final timestamp is determinative.
	endSortedSpans := make([]Span, len(r.Spans))
	for i, span := range r.Spans {
		endSortedSpans[i] = span
	}

	sort.SliceStable(endSortedSpans, func(i, j int) bool {
		return endSortedSpans[i].EndedAt().Before(endSortedSpans[j].EndedAt())
	})

	for peerID := range r.Involved {
		pi := PeerInfo{
			ID:           peerID,
			Distance:     hex.EncodeToString(content.DistanceTo(peerID)),
			IsBootstrap:  r.IsBootstrapPeer(peerID),
			AgentVersion: r.AgentVersion(peerID),
			firstSpan:    r.FirstSpan(peerID),
		}

		for _, span := range endSortedSpans {
			srs, ok := span.(*SendRequestSpan)
			if span.RemoteID() == peerID || !ok || !srs.ReturnedCloserPeer(peerID) || span.StartedAt().After(pi.firstSpan.StartedAt()) {
				continue
			}

			pi.RelDiscoveredAt = srs.StartedAt().Add(srs.Duration()).Sub(r.StartedAt).Seconds()
			pi.DiscoveredAt = srs.StartedAt().Add(srs.Duration())
			pi.DiscoveredFrom = srs.RemoteID()
			break
		}

		peerInfos[peerID.Pretty()] = pi
	}

	return peerInfos
}

// IsBootstrapPeer checks if there is a dial in the spans before another span type
func (r *Run) IsBootstrapPeer(peerID peer.ID) bool {
	for _, span := range r.Spans {
		if span.RemoteID() != peerID {
			continue
		}

		_, ok := span.(*DialSpan)
		if ok {
			return false
		} else {
			return !span.StartedAt().After(r.StartedAt.Add(100 * time.Millisecond))
		}
	}

	log.WithField("peerID", FmtPeerID(peerID)).Warn("unexpected peer")
	return false
}

func (r *Run) AgentVersion(peerID peer.ID) string {
	for _, span := range r.Spans {
		if span.RemoteID() != peerID {
			continue
		}

		if srs, ok := span.(*SendRequestSpan); ok && srs.AgentVersion != "" {
			return srs.AgentVersion
		}

		if sms, ok := span.(*SendMessageSpan); ok && sms.AgentVersion != "" {
			return sms.AgentVersion
		}
	}

	return ""
}

func (r *Run) FirstSpan(peerID peer.ID) Span {
	for _, span := range r.Spans {
		if span.RemoteID() != peerID {
			continue
		}
		return span
	}

	panic("unexpected")
}

func (r *Run) PeerOrder(peerInfos map[string]PeerInfo) []peer.ID {
	var bootstrapPeers []PeerInfo
	var discoveredPeers []PeerInfo
	for _, pi := range peerInfos {
		if pi.IsBootstrap {
			bootstrapPeers = append(bootstrapPeers, pi)
		} else {
			discoveredPeers = append(discoveredPeers, pi)
		}
	}

	sort.SliceStable(bootstrapPeers, func(i, j int) bool {
		return bootstrapPeers[i].firstSpan.Duration() >= bootstrapPeers[j].firstSpan.Duration()
	})

	sort.SliceStable(discoveredPeers, func(i, j int) bool {
		return discoveredPeers[i].firstSpan.StartedAt().Before(discoveredPeers[j].firstSpan.StartedAt())
	})

	var peerOrder []peer.ID
	for _, pi := range append(bootstrapPeers, discoveredPeers...) {
		peerOrder = append(peerOrder, pi.ID)
	}

	return peerOrder
}
