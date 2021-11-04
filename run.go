package main

import (
	"encoding/hex"
	"time"

	u "github.com/ipfs/go-ipfs-util"
	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	log "github.com/sirupsen/logrus"
)

type Run struct {
	StartedAt time.Time
	EndedAt   time.Time
	LocalID   peer.ID
	Events    []Event
}

type Span struct {
	RelStart  float64
	DurationS float64
	Start     time.Time
	End       time.Time
	PeerID    peer.ID
	Type      SpanType
	Error     string
	Details   string
}

func (r *Run) Data(content *Content) RunData {
	peerInfos := r.GatherPeerInfos(content)
	return RunData{
		StartedAt: r.StartedAt,
		EndedAt:   r.EndedAt,
		LocalID:   r.LocalID,
		Distance:  hex.EncodeToString(u.XOR(kbucket.ConvertPeerID(r.LocalID), kbucket.ConvertKey(string(content.mhash)))),
		Spans:     r.DetectSpans(),
		PeerInfos: peerInfos,
		PeerOrder: r.PeerOrder(peerInfos),
	}
}

func (r *Run) DetectSpans() []Span {
	type SpanState struct {
		Start   time.Time
		Details string
		Count   int
	}
	var spans []Span
	spanStates := map[peer.ID]map[SpanType]*SpanState{}

	for _, event := range r.Events {
		if event.LocalID() != r.LocalID {
			panic("unexpected peer id")
		}

		// Check if we are already tracking this peer - if not create a map for it.
		if _, found := spanStates[event.RemoteID()]; !found {
			spanStates[event.RemoteID()] = map[SpanType]*SpanState{}
		}

		spanState, found := spanStates[event.RemoteID()][event.Span()]

		// Check if it's the start of a span or not
		if event.IsStart() {
			// If we don't have an "open" start event in our state we create the SpanState
			// and set the count to 1 (number of start events)
			// If we already have an open start event in our state we just increment the
			// counter to keep track how many open events we came across
			if !found {
				spanState = &SpanState{
					Start:   event.TimeStamp(),
					Details: event.Details(),
					Count:   1,
				}
			} else {
				spanState.Count += 1
			}

			spanStates[event.RemoteID()][event.Span()] = spanState
		} else {
			// If we received an end event while there is no open span state we just do nothing
			if !found {
				continue
			}

			// Decrement the count
			spanState.Count -= 1

			// If the end event contains an error and this was not the last event for this
			// open span we do nothing and wait for the last or a successful end event
			if event.Error() != nil && spanState.Count != 0 {
				spanStates[event.RemoteID()][event.Span()] = spanState
				continue
			}

			if event.Details() != "" && spanState.Details != event.Details() {
				log.WithFields(log.Fields{
					"start": spanStates[event.RemoteID()][event.Span()].Details,
					"end":   event.Details(),
				}).Warnln("Unexpected details combination")
			}

			// Create span
			span := Span{
				RelStart:  spanState.Start.Sub(r.StartedAt).Seconds(),
				DurationS: event.TimeStamp().Sub(spanState.Start).Seconds(),
				Start:     spanState.Start,
				End:       event.TimeStamp(),
				PeerID:    event.RemoteID(),
				Type:      event.Span(),
				Details:   spanState.Details,
			}

			if event.Error() != nil {
				span.Error = event.Error().Error()
			}

			spans = append(spans, span)

			// Delete span state so that subsequent events of this span type can be tracked again.
			delete(spanStates[event.RemoteID()], event.Span())
		}
	}
	return spans
}

