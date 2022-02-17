package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/lib"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type RoutingTableService interface {
	SaveRoutingTable(ctx context.Context, h *dht.Host) (*models.RoutingTableSnapshot, error)
}

var _ RoutingTableService = &RoutingTable{}

type RoutingTable struct {
	peerService PeerService
	rtRepo      repo.RoutingTableRepo
}

func NewRoutingTableService(peerService PeerService, rtRepo repo.RoutingTableRepo) RoutingTableService {
	return &RoutingTable{
		peerService: peerService,
		rtRepo:      rtRepo,
	}
}

func (rts *RoutingTable) SaveRoutingTable(ctx context.Context, h *dht.Host) (*models.RoutingTableSnapshot, error) {
	localDbPeer, err := rts.peerService.UpsertLocalPeer(h.Host)
	if err != nil {
		return nil, errors.Wrap(err, "upsert local peer")
	}

	rt := h.DHT.RoutingTable()
	swarm := h.Host.Network()

	snapshot, err := rts.rtRepo.SaveSnapshot(ctx, localDbPeer.ID, lib.DefaultBucketSize, len(rt.GetPeerInfos()))
	if err != nil {
		return nil, errors.Wrap(err, "insert routing table")
	}

	for _, peerInfo := range rt.GetPeerInfos() {
		dbpeer, err := rts.peerService.UpsertPeer(h.Host, peerInfo.Id)
		if err != nil {
			return nil, errors.Wrap(err, "upsert peer")
		}

		var connectedAt *time.Time
		for _, conn := range swarm.ConnsToPeer(peerInfo.Id) {
			opened := conn.Stat().Opened
			if connectedAt == nil || connectedAt.After(opened) {
				connectedAt = &opened
			}
		}

		rte := &models.RoutingTableEntry{
			RoutingTableSnapshotID:        snapshot.ID,
			PeerID:                        dbpeer.ID,
			Bucket:                        lib.BucketIdForPeer(h.PeerID, peerInfo.Id),
			LastUsefulAt:                  null.NewTime(peerInfo.LastUsefulAt, !peerInfo.LastUsefulAt.IsZero()),
			LastSuccessfulOutboundQueryAt: peerInfo.LastSuccessfulOutboundQueryAt,
			AddedAt:                       peerInfo.AddedAt,
			ConnectedAt:                   null.TimeFromPtr(connectedAt),
		}

		if _, err = rts.rtRepo.SaveRoutingTableEntry(ctx, rte); err != nil {
			return nil, errors.Wrap(err, "insert routing table entry")
		}
	}

	return snapshot, nil
}
