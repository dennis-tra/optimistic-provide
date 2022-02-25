package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type ConnectionService interface {
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, conns []*ConnectionSpan) (models.ConnectionSlice, error)
}

var _ ConnectionService = &Connection{}

type Connection struct {
	peerService PeerService
	maService   MultiAddressService
	repo        repo.ConnectionRepo
}

func NewConnectionService(peerService PeerService, maService MultiAddressService, repo repo.ConnectionRepo) ConnectionService {
	return &Connection{
		peerService: peerService,
		maService:   maService,
		repo:        repo,
	}
}

func (c *Connection) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, conns []*ConnectionSpan) (models.ConnectionSlice, error) {
	log.Info("Saving Connections")

	dbConns := make([]*models.Connection, len(conns))
	for i, conn := range conns {
		remotePeer, err := c.peerService.UpsertPeer(ctx, exec, h, conn.RemotePeerID)
		if err != nil {
			return nil, err
		}

		maddr, err := c.maService.UpsertMultiAddress(ctx, exec, conn.Maddr)
		if err != nil {
			return nil, err
		}

		dbConn := &models.Connection{
			LocalID:        h.DBPeer.ID,
			RemoteID:       remotePeer.ID,
			MultiAddressID: maddr.ID,
			StartedAt:      conn.Start,
			EndedAt:        conn.End,
		}

		if err = dbConn.Insert(ctx, exec, boil.Infer()); err != nil {
			return nil, err
		}

		dbConns[i] = dbConn
	}

	return dbConns, nil
}