func (r *Run) GatherPeerInfos(content *Content) map[string]PeerInfo {
	peerInfos := map[string]PeerInfo{}

	for _, outerEvt := range r.Events {
		if _, exists := peerInfos[outerEvt.RemoteID().Pretty()]; exists {
			continue
		}

		peerID := outerEvt.RemoteID()
		pi := PeerInfo{
			ID:       peerID,
			Distance: hex.EncodeToString(u.XOR(kbucket.ConvertPeerID(peerID), kbucket.ConvertKey(string(content.cid.Hash())))),
		}

		for _, event := range r.Events {
			switch evt := event.(type) {
			case *SendRequestStart:
				if event.RemoteID() == peerID && evt.AgentVersion != "" {
					pi.AgentVersion = evt.AgentVersion
				}
			case *SendRequestEnd:
				if event.RemoteID() == peerID && evt.AgentVersion != "" {
					pi.AgentVersion = evt.AgentVersion
				}
				// Track which peer discovered this peer
				// If there is no tracked discovery (pi.RelDiscoveredAt <= 0) AND
				// the event originated from another peer (event.RemoteID() != peerID) AND
				// there was no error (evt.Error() == nil) AND
				// It was a find node request (evt.Response.Type == pb.Message_FIND_NODE)
				// Then loop through the returned peers and see if it is part of that.
				if pi.RelDiscoveredAt <= 0 && event.RemoteID() != peerID && evt.Error() == nil && (evt.Response.Type == pb.Message_FIND_NODE || evt.Response.Type == pb.Message_GET_PROVIDERS) {
					for _, closer := range pb.PBPeersToPeerInfos(evt.Response.CloserPeers) {
						if closer.ID == peerID {
							pi.DiscoveredAt = evt.Time
							pi.DiscoveredFrom = evt.RemoteID()
							pi.RelDiscoveredAt = evt.Time.Sub(r.StartedAt).Seconds()
						}
					}
				}
			case *SendMessageStart:
				if event.RemoteID() == peerID && evt.AgentVersion != "" {
					pi.AgentVersion = evt.AgentVersion
				}
			case *SendMessageEnd:
				if event.RemoteID() == peerID && evt.AgentVersion != "" {
					pi.AgentVersion = evt.AgentVersion
				}
			}
		}

		peerInfos[peerID.Pretty()] = pi

	}
	return peerInfos
}

func (r *Run) PeerOrder(peerInfos map[string]PeerInfo) []peer.ID {
	// The below code is ugly
	// This should sort the peers by their time they were discovered OR
	// if they are bootstrap peers by their sendRequest finish time.
	type sortInfo struct {
		id           peer.ID
		discoveredAt *time.Time
		sendReqEnd   *time.Time
	}

	var sortInfos []sortInfo
OUTER:
	for _, peerInfo := range peerInfos {
		pi := peerInfo
		if pi.DiscoveredFrom.String() != "" {
			for i, p := range sortInfos {
				if p.discoveredAt == nil {
					continue
				}
				if p.discoveredAt.After(pi.DiscoveredAt) {
					sortInfos = append(sortInfos[:i+1], sortInfos[i:]...)
					sortInfos[i] = sortInfo{
						id:           pi.ID,
						discoveredAt: &pi.DiscoveredAt,
					}
					continue OUTER
				}
			}
			sortInfos = append(sortInfos,
				sortInfo{
					id:           pi.ID,
					discoveredAt: &pi.DiscoveredAt,
				})
		} else {
			for _, event := range r.Events {
				_, ok := event.(*SendRequestEnd)
				if event.RemoteID() != pi.ID || event.Error() != nil || !ok {
					continue
				}

				for i, p := range sortInfos {
					if p.sendReqEnd == nil {
						continue
					}

					if (p.sendReqEnd.After(event.TimeStamp()) && event.TimeStamp().After(r.StartedAt)) || p.discoveredAt != nil {
						sortInfos = append(sortInfos[:i+1], sortInfos[i:]...)
						sortInfos[i] = sortInfo{
							id:           pi.ID,
							discoveredAt: &pi.DiscoveredAt,
						}
						continue OUTER
					}
				}
				sortInfos = append([]sortInfo{{
					id:           pi.ID,
					discoveredAt: &pi.DiscoveredAt,
				}}, sortInfos...)
				continue OUTER
			}
		}
	}

	var peerOrder []peer.ID
	for _, si := range sortInfos {
		peerOrder = append(peerOrder, si.id)
	}

	return peerOrder
}
