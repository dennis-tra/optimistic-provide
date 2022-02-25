package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ipfs/go-cid"

	kaddht "github.com/libp2p/go-libp2p-kad-dht"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	"github.com/libp2p/go-libp2p-kad-dht/qpeerset"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/wrap"
)

type RetrievalState struct {
	h       *dht.Host
	content cid.Cid
	cancel  context.CancelFunc

	dialsLk sync.RWMutex
	dials   []*DialSpan

	getProvidersLk sync.RWMutex
	getProviders   []*GetProvidersSpan

	connectionsStartedLk sync.RWMutex
	connectionsStarted   map[peer.ID]time.Time

	connectionsLk sync.RWMutex
	connections   []*ConnectionSpan

	peerSet *qpeerset.QueryPeerset

	relevantPeers sync.Map
}

func NewRetrievalState(h *dht.Host, contentID cid.Cid) *RetrievalState {
	return &RetrievalState{
		h:                    h,
		content:              contentID,
		dialsLk:              sync.RWMutex{},
		dials:                []*DialSpan{},
		getProvidersLk:       sync.RWMutex{},
		getProviders:         []*GetProvidersSpan{},
		connectionsStartedLk: sync.RWMutex{},
		connectionsStarted:   map[peer.ID]time.Time{},
		connectionsLk:        sync.RWMutex{},
		connections:          []*ConnectionSpan{},
		relevantPeers:        sync.Map{},
		peerSet:              qpeerset.NewQueryPeerset(string(contentID.Hash())),
	}
}

func (rs *RetrievalState) Register(ctx context.Context) context.Context {
	if rs.cancel != nil {
		panic("already registered for events")
	}
	ctx, cancel := context.WithCancel(ctx)
	ctx, lookupEvents := kaddht.RegisterForLookupEvents(ctx)
	ctx, queryEvents := routing.RegisterForQueryEvents(ctx)
	ctx, rpcEvents := wrap.RegisterForRPCEvents(ctx)

	go rs.consumeLookupEvents(lookupEvents)
	go rs.consumeQueryEvents(queryEvents)
	go rs.consumeRPCEvents(rpcEvents)

	go rs.consumeRPCEvents(rpcEvents)

	for _, notifier := range rs.h.Transports {
		notifier.Notify(rs)
	}

	rs.h.Host.Network().Notify(rs)

	rs.cancel = cancel

	return ctx
}

func (rs *RetrievalState) Unregister() {
	rs.h.Host.Network().StopNotify(rs)

	for _, notifier := range rs.h.Transports {
		notifier.StopNotify(rs)
	}

	rs.cancel()
	rs.cancel = nil

	rs.filterDials()
	rs.filterConnections()
}

func (rs *RetrievalState) consumeQueryEvents(queryEvents <-chan *routing.QueryEvent) {
	for event := range queryEvents {
		switch event.Type {
		case routing.DialingPeer:
			rs.trackConnectionStart(event.ID)
		}
	}
}

func (rs *RetrievalState) trackConnectionStart(p peer.ID) {
	rs.connectionsStartedLk.Lock()
	defer rs.connectionsStartedLk.Unlock()
	rs.connectionsStarted[p] = time.Now()
}

func (rs *RetrievalState) trackConnectionEnd(p peer.ID, maddr ma.Multiaddr) {
	end := time.Now()

	rs.connectionsStartedLk.Lock()
	started, ok := rs.connectionsStarted[p]
	if !ok {
		rs.connectionsStartedLk.Unlock()
		return
	}
	delete(rs.connectionsStarted, p)
	rs.connectionsStartedLk.Unlock()

	rs.connectionsLk.Lock()
	rs.connections = append(rs.connections, &ConnectionSpan{
		RemotePeerID: p,
		Maddr:        maddr,
		Start:        started,
		End:          end,
	})
	rs.connectionsLk.Unlock()
}

func (rs *RetrievalState) filterDials() {
	rs.dialsLk.Lock()
	defer rs.dialsLk.Unlock()

	var relevantDials []*DialSpan
	for _, dial := range rs.dials {
		if _, found := rs.relevantPeers.Load(dial.RemotePeerID); found {
			relevantDials = append(relevantDials, dial)
		}
	}
	rs.dials = relevantDials
}

func (rs *RetrievalState) filterConnections() {
	rs.connectionsLk.Lock()
	defer rs.connectionsLk.Unlock()

	var relevantConnections []*ConnectionSpan
	for _, conn := range rs.connections {
		if _, found := rs.relevantPeers.Load(conn.RemotePeerID); found {
			relevantConnections = append(relevantConnections, conn)
		}
	}
	rs.connections = relevantConnections
}

