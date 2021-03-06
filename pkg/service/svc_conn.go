package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type ConnectionService interface {
	List(ctx context.Context, provide *models.Provide) (models.ConnectionSlice, error)
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, conns []*ConnectionSpan, peerInfos map[peer.ID]*PeerInfo) (models.ConnectionSlice, error)
}

var _ ConnectionService = &Connection{}

type Connection struct {
	peerService PeerService
	maService   MultiAddressService
	repo        repo.ConnectionRepo
}

func (c *Connection) List(ctx context.Context, provide *models.Provide) (models.ConnectionSlice, error) {
	return c.repo.ListFromProvide(ctx, provide)
}

func NewConnectionService(peerService PeerService, maService MultiAddressService, repo repo.ConnectionRepo) ConnectionService {
	return &Connection{
		peerService: peerService,
		maService:   maService,
		repo:        repo,
	}
}

func (c *Connection) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, conns []*ConnectionSpan, peerInfos map[peer.ID]*PeerInfo) (models.ConnectionSlice, error) {
	log.Info("Saving Connections")

	dbConns := make([]*models.Connection, len(conns))
	for i, conn := range conns {
		remotePeer, err := c.peerService.UpsertPeerForInfo(ctx, exec, h, conn.RemotePeerID, peerInfos[conn.RemotePeerID])
		if err != nil {
			return nil, err
		}

		maddr, err := c.maService.UpsertMultiAddress(ctx, exec, conn.Maddr)
		if err != nil {
			return nil, err
		}

		dbConn := &models.Connection{
			LocalID:        h.DBHost.PeerID,
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
