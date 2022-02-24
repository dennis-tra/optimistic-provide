package service

import (
	"context"
	"sort"
	"time"

	"github.com/libp2p/go-libp2p-core/host"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type ProvidersService interface {
	Save(ctx context.Context, h host.Host, retrievalID int, provider peer.AddrInfo) (*models.Provider, error)
}

var _ ProvidersService = &Providers{}

type Providers struct {
	peerService   PeerService
	maService     MultiAddressService
	providersRepo repo.ProvidersRepo
}

func NewProvidersService(peerService PeerService, maService MultiAddressService) *Providers {
	return &Providers{
		peerService: peerService,
		maService:   maService,
	}
}

func (p Providers) Save(ctx context.Context, h host.Host, retrievalID int, provider peer.AddrInfo) (*models.Provider, error) {
	foundAt := time.Now()
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

	dbProvider := &models.Provider{
		RetrievalID:     retrievalID,
		RemoteID:        dbPeer.ID,
		MultiAddressIds: dbMaddrIDsInt64,
		FoundAt:         foundAt,
	}

	return p.providersRepo.Save(ctx, dbProvider)
}
