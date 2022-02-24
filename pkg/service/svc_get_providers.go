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
	ppService   ProviderPeersService
	gpRepo      repo.GetProvidersRepo
}

func NewGetProvidersService(peerService PeerService, ppService ProviderPeersService, gpRepo repo.GetProvidersRepo) GetProvidersService {
	return &GetProviders{
		peerService: peerService,
		gpRepo:      gpRepo,
		ppService:   ppService,
	}
}

func (gp *GetProviders) List(ctx context.Context, retrievalID int) ([]*models.GetProvider, error) {
	return gp.gpRepo.List(ctx, retrievalID)
}

func (gp *GetProviders) Save(ctx context.Context, h host.Host, retrievalID int, fnReqs []*GetProvidersSpan) error {
	log.Info("Saving get providers requests...")
	defer log.Info("Done saving get providers requests")

	localPeer, err := gp.peerService.UpsertLocalPeer(h)
	if err != nil {
		return err
	}

	for _, gpReq := range fnReqs {

		remotePeer, err := gp.peerService.UpsertPeer(h, gpReq.RemotePeerID)
		if err != nil {
			return err
		}

		errStr := null.NewString("", false)
		ppCount := null.NewInt(0, false)
		if gpReq.Error != nil {
			errStr = null.StringFrom(gpReq.Error.Error())
		} else {
			ppCount = null.IntFrom(len(gpReq.Providers))
		}

		dbgp := &models.GetProvider{
			RetrievalID:        retrievalID,
			LocalID:            localPeer.ID,
			RemoteID:           remotePeer.ID,
			StartedAt:          gpReq.Start,
			EndedAt:            gpReq.End,
			Error:              errStr,
			ProviderPeersCount: ppCount,
		}
		dbgp, err = gp.gpRepo.Save(ctx, dbgp)
		if err != nil {
			return err
		}

		for _, provider := range gpReq.Providers {
			if _, err = gp.ppService.Save(ctx, h, dbgp.ID, *provider); err != nil {
				return err
			}
		}
	}
	return nil
}
