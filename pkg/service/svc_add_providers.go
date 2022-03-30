package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
	"github.com/volatiletech/sqlboiler/v4/boil"

	ks "github.com/whyrusleeping/go-keyspace"

	"github.com/friendsofgo/errors"

	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type AddProvidersService interface {
	List(ctx context.Context, provide *models.Provide) ([]*models.AddProviderRPC, error)
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, apReqs []*AddProvidersSpan) (models.AddProviderRPCSlice, error)
}

var _ AddProvidersService = &AddProviders{}

type AddProviders struct {
	peerService PeerService
	maService   MultiAddressService
	apRepo      repo.AddProvidersRepo
	cpRepo      repo.CloserPeersRepo
}

func NewAddProvidersService(peerService PeerService, maService MultiAddressService, apRepo repo.AddProvidersRepo, cpRepo repo.CloserPeersRepo) AddProvidersService {
	return &AddProviders{
		peerService: peerService,
		maService:   maService,
		apRepo:      apRepo,
		cpRepo:      cpRepo,
	}
}

func (ap *AddProviders) List(ctx context.Context, provide *models.Provide) ([]*models.AddProviderRPC, error) {
	return ap.apRepo.List(ctx, provide)
}

func (ap *AddProviders) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, apReqs []*AddProvidersSpan) (models.AddProviderRPCSlice, error) {
	log.Info("Saving Add Provider RPCs")

	dbaps := make([]*models.AddProviderRPC, len(apReqs))
	for i, apReq := range apReqs {
		remotePeer, err := ap.peerService.UpsertPeer(ctx, exec, h, apReq.RemotePeerID)
		if err != nil {
			return nil, err
		}

		maids, err := ap.maService.UpsertMultiAddresses(ctx, exec, apReq.ProviderAddrs)
		if err != nil {
			return nil, errors.Wrap(err, "upsert multi address")
		}

		dbap := &models.AddProviderRPC{
			LocalID:         h.DBHost.PeerID,
			RemoteID:        remotePeer.ID,
			Distance:        ks.XORKeySpace.Key([]byte(apReq.RemotePeerID)).Distance(ks.XORKeySpace.Key(apReq.Content.CID.Hash())).Bytes(),
			MultiAddressIds: maids,
			StartedAt:       apReq.Start,
			EndedAt:         apReq.End,
			Error:           null.StringFromPtr(util.ErrorStr(apReq.Error)),
		}

		if err = dbap.Insert(ctx, exec, boil.Infer()); err != nil {
			return nil, err
		}

		dbaps[i] = dbap
	}

	return dbaps, nil
}
