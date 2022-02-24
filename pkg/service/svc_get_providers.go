package service

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type GetProvidersService interface {
	Save(ctx context.Context, h host.Host, retrievalID int, fnReqs []*GetProvidersSpan) error
	List(ctx context.Context, retrievalID int) ([]*models.GetProvider, error)
}

var _ GetProvidersService = &GetProviders{}

type GetProviders struct {
	peerService PeerService
	fnRepo      repo.GetProvidersRepo
	cpRepo      repo.CloserPeersRepo
}

func NewGetProvidersService(peerService PeerService, fnRepo repo.GetProvidersRepo, cpRepo repo.CloserPeersRepo) GetProvidersService {
	return &GetProviders{
		peerService: peerService,
		fnRepo:      fnRepo,
		cpRepo:      cpRepo,
	}
}

func (fn *GetProviders) List(ctx context.Context, retrievalID int) ([]*models.GetProvider, error) {
	return fn.fnRepo.List(ctx, retrievalID)
}

func (fn *GetProviders) Save(ctx context.Context, h host.Host, retrievalID int, fnReqs []*GetProvidersSpan) error {
	log.Info("Saving get providers requests...")
	defer log.Info("Done saving get providers requests")

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

		dbfn := &models.GetProvider{
			RetrievalID:      retrievalID,
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
				ProvideID:     retrievalID,
				GetProviderID: dbfn.ID,
				PeerID:        cp.ID,
			}
			if _, err = fn.cpRepo.Save(ctx, dbcp); err != nil {
				return err
			}
		}

	}
	return nil
}
