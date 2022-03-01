package service

import (
	"sort"
	"time"

	"github.com/libp2p/go-libp2p-core/network"

	kbucket "github.com/libp2p/go-libp2p-kbucket"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/volatiletech/null/v8"
	"go.uber.org/atomic"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
)

type RoutingTableListener struct {
	h          *dht.Host
	updateChan chan types.RoutingTableUpdate
	stopped    *atomic.Bool
}

func NewRoutingTableListener(h *dht.Host) *RoutingTableListener {
	return &RoutingTableListener{
		h:          h,
		updateChan: make(chan types.RoutingTableUpdate),
		stopped:    atomic.NewBool(false),
	}
}

var _ dht.RoutingTableListener = (*RoutingTableListener)(nil)

func (r *RoutingTableListener) Updates() <-chan types.RoutingTableUpdate {
	return r.updateChan
}

func (r *RoutingTableListener) Stop() {
	old := r.stopped.Swap(true)
	if old {
		return
	}
	close(r.updateChan)
}

func (r *RoutingTableListener) OnClose() {
	r.Stop()
}

func (r *RoutingTableListener) PeerAdded(p peer.ID) {
	infos := r.h.DHT.RoutingTable().GetPeerInfos()
	swarm := r.h.Network()

	var peerInfo *kbucket.PeerInfo
	for _, info := range infos {
		if info.Id == p {
			peerInfo = &info
			break
		}
	}
	if peerInfo == nil {
		// TODO: log
		return
	}

	r.updateChan <- types.RoutingTableUpdate{
		Type:   types.RoutingTableUpdateTypePEERADDED,
		Update: r.buildRoutingTablePeer(swarm, *peerInfo),
	}
}

func (r *RoutingTableListener) PeerRemoved(p peer.ID) {
	r.updateChan <- types.RoutingTableUpdate{
		Type:   types.RoutingTableUpdateTypePEERREMOVED,
		Update: p.String(),
	}
}

func (r *RoutingTableListener) SendFullUpdate() {
	r.updateChan <- types.RoutingTableUpdate{
		Type:   types.RoutingTableUpdateTypeFULL,
		Update: r.BuildUpdate(),
	}
}

func (r *RoutingTableListener) BuildUpdate() types.RoutingTablePeers {
	infos := r.h.DHT.RoutingTable().GetPeerInfos()
	swarm := r.h.Network()

	var rtp types.RoutingTablePeers = make([]types.RoutingTablePeer, len(infos))
	for i, info := range infos {
		rtp[i] = r.buildRoutingTablePeer(swarm, info)
	}

	sort.Slice(rtp, func(i, j int) bool {
		ti, _ := time.Parse(time.RFC3339Nano, rtp[i].AddedAt)
		tj, _ := time.Parse(time.RFC3339Nano, rtp[j].AddedAt)
		return ti.After(tj)
	})

	return rtp
}

func (r *RoutingTableListener) buildRoutingTablePeer(swarm network.Network, peerInfo kbucket.PeerInfo) types.RoutingTablePeer {
	var connectedAt *time.Time
	for _, conn := range swarm.ConnsToPeer(peerInfo.Id) {
		opened := conn.Stat().Opened
		if connectedAt == nil || connectedAt.After(opened) {
			connectedAt = &opened
		}
	}

	av := ""
	if agent, err := r.h.Peerstore().Get(peerInfo.Id, "AgentVersion"); err == nil {
		av = agent.(string)
	}

	return types.RoutingTablePeer{
		PeerId:                        peerInfo.Id.String(),
		AddedAt:                       peerInfo.AddedAt.Format(time.RFC3339Nano),
		AgentVersion:                  null.NewString(av, av != "").Ptr(),
		Bucket:                        int(util.BucketIdForPeer(r.h.ID(), peerInfo.Id)),
		ConnectedSince:                util.TimeToStr(connectedAt),
		LastSuccessfulOutboundQueryAt: peerInfo.LastSuccessfulOutboundQueryAt.Format(time.RFC3339Nano),
		LastUsefulAt:                  util.TimeToStr(&peerInfo.LastUsefulAt),
	}
}
