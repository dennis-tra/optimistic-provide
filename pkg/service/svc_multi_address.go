package service

import (
	"context"

	lru "github.com/hashicorp/golang-lru"

	"github.com/dennis-tra/optimistic-provide/pkg/maxmind"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/pkg/errors"
)

type MultiAddressService interface {
	UpsertMultiAddress(ctx context.Context, maddr ma.Multiaddr) (*models.MultiAddress, error)
}

var _ MultiAddressService = &MultiAddress{}

type MultiAddress struct {
	mmclient *maxmind.Client
	cache    *lru.Cache
	iaRepo   repo.IPAddressRepo
	maRepo   repo.MultiAddressRepo
}

func NewMultiAddressService(maRepo repo.MultiAddressRepo, iaRepo repo.IPAddressRepo) *MultiAddress {
	mmclient, err := maxmind.NewClient()
	if err != nil {
		panic(err)
	}

	cache, err := lru.New(1000)
	if err != nil {
		panic(err)
	}

	return &MultiAddress{
		mmclient: mmclient,
		cache:    cache,
		iaRepo:   iaRepo,
		maRepo:   maRepo,
	}
}

func (ma *MultiAddress) UpsertMultiAddress(ctx context.Context, maddr ma.Multiaddr) (*models.MultiAddress, error) {
	isPublic := manet.IsPublicAddr(maddr)

	infos, err := ma.mmclient.MaddrInfo(ctx, maddr)
	if err != nil {
		return nil, errors.Wrap(err, "resolve maddr infos")
	}

	ipAddresses := []*models.IPAddress{}
	for address, info := range infos {
		dbIPAddress, err := ma.iaRepo.UpsertIPAddress(ctx, address, info, isPublic)
		if err != nil {
			return nil, errors.Wrap(err, "upsert ip address")
		}
		ipAddresses = append(ipAddresses, dbIPAddress)
	}

	dbMaddr := &models.MultiAddress{
		Maddr: maddr.String(),
	}

	return ma.maRepo.UpsertMultiAddress(ctx, dbMaddr, ipAddresses, isPublic)
}
