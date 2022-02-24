package service

import (
	"context"

	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/libp2p/go-libp2p-core/host"
)

type FindNodesService interface {
	Save(ctx context.Context, h host.Host, provideID int, fnReqs []*FindNodesSpan) error
	List(ctx context.Context, provideID int) ([]*models.FindNode, error)
}

var _ FindNodesService = &FindNodes{}

type FindNodes struct {
	peerService PeerService
	fnRepo      repo.FindNodesRepo
	cpRepo      repo.CloserPeersRepo
}

func NewFindNodesService(peerService PeerService, fnRepo repo.FindNodesRepo, cpRepo repo.CloserPeersRepo) FindNodesService {
	return &FindNodes{
		peerService: peerService,
		fnRepo:      fnRepo,
		cpRepo:      cpRepo,
	}
}

func (fn *FindNodes) List(ctx context.Context, provideID int) ([]*models.FindNode, error) {
	return fn.fnRepo.List(ctx, provideID)
}

func (fn *FindNodes) Save(ctx context.Context, h host.Host, provideID int, fnReqs []*FindNodesSpan) error {
	log.Info("Saving find nodes requests...")
	defer log.Info("Done saving find nodes requests")

	localPeer, err := fn.peerService.UpsertLocalPeer(h)
	if err != nil {
		return err
	}

	for _, fnReq := range fnReqs {

		remotePeer, err := fn.peerService.UpsertPeer(h, fnReq.RemotePeerID)
		if err != nil {
			return err
		}

		errStr := null.NewString("", false)
		cpCount := null.NewInt(0, false)
		if fnReq.Error != nil {
			errStr = null.StringFrom(fnReq.Error.Error())
		} else {
			cpCount = null.IntFrom(len(fnReq.CloserPeers))
		}

		dbfn := &models.FindNode{
			ProvideID:        provideID,
			LocalID:          localPeer.ID,
			RemoteID:         remotePeer.ID,
			StartedAt:        fnReq.Start,
			EndedAt:          fnReq.End,
			Error:            errStr,
			CloserPeersCount: cpCount,
		}
		dbfn, err = fn.fnRepo.Save(ctx, dbfn)
		if err != nil {
			return err
		}

		for _, closerPeer := range fnReq.CloserPeers {

			cp, err := fn.peerService.UpsertPeer(h, closerPeer.ID)
			if err != nil {
				return err
			}

			dbcp := &models.CloserPeer{
				ProvideID:  provideID,
				FindNodeID: dbfn.ID,
				PeerID:     cp.ID,
			}
			if _, err = fn.cpRepo.Save(ctx, dbcp); err != nil {
				return err
			}
		}

	}
	return nil
}
