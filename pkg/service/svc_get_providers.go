package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/util"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type GetProvidersService interface {
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, fnReqs []*GetProvidersSpan) (models.GetProvidersRPCSlice, error)
}

var _ GetProvidersService = &GetProviders{}

type GetProviders struct {
	peerService PeerService
	maService   MultiAddressService
	gpRepo      repo.GetProvidersRepo
}

func NewGetProvidersService(peerService PeerService, maService MultiAddressService, gpRepo repo.GetProvidersRepo) GetProvidersService {
	return &GetProviders{
		peerService: peerService,
		maService:   maService,
		gpRepo:      gpRepo,
	}
}

func (gp *GetProviders) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, gpReqs []*GetProvidersSpan) (models.GetProvidersRPCSlice, error) {
	dbGetProviders := make([]*models.GetProvidersRPC, len(gpReqs))
	for i, gpReq := range gpReqs {

		remotePeer, err := gp.peerService.UpsertPeer(ctx, exec, h, gpReq.RemotePeerID)
		if err != nil {
			return nil, err
		}

		dbgp := &models.GetProvidersRPC{
			QueryID:            gpReq.QueryID.String(),
			LocalID:            h.DBHost.PeerID,
			RemoteID:           remotePeer.ID,
			StartedAt:          gpReq.Start,
			EndedAt:            gpReq.End,
			Error:              null.StringFromPtr(util.ErrorStr(gpReq.Error)),
			ProviderPeersCount: null.NewInt(len(gpReq.Providers), gpReq.Error == nil),
		}

		if err = dbgp.Insert(ctx, exec, boil.Infer()); err != nil {
			return nil, err
		}

		dbProviders := make([]*models.ProviderPeer, len(gpReq.Providers))
		for j, provider := range gpReq.Providers {
			dbPeer, err := gp.peerService.UpsertPeer(ctx, exec, h, provider.ID)
			if err != nil {
				return nil, err
			}

			maids, err := gp.maService.UpsertMultiAddresses(ctx, exec, provider.Addrs)
			if err != nil {
				return nil, err
			}

			dbProvider := &models.ProviderPeer{
				ProviderID:      dbPeer.ID,
				MultiAddressIds: maids,
			}

			dbProviders[j] = dbProvider
		}

		if err = dbgp.AddProviderPeers(ctx, exec, true, dbProviders...); err != nil {
			return nil, err
		}

		dbGetProviders[i] = dbgp
	}

	return dbGetProviders, nil
}
