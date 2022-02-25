package wrap

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
)

// RPCEventBufferSize defines the number of events to buffer.
var RPCEventBufferSize = 16

// RPCSendRequestStartedEvent .
type RPCSendRequestStartedEvent struct {
	QueryID    uuid.UUID
	RemotePeer peer.ID
	StartedAt  time.Time
	Request    *pb.Message
}

// RPCSendRequestEndedEvent .
type RPCSendRequestEndedEvent struct {
	QueryID    uuid.UUID
	RemotePeer peer.ID
	StartedAt  time.Time
	EndedAt    time.Time
	Request    *pb.Message
	Response   *pb.Message
	Error      error
}

// RPCSendMessageStartedEvent .
type RPCSendMessageStartedEvent struct {
	QueryID    uuid.UUID
	RemotePeer peer.ID
	StartedAt  time.Time
	Message    *pb.Message
}

// RPCSendMessageEndedEvent .
type RPCSendMessageEndedEvent struct {
	QueryID    uuid.UUID
	RemotePeer peer.ID
	StartedAt  time.Time
	EndedAt    time.Time
	Message    *pb.Message
	Error      error
}

type (
	rpcEventKey  struct{}
	eventChannel struct {
		mu  sync.Mutex
		ctx context.Context
		ch  chan<- interface{}
	}
)

// waitThenClose is spawned in a goroutine when the channel is registered. This
// safely cleans up the channel when the context has been canceled.
func (e *eventChannel) waitThenClose() {
	<-e.ctx.Done()
	e.mu.Lock()
	close(e.ch)
	// 1. Signals that we're done.
	// 2. Frees memory (in case we end up hanging on to this for a while).
	e.ch = nil
	e.mu.Unlock()
}

// send sends an event on the event channel, aborting if either the passed or
// the internal context expire.
func (e *eventChannel) send(ctx context.Context, ev interface{}) {
	e.mu.Lock()
	// Closed.
	if e.ch == nil {
		e.mu.Unlock()
		return
	}
	// in case the passed context is unrelated, wait on both.
	select {
	case e.ch <- ev:
	case <-e.ctx.Done():
	case <-ctx.Done():
	}
	e.mu.Unlock()
}

// RegisterForRPCEvents registers a query event channel with the given
// context. The returned context can be passed to DHT queries to receive query
// events on the returned channels.
//
// The passed context MUST be canceled when the caller is no longer interested
// in query events.
func RegisterForRPCEvents(ctx context.Context) (context.Context, <-chan interface{}) {
	ch := make(chan interface{}, RPCEventBufferSize)
	ech := &eventChannel{ch: ch, ctx: ctx}
	go ech.waitThenClose()
	return context.WithValue(ctx, rpcEventKey{}, ech), ch
}

// PublishRPCEvent publishes a query event to the query event channel
// associated with the given context, if any.
func PublishRPCEvent(ctx context.Context, ev interface{}) {
	ich := ctx.Value(rpcEventKey{})
	if ich == nil {
		return
	}

	// We *want* to panic here.
	ech := ich.(*eventChannel)
	ech.send(ctx, ev)
}

// SubscribesToRPCEvents returns true if the context subscribes to query
// events. If this function returns falls, calling `PublishRPCEvent` on the
// context will be a no-op.
func SubscribesToRPCEvents(ctx context.Context) bool {
	return ctx.Value(rpcEventKey{}) != nil
}
