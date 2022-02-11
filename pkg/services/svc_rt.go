package services

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/db/models"
	"github.com/dennis-tra/optimistic-provide/pkg/host"
	"github.com/dennis-tra/optimistic-provide/pkg/utils"
)

type RoutingTableService struct {
	dbc *db.Client
}

func NewRoutingTableService(dbc *db.Client) *RoutingTableService {
	return &RoutingTableService{
		dbc: dbc,
	}
}

func (rts RoutingTableService) SaveRoutingTable(ctx context.Context, h *host.Host) (*models.RoutingTableSnapshot, error) {
	localDbPeer, err := rts.dbc.UpsertLocalPeer(h.PeerID)
	if err != nil {
		return nil, errors.Wrap(err, "upsert local peer")
	}

	ps := h.Host.Peerstore()
	rt := h.DHT.RoutingTable()
	swarm := h.Host.Network()

	txn, err := rts.dbc.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "begin txn")
	}

	dbrt := &models.RoutingTableSnapshot{
		PeerID:     localDbPeer.ID,
		BucketSize: utils.DefaultBucketSize,
		EntryCount: len(rt.GetPeerInfos()),
	}
	if err := dbrt.Insert(ctx, txn, boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "insert routing table")
	}

	for _, peerInfo := range rt.GetPeerInfos() {
		av := ""
		if agent, err := ps.Get(peerInfo.Id, "AgentVersion"); err == nil {
			av = agent.(string)
		}

		protocols := []string{}
		if prots, err := ps.GetProtocols(peerInfo.Id); err == nil {
			protocols = prots
		}

		dbpeer, err := rts.dbc.UpsertPeer(txn, peerInfo.Id, av, protocols)
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

		dbrte := models.RoutingTableEntry{
			RoutingTableSnapshotID:        dbrt.ID,
			PeerID:                        dbpeer.ID,
			Bucket:                        utils.BucketIdForPeer(h.PeerID, peerInfo.Id),
			LastUsefulAt:                  null.NewTime(peerInfo.LastUsefulAt, !peerInfo.LastUsefulAt.IsZero()),
			LastSuccessfulOutboundQueryAt: peerInfo.LastSuccessfulOutboundQueryAt,
			AddedAt:                       peerInfo.AddedAt,
			ConnectedAt:                   null.TimeFromPtr(connectedAt),
		}

		if err = dbrte.Insert(ctx, txn, boil.Infer()); err != nil {
			return nil, errors.Wrap(err, "insert routing table entry")
		}
	}

	if err := txn.Commit(); err != nil {
		return nil, txn.Rollback()
	}
	return dbrt, nil
}
