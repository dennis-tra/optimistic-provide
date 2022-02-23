package service

import (
	"sort"
	"time"

	"github.com/volatiletech/null/v8"

	"go.uber.org/atomic"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
	"github.com/libp2p/go-libp2p-core/peer"
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

func (r *RoutingTableListener) PeerAdded(p peer.ID) {
	r.SendUpdate()
}

func (r *RoutingTableListener) PeerRemoved(p peer.ID) {
	r.SendUpdate()
}

func (r *RoutingTableListener) SendUpdate() {
	r.updateChan <- r.BuildUpdate()
}

func (r *RoutingTableListener) BuildUpdate() types.RoutingTableUpdate {
	infos := r.h.DHT.RoutingTable().GetPeerInfos()
	swarm := r.h.Network()

	var rtp types.RoutingTableUpdate = make([]types.RoutingTablePeer, len(infos))
	for i, info := range infos {
		var connectedAt *time.Time
		for _, conn := range swarm.ConnsToPeer(info.Id) {
			opened := conn.Stat().Opened
			if connectedAt == nil || connectedAt.After(opened) {
				connectedAt = &opened
			}
		}

		av := ""
		if agent, err := r.h.Peerstore().Get(info.Id, "AgentVersion"); err == nil {
			av = agent.(string)
		}

		rtp[i] = types.RoutingTablePeer{
			PeerId:                        info.Id.String(),
			AddedAt:                       info.AddedAt.Format(time.RFC3339),
			AgentVersion:                  null.NewString(av, av != "").Ptr(),
			Bucket:                        int(util.BucketIdForPeer(r.h.ID(), info.Id)),
			ConnectedSince:                util.TimeToStr(connectedAt),
			LastSuccessfulOutboundQueryAt: info.LastSuccessfulOutboundQueryAt.Format(time.RFC3339),
			LastUsefulAt:                  util.TimeToStr(&info.LastUsefulAt),
		}
	}

	sort.Slice(rtp, func(i, j int) bool {
		ti, _ := time.Parse(time.RFC3339, rtp[i].AddedAt)
		tj, _ := time.Parse(time.RFC3339, rtp[j].AddedAt)
		return ti.After(tj)
	})

	return rtp
}

func (r *RoutingTableListener) OnClose() {
	r.Stop()
}
