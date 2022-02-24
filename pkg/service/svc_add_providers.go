package service

import (
	"context"
	"sort"

	ks "github.com/whyrusleeping/go-keyspace"

	"github.com/friendsofgo/errors"

	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/libp2p/go-libp2p-core/host"
)

type AddProvidersService interface {
	Save(ctx context.Context, h host.Host, provideID int, apReqs []*AddProvidersSpan) error
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

func (ap *AddProviders) Save(ctx context.Context, h host.Host, provideID int, apReqs []*AddProvidersSpan) error {
	log.Info("Saving add provider requests...")
	defer log.Info("Done saving add provider requests")

	localPeer, err := ap.peerService.UpsertLocalPeer(h)
	if err != nil {
		return err
	}

	for _, fnReq := range apReqs {

		remotePeer, err := ap.peerService.UpsertPeer(h, fnReq.RemotePeerID)
		if err != nil {
			return err
		}

		errStr := null.NewString("", false)
		if fnReq.Error != nil {
			errStr = null.StringFrom(fnReq.Error.Error())
		}

		maids := make([]int, len(fnReq.ProviderAddrs))
		for i, addr := range fnReq.ProviderAddrs {
			dbma, err := ap.maService.UpsertMultiAddress(ctx, addr)
			if err != nil {
				return errors.Wrap(err, "upsert multi address")
			}
			maids[i] = int(dbma.ID)
		}

		sort.Ints(maids)
		maids64 := make([]int64, len(maids))
		for i, x := range maids {
			maids64[i] = int64(x)
		}

		dbap := &models.AddProvider{
			ProvideID:       provideID,
			LocalID:         localPeer.ID,
			RemoteID:        remotePeer.ID,
			Distance:        ks.XORKeySpace.Key([]byte(fnReq.RemotePeerID)).Distance(ks.XORKeySpace.Key(fnReq.Content.CID.Hash())).Bytes(),
			MultiAddressIds: maids64,
			StartedAt:       fnReq.Start,
			EndedAt:         fnReq.End,
			Error:           errStr,
		}
		dbap, err = ap.apRepo.Save(ctx, dbap)
		if err != nil {
			return err
		}
	}

	return nil
}