////

func (rs *RetrievalState) consumeLookupEvents(lookupEvents <-chan *kaddht.LookupEvent) {
	for event := range lookupEvents {
		if event.Response != nil {
			rs.trackLookupResponse(event)
		} else if event.Request != nil {
			rs.trackLookupRequest(event)
		}
	}
}

func (rs *RetrievalState) trackLookupRequest(evt *kaddht.LookupEvent) {
	for _, p := range evt.Request.Waiting {
		rs.peerSet.SetState(p.Peer, qpeerset.PeerWaiting)
	}
}

func (rs *RetrievalState) trackLookupResponse(evt *kaddht.LookupEvent) {
	for _, p := range evt.Response.Heard {
		if p.Peer == rs.h.ID() { // don't add self.
			continue
		}
		rs.peerSet.TryAdd(p.Peer, evt.Response.Cause.Peer)
	}
	for _, p := range evt.Response.Queried {
		if p.Peer == rs.h.ID() { // don't add self.
			continue
		}
		if st := rs.peerSet.GetState(p.Peer); st == qpeerset.PeerWaiting {
			rs.peerSet.SetState(p.Peer, qpeerset.PeerQueried)
		} else {
			panic(fmt.Errorf("kademlia protocol error: tried to transition to the queried state from state %v", st))
		}
	}
	for _, p := range evt.Response.Unreachable {
		if p.Peer == rs.h.ID() { // don't add self.
			continue
		}

		if st := rs.peerSet.GetState(p.Peer); st == qpeerset.PeerWaiting {
			rs.peerSet.SetState(p.Peer, qpeerset.PeerUnreachable)
		} else {
			panic(fmt.Errorf("kademlia protocol error: tried to transition to the unreachable state from state %v", st))
		}
	}
}

func (rs *RetrievalState) consumeRPCEvents(rpcEvents <-chan interface{}) {
	for event := range rpcEvents {
		switch evt := event.(type) {
		case *wrap.RPCSendRequestStartedEvent:
		case *wrap.RPCSendRequestEndedEvent:
			switch evt.Request.Type {
			case pb.Message_GET_PROVIDERS:
				rs.trackGetProvidersRequest(evt)
			default:
				log.Warn(evt)
			}
		case *wrap.RPCSendMessageStartedEvent:
		case *wrap.RPCSendMessageEndedEvent:
			switch evt.Message.Type {
			default:
				log.Warn(evt)
			}
		default:
			log.Warn(event)
		}
	}
}

func (rs *RetrievalState) trackGetProvidersRequest(evt *wrap.RPCSendRequestEndedEvent) {
	gps := &GetProvidersSpan{
		RemotePeerID: evt.RemotePeer,
		Start:        evt.StartedAt,
		End:          evt.EndedAt,
		Error:        evt.Error,
	}
	if evt.Response != nil {
		gps.Providers = pb.PBPeersToPeerInfos(evt.Response.ProviderPeers)
	}

	rs.getProvidersLk.Lock()
	defer rs.getProvidersLk.Unlock()

	rs.getProviders = append(rs.getProviders, gps)
	rs.relevantPeers.Store(evt.RemotePeer, struct{}{})
	for _, p := range gps.Providers {
		rs.relevantPeers.Store(p.ID, struct{}{})
	}
}

////

func (rs *RetrievalState) Listen(network network.Network, multiaddr ma.Multiaddr) {
}

func (rs *RetrievalState) ListenClose(network network.Network, multiaddr ma.Multiaddr) {
}

func (rs *RetrievalState) Connected(network network.Network, conn network.Conn) {
	rs.trackConnectionEnd(conn.RemotePeer(), conn.RemoteMultiaddr())
}

func (rs *RetrievalState) Disconnected(network network.Network, conn network.Conn) {
}

func (rs *RetrievalState) OpenedStream(network network.Network, stream network.Stream) {
}

func (rs *RetrievalState) ClosedStream(network network.Network, stream network.Stream) {
}

////

func (rs *RetrievalState) DialStarted(trpt string, raddr ma.Multiaddr, p peer.ID, start time.Time) {
}

func (rs *RetrievalState) DialEnded(trpt string, raddr ma.Multiaddr, p peer.ID, start time.Time, end time.Time, err error) {
	d := &DialSpan{
		RemotePeerID: p,
		Maddr:        raddr,
		Start:        start,
		End:          end,
		Trpt:         trpt,
		Error:        err,
	}
	rs.dialsLk.Lock()
	defer rs.dialsLk.Unlock()

	rs.dials = append(rs.dials, d)
}
