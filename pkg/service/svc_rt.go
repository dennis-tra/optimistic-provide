package service

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
)

type RoutingTableService interface {
	Find(ctx context.Context, id int) (*models.RoutingTableSnapshot, error)
	FindAll(ctx context.Context, hostID peer.ID) ([]*models.RoutingTableSnapshot, error)
	FindByIDAndHostID(ctx context.Context, id int, hostID peer.ID) (*models.RoutingTableSnapshot, error)
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host) (*models.RoutingTableSnapshot, error)
	SaveTxn(ctx context.Context, h *dht.Host) (*models.RoutingTableSnapshot, error)
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

func (rts *RoutingTable) FindByIDAndHostID(ctx context.Context, id int, hostID peer.ID) (*models.RoutingTableSnapshot, error) {
	panic("implement me")
}

func (rts *RoutingTable) Find(ctx context.Context, id int) (*models.RoutingTableSnapshot, error) {
	return rts.rtRepo.Find(ctx, id)
}

func (rts *RoutingTable) FindAll(ctx context.Context, hostID peer.ID) ([]*models.RoutingTableSnapshot, error) {
	return rts.rtRepo.FindAll(ctx, hostID.String())
}

func (rts *RoutingTable) SaveTxn(ctx context.Context, h *dht.Host) (*models.RoutingTableSnapshot, error) {
	txn, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}

	snapshot, err := rts.Save(ctx, txn, h)
	if err != nil {
		if err2 := txn.Rollback(); err != nil {
			log.Warn("error rolling back transaction", err2)
		}
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		if err2 := txn.Rollback(); err != nil {
			log.Warn("error rolling back transaction", err2)
		}
		return nil, errors.Wrap(err, "committing transaction")
	}

	return snapshot, nil
}

func (rts *RoutingTable) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host) (*models.RoutingTableSnapshot, error) {
	log.Info("Saving Routing Table")

	rt := h.DHT.RoutingTable()
	swarm := h.Host.Network()

	snapshot, err := rts.rtRepo.SaveSnapshot(ctx, exec, h.DBHost.PeerID, util.DefaultBucketSize, len(rt.GetPeerInfos()))
	if err != nil {
		return nil, errors.Wrap(err, "insert routing table")
	}

	for _, peerInfo := range rt.GetPeerInfos() {
		dbpeer, err := rts.peerService.UpsertPeer(ctx, exec, h.Host, peerInfo.Id)
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
			Bucket:                        util.BucketIdForPeer(h.ID(), peerInfo.Id),
			LastUsefulAt:                  null.NewTime(peerInfo.LastUsefulAt, !peerInfo.LastUsefulAt.IsZero()),
			LastSuccessfulOutboundQueryAt: peerInfo.LastSuccessfulOutboundQueryAt,
			AddedAt:                       peerInfo.AddedAt,
			ConnectedSince:                null.TimeFromPtr(connectedAt),
		}

		if _, err = rts.rtRepo.SaveRoutingTableEntry(ctx, exec, rte); err != nil {
			return nil, errors.Wrap(err, "insert routing table entry")
		}
	}
	return snapshot, nil
}
