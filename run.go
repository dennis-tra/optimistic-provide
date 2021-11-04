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
	Involved  map[peer.ID]struct{}
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

	for peerID := range r.Involved {
		pi := PeerInfo{
			ID:           peerID,
			Distance:     hex.EncodeToString(content.DistanceTo(peerID)),
			IsBootstrap:  r.IsBootstrapPeer(peerID),
			AgentVersion: r.AgentVersion(peerID),
		}

		for _, span := range r.Spans {
			srs, ok := span.(*SendRequestSpan)
			if span.RemoteID() == peerID || !ok || !srs.ReturnedCloserPeer(peerID) {
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
		return !ok
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

func (r *Run) EarliestSendRequestEnd(peerID peer.ID) *time.Time {
	var earliest *time.Time

	for _, span := range r.Spans {
		srs, ok := span.(*SendRequestSpan)
		if span.RemoteID() != peerID || span.Error() != nil || !ok {
			continue
		}

		if earliest == nil || earliest.After(srs.End) {
			earliest = &srs.End
		}
	}
	return earliest
}

func (r *Run) EarliestSendRequestStart(peerID peer.ID) *time.Time {
	var earliest *time.Time

	for _, span := range r.Spans {
		srs, ok := span.(*SendRequestSpan)
		if span.RemoteID() != peerID || span.Error() != nil || !ok {
			continue
		}

		if earliest == nil || earliest.After(srs.Start) {
			earliest = &srs.Start
		}
	}
	return earliest
}

func (r *Run) EarliestDialStart(peerID peer.ID) *time.Time {
	var earliest *time.Time

	for _, span := range r.Spans {
		srs, ok := span.(*DialSpan)
		if span.RemoteID() != peerID || span.Error() != nil || !ok {
			continue
		}

		if earliest == nil || earliest.After(srs.Start) {
			earliest = &srs.Start
		}
	}
	return earliest
}

func (r *Run) PeerOrder(peerInfos map[string]PeerInfo) []peer.ID {
	var bootstrapPeers []PeerInfo
	var discoveredPeers []PeerInfo
	for _, pi := range peerInfos {
		if pi.DiscoveredFrom == "" {
			bootstrapPeers = append(bootstrapPeers, pi)
		} else {
			discoveredPeers = append(discoveredPeers, pi)
		}
	}

	sort.SliceStable(bootstrapPeers, func(i, j int) bool {
		bpi := r.EarliestSendRequestEnd(bootstrapPeers[i].ID)
		bpj := r.EarliestSendRequestEnd(bootstrapPeers[j].ID)
		if bpi == nil && bpj == nil {
			return true
		} else if bpi == nil && bpj != nil {
			return false
		} else if bpi != nil && bpj == nil {
			return true
		} else {
			return bpi.Before(*bpj)
		}
	})

	sort.SliceStable(discoveredPeers, func(i, j int) bool {
		if discoveredPeers[i].DiscoveredAt == discoveredPeers[j].DiscoveredAt {
			pi := r.EarliestDialStart(discoveredPeers[i].ID)
			if pi == nil {
				pi = r.EarliestSendRequestStart(discoveredPeers[i].ID)
			}

			pj := r.EarliestDialStart(discoveredPeers[j].ID)
			if pj == nil {
				pj = r.EarliestSendRequestStart(discoveredPeers[j].ID)
			}

			if pi == nil && pj == nil {
				return false
			} else if pi == nil && pj != nil {
				return false
			} else if pi != nil && pj == nil {
				return true
			} else {
				return pi.Before(*pj)
			}
		}
		return discoveredPeers[i].DiscoveredAt.After(discoveredPeers[j].DiscoveredAt)
	})

	var peerOrder []peer.ID
	for _, pi := range append(bootstrapPeers, discoveredPeers...) {
		peerOrder = append(peerOrder, pi.ID)
	}

	return peerOrder
}
