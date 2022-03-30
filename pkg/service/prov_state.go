package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	"github.com/libp2p/go-libp2p-kad-dht/qpeerset"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
	"github.com/dennis-tra/optimistic-provide/pkg/wrap"
)

type ProvideState struct {
	h       *dht.Host
	content *util.Content
	cancel  context.CancelFunc

	dialsLk sync.RWMutex
	dials   []*DialSpan

	findNodesLk sync.RWMutex
	findNodes   []*FindNodesSpan

	addProvidersLk sync.RWMutex
	addProviders   []*AddProvidersSpan

	connectionsStartedLk sync.RWMutex
	connectionsStarted   map[peer.ID]time.Time

	connectionsLk sync.RWMutex
	connections   []*ConnectionSpan

	peerSet map[uuid.UUID]*qpeerset.QueryPeerset

	relevantPeers sync.Map
}

func NewProvideState(h *dht.Host, content *util.Content) *ProvideState {
	return &ProvideState{
		h:                    h,
		content:              content,
		dialsLk:              sync.RWMutex{},
		dials:                []*DialSpan{},
		findNodesLk:          sync.RWMutex{},
		findNodes:            []*FindNodesSpan{},
		addProvidersLk:       sync.RWMutex{},
		addProviders:         []*AddProvidersSpan{},
		connectionsStartedLk: sync.RWMutex{},
		connectionsStarted:   map[peer.ID]time.Time{},
		connectionsLk:        sync.RWMutex{},
		connections:          []*ConnectionSpan{},
		peerSet:              map[uuid.UUID]*qpeerset.QueryPeerset{},
		relevantPeers:        sync.Map{},
	}
}

func (ps *ProvideState) Register(ctx context.Context) context.Context {
	if ps.cancel != nil {
		panic("already registered for events")
	}

	ctx, cancel := context.WithCancel(ctx)
	ctx, lookupEvents := kaddht.RegisterForLookupEvents(ctx)
	ctx, queryEvents := routing.RegisterForQueryEvents(ctx)
	ctx, rpcEvents := wrap.RegisterForRPCEvents(ctx)

	go ps.consumeLookupEvents(lookupEvents)
	go ps.consumeQueryEvents(queryEvents)
	go ps.consumeRPCEvents(rpcEvents)

	for _, notifier := range ps.h.Transports {
		notifier.Notify(ps)
	}

	ps.h.Host.Network().Notify(ps)

	ps.cancel = cancel

	return ctx
}

func (ps *ProvideState) Unregister() {
	ps.h.Host.Network().StopNotify(ps)

	for _, notifier := range ps.h.Transports {
		notifier.StopNotify(ps)
	}

	ps.cancel()
	ps.cancel = nil

	ps.filterDials()
	ps.filterConnections()
}

func (ps *ProvideState) consumeLookupEvents(lookupEvents <-chan *kaddht.LookupEvent) {
	for event := range lookupEvents {
		ps.ensurePeerset(event)
		if event.Response != nil {
			ps.trackLookupResponse(event)
		} else if event.Request != nil {
			ps.trackLookupRequest(event)
		}
	}
}

func (ps *ProvideState) consumeQueryEvents(queryEvents <-chan *routing.QueryEvent) {
	for event := range queryEvents {
		switch event.Type {
		case routing.DialingPeer:
			ps.trackConnectionStart(event.ID)
		}
	}
}

func (ps *ProvideState) consumeRPCEvents(rpcEvents <-chan interface{}) {
	for event := range rpcEvents {
		switch evt := event.(type) {
		case *wrap.RPCSendRequestStartedEvent:
		case *wrap.RPCSendRequestEndedEvent:
			switch evt.Request.Type {
			case pb.Message_FIND_NODE:
				ps.trackFindNodeRequest(evt)
			default:
				log.Warn(evt)
			}
		case *wrap.RPCSendMessageStartedEvent:
		case *wrap.RPCSendMessageEndedEvent:
			switch evt.Message.Type {
			case pb.Message_ADD_PROVIDER:
				ps.trackAddProvidersRequest(evt)
			default:
				log.Warn(evt)
			}
		default:
			log.Warn(event)
		}
	}
}

func (ps *ProvideState) ensurePeerset(evt *kaddht.LookupEvent) {
	if _, found := ps.peerSet[evt.ID]; !found {
		ps.peerSet[evt.ID] = qpeerset.NewQueryPeerset(string(ps.content.CID.Hash()))
	}
}

