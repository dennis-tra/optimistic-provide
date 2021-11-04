package main

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
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
