package services

import (
	"context"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/routing"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	ma "github.com/multiformats/go-multiaddr"
)

type ProvideState struct {
	h *dht.Host

	dialsLk sync.RWMutex
	dials   []*DialSpan

	connectionsStartedLk sync.RWMutex
	connectionsStarted   map[peer.ID]time.Time

	connectionsLk sync.RWMutex
	connections   []*ConnectionSpan
}

func (ps *ProvideState) Register(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	ctx, lookupEvents := kaddht.RegisterForLookupEvents(ctx)
	ctx, queryEvents := routing.RegisterForQueryEvents(ctx)

	go ps.consumeLookupEvents(lookupEvents)
	go ps.consumeQueryEvents(queryEvents)

	for _, notifier := range ps.h.Transports {
		notifier.Notify(ps)
	}

	ps.h.Host.Network().Notify(ps)

	return ctx, cancel
}

func (ps *ProvideState) Unregister(cancel context.CancelFunc) {
	ps.h.Host.Network().StopNotify(ps)

	for _, notifier := range ps.h.Transports {
		notifier.StopNotify(ps)
	}

	cancel()
}

func (ps *ProvideState) consumeQueryEvents(queryEvents <-chan *routing.QueryEvent) {
	for event := range queryEvents {
		switch event.Type {
		case routing.DialingPeer:
			log.Infow("DialingPeer", "peerID", event.ID.Pretty())
			ps.connectionsStartedLk.Lock()
			ps.connectionsStarted[event.ID] = time.Now()
			ps.connectionsStartedLk.Unlock()
		case routing.SendingQuery:
			log.Infow("SendingQuery", "peerID", event.ID.Pretty())
		case routing.PeerResponse:
			log.Infow("PeerResponse", "peerID", event.ID.Pretty())
		case routing.QueryError:
			log.Infow("QueryError", "peerID", event.ID.Pretty())
		}
	}
}

func (ps *ProvideState) consumeLookupEvents(lookupEvents <-chan *kaddht.LookupEvent) {
	for event := range lookupEvents {
		_ = event
	}
}

func (ps *ProvideState) Listen(network network.Network, multiaddr ma.Multiaddr) {
}

func (ps *ProvideState) ListenClose(network network.Network, multiaddr ma.Multiaddr) {
}

func (ps *ProvideState) Connected(network network.Network, conn network.Conn) {
	end := time.Now()
	ps.connectionsStartedLk.Lock()
	defer ps.connectionsStartedLk.Unlock()
	started, ok := ps.connectionsStarted[conn.RemotePeer()]
	if !ok {
		return
	}
	delete(ps.connectionsStarted, conn.RemotePeer())

	ps.connectionsLk.Lock()
	defer ps.connectionsLk.Unlock()
	ps.connections = append(ps.connections, &ConnectionSpan{
		RemotePeerID: conn.RemotePeer(),
		Maddr:        conn.RemoteMultiaddr(),
		Start:        started,
		End:          end,
	})
}

func (ps *ProvideState) Disconnected(network network.Network, conn network.Conn) {
}

func (ps *ProvideState) OpenedStream(network network.Network, stream network.Stream) {
}

func (ps *ProvideState) ClosedStream(network network.Network, stream network.Stream) {
}

func (ps *ProvideState) DialStarted(trpt string, raddr ma.Multiaddr, p peer.ID, start time.Time) {
}

func (ps *ProvideState) DialEnded(trpt string, raddr ma.Multiaddr, p peer.ID, start time.Time, end time.Time, err error) {
	d := &DialSpan{
		RemotePeerID: p,
		Maddr:        raddr,
		Start:        start,
		End:          end,
		Trpt:         trpt,
		Error:        err,
	}
	ps.dialsLk.Lock()
	defer ps.dialsLk.Unlock()

	ps.dials = append(ps.dials, d)
}
