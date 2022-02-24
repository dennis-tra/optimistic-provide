package service

import (
	"context"

	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/libp2p/go-libp2p-core/host"
)

type ConnectionService interface {
	Save(ctx context.Context, h host.Host, op HostOperation, id int, conns []*ConnectionSpan) error
	List(ctx context.Context, op HostOperation, id int) ([]*models.Connection, error)
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

func (c *Connection) List(ctx context.Context, op HostOperation, id int) ([]*models.Connection, error) {
	switch op {
	case HostOperationProvide:
		return c.repo.ListFromProvide(ctx, id)
	case HostOperationRetrieval:
		return c.repo.ListFromRetrieval(ctx, id)
	default:
		panic(op)
	}
}

func (c *Connection) Save(ctx context.Context, h host.Host, op HostOperation, id int, conns []*ConnectionSpan) error {
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
			LocalID:        localPeer.ID,
			RemoteID:       remotePeer.ID,
			MultiAddressID: maddr.ID,
			StartedAt:      conn.Start,
			EndedAt:        conn.End,
		}

		switch op {
		case HostOperationProvide:
			dbConn.ProvideID = null.IntFrom(id)
		case HostOperationRetrieval:
			dbConn.RetrievalID = null.IntFrom(id)
		default:
			panic(op)
		}

		if dbConn, err = c.repo.Save(ctx, dbConn); err != nil {
			return err
		}
	}

	return nil
}
