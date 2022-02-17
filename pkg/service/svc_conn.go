package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/libp2p/go-libp2p-core/host"
)

type ConnectionService interface {
	Save(ctx context.Context, h host.Host, provideID int, conns []*ConnectionSpan) error
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

func (c *Connection) Save(ctx context.Context, h host.Host, provideID int, conns []*ConnectionSpan) error {
	log.Info("Saving connections...")
	defer log.Info("Done saving connections")

	localPeer, err := c.peerService.UpsertLocalPeer(h)
	if err != nil {
		return err
	}

	for _, conn := range conns {
		remotePeer, err := c.peerService.UpsertPeer(h, conn.RemotePeerID)
		if err != nil {
			return err
		}

		maddr, err := c.maService.UpsertMultiAddress(ctx, conn.Maddr)
		if err != nil {
			return err
		}

		dbConn := &models.Connection{
			ProvideID:      provideID,
			LocalID:        localPeer.ID,
			RemoteID:       remotePeer.ID,
			MultiAddressID: maddr.ID,
			StartedAt:      conn.Start,
			EndedAt:        conn.End,
		}

		if dbConn, err = c.repo.Save(ctx, dbConn); err != nil {
			return err
		}
	}

	return nil
}
