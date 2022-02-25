package service

import (
	"context"
	"sort"

	"github.com/libp2p/go-libp2p-core/host"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type ProviderPeersService interface {
	Save(ctx context.Context, h host.Host, getProviderID int, provider peer.AddrInfo) (*models.ProviderPeer, error)
}

var _ ProviderPeersService = &Providers{}

type Providers struct {
	peerService   PeerService
	maService     MultiAddressService
	providersRepo repo.ProviderPeersRepo
}

func NewProviderPeersService(peerService PeerService, maService MultiAddressService, providersRepo repo.ProviderPeersRepo) *Providers {
	return &Providers{
		peerService:   peerService,
		maService:     maService,
		providersRepo: providersRepo,
	}
}

func (p Providers) Save(ctx context.Context, h host.Host, getProvidersID int, provider peer.AddrInfo) (*models.ProviderPeer, error) {
	dbPeer, err := p.peerService.UpsertPeer(h, provider.ID)
	if err != nil {
		return nil, err
	}

	dbMaddrIDs := make([]int, len(provider.Addrs))
	for i, maddr := range provider.Addrs {
		dbMaddr, err := p.maService.UpsertMultiAddress(ctx, maddr)
		if err != nil {
			return nil, err
		}
		dbMaddrIDs[i] = dbMaddr.ID
	}

	sort.Ints(dbMaddrIDs)

	dbMaddrIDsInt64 := make([]int64, len(dbMaddrIDs))
	for i, id := range dbMaddrIDs {
		dbMaddrIDsInt64[i] = int64(id)
	}

	dbProvider := &models.ProviderPeer{
		GetProvidersID:  getProvidersID,
		ProviderID:      dbPeer.ID,
		MultiAddressIds: dbMaddrIDsInt64,
	}

	return p.providersRepo.Save(ctx, dbProvider)
}