func (ps *ProvideState) trackLookupResponse(evt *kaddht.LookupEvent) {
	for _, p := range evt.Response.Heard {
		if p.Peer == ps.h.ID() { // don't add self.
			continue
		}
		ps.peerSet[evt.ID].TryAdd(p.Peer, evt.Response.Cause.Peer)
	}
	for _, p := range evt.Response.Queried {
		if p.Peer == ps.h.ID() { // don't add self.
			continue
		}
		if st := ps.peerSet[evt.ID].GetState(p.Peer); st == qpeerset.PeerWaiting {
			ps.peerSet[evt.ID].SetState(p.Peer, qpeerset.PeerQueried)
		} else {
			panic(fmt.Errorf("kademlia protocol error: tried to transition to the queried state from state %v", st))
		}
	}
	for _, p := range evt.Response.Unreachable {
		if p.Peer == ps.h.ID() { // don't add self.
			continue
		}

		if st := ps.peerSet[evt.ID].GetState(p.Peer); st == qpeerset.PeerWaiting {
			ps.peerSet[evt.ID].SetState(p.Peer, qpeerset.PeerUnreachable)
		} else {
			panic(fmt.Errorf("kademlia protocol error: tried to transition to the unreachable state from state %v", st))
		}
	}
}

func (ps *ProvideState) trackLookupRequest(evt *kaddht.LookupEvent) {
	for _, p := range evt.Request.Waiting {
		ps.peerSet[evt.ID].SetState(p.Peer, qpeerset.PeerWaiting)
	}
}

func (ps *ProvideState) trackConnectionStart(p peer.ID) {
	ps.connectionsStartedLk.Lock()
	defer ps.connectionsStartedLk.Unlock()
	ps.connectionsStarted[p] = time.Now()
}

func (ps *ProvideState) trackConnectionEnd(p peer.ID, maddr ma.Multiaddr) {
	end := time.Now()

	ps.connectionsStartedLk.Lock()
	started, ok := ps.connectionsStarted[p]
	if !ok {
		ps.connectionsStartedLk.Unlock()
		return
	}
	delete(ps.connectionsStarted, p)
	ps.connectionsStartedLk.Unlock()

	ps.connectionsLk.Lock()
	ps.connections = append(ps.connections, &ConnectionSpan{
		RemotePeerID: p,
		Maddr:        maddr,
		Start:        started,
		End:          end,
	})
	ps.connectionsLk.Unlock()
}

func (ps *ProvideState) trackFindNodeRequest(evt *wrap.RPCSendRequestEndedEvent) {
	fns := &FindNodesSpan{
		QueryID:      evt.QueryID,
		RemotePeerID: evt.RemotePeer,
		Start:        evt.StartedAt,
		End:          evt.EndedAt,
		Error:        evt.Error,
	}
	if evt.Response != nil {
		fns.CloserPeers = pb.PBPeersToPeerInfos(evt.Response.CloserPeers)
	}

	ps.findNodesLk.Lock()
	defer ps.findNodesLk.Unlock()

	ps.findNodes = append(ps.findNodes, fns)
	ps.relevantPeers.Store(evt.RemotePeer, struct{}{})
	for _, p := range fns.CloserPeers {
		ps.relevantPeers.Store(p.ID, struct{}{})
	}
}

func (ps *ProvideState) trackAddProvidersRequest(evt *wrap.RPCSendMessageEndedEvent) {
	aps := &AddProvidersSpan{
		QueryID:       evt.QueryID,
		Content:       ps.content,
		RemotePeerID:  evt.RemotePeer,
		Start:         evt.StartedAt,
		End:           evt.EndedAt,
		ProviderAddrs: pb.PBPeersToPeerInfos(evt.Message.ProviderPeers)[0].Addrs,
		Error:         evt.Error,
	}

	ps.addProvidersLk.Lock()
	defer ps.addProvidersLk.Unlock()

	ps.addProviders = append(ps.addProviders, aps)
	ps.relevantPeers.Store(evt.RemotePeer, struct{}{})
}

func (ps *ProvideState) filterDials() {
	ps.dialsLk.Lock()
	defer ps.dialsLk.Unlock()

	var relevantDials []*DialSpan
	for _, dial := range ps.dials {
		if _, found := ps.relevantPeers.Load(dial.RemotePeerID); found {
			relevantDials = append(relevantDials, dial)
		}
	}
	ps.dials = relevantDials
}

func (ps *ProvideState) filterConnections() {
	ps.connectionsLk.Lock()
	defer ps.connectionsLk.Unlock()

	var relevantConnections []*ConnectionSpan
	for _, conn := range ps.connections {
		if _, found := ps.relevantPeers.Load(conn.RemotePeerID); found {
			relevantConnections = append(relevantConnections, conn)
		}
	}
	ps.connections = relevantConnections
}

////

func (ps *ProvideState) Listen(network network.Network, multiaddr ma.Multiaddr) {
}

func (ps *ProvideState) ListenClose(network network.Network, multiaddr ma.Multiaddr) {
}

func (ps *ProvideState) Connected(network network.Network, conn network.Conn) {
	ps.trackConnectionEnd(conn.RemotePeer(), conn.RemoteMultiaddr())
}

func (ps *ProvideState) Disconnected(network network.Network, conn network.Conn) {
}

func (ps *ProvideState) OpenedStream(network network.Network, stream network.Stream) {
}

func (ps *ProvideState) ClosedStream(network network.Network, stream network.Stream) {
}

////

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
