package service

import (
	"context"
	"sort"

	"github.com/volatiletech/sqlboiler/v4/boil"

	lru "github.com/hashicorp/golang-lru"

	"github.com/dennis-tra/optimistic-provide/pkg/maxmind"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/pkg/errors"
)

type MultiAddressService interface {
	UpsertMultiAddress(ctx context.Context, exec boil.ContextExecutor, maddr ma.Multiaddr) (*models.MultiAddress, error)
	UpsertMultiAddresses(ctx context.Context, exec boil.ContextExecutor, maddr []ma.Multiaddr) ([]int64, error)
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

	cache, err := lru.New(10000)
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

func (ma *MultiAddress) UpsertMultiAddress(ctx context.Context, exec boil.ContextExecutor, maddr ma.Multiaddr) (*models.MultiAddress, error) {
	cached, found := ma.cache.Get(maddr.String())
	if found {
		return cached.(*models.MultiAddress), nil
	}

	isPublic := manet.IsPublicAddr(maddr)

	infos, err := ma.mmclient.MaddrInfo(ctx, maddr)
	if err != nil {
		dbMaddr := &models.MultiAddress{
			Maddr: maddr.String(),
		}
		return ma.maRepo.UpsertMultiAddress(ctx, exec, dbMaddr, nil, isPublic)
	}

	ipAddresses := []*models.IPAddress{}
	for address, info := range infos {
		dbIPAddress, err := ma.iaRepo.UpsertIPAddress(ctx, exec, address, info, isPublic)
		if err != nil {
			return nil, errors.Wrap(err, "upsert ip address")
		}
		ipAddresses = append(ipAddresses, dbIPAddress)
	}

	dbMaddr := &models.MultiAddress{
		Maddr: maddr.String(),
	}

	dbMaddr, err = ma.maRepo.UpsertMultiAddress(ctx, exec, dbMaddr, ipAddresses, isPublic)
	if err != nil {
		return nil, err
	}

	ma.cache.Add(maddr.String(), dbMaddr)

	return dbMaddr, err
}

func (ma *MultiAddress) UpsertMultiAddresses(ctx context.Context, exec boil.ContextExecutor, maddrs []ma.Multiaddr) ([]int64, error) {
	maids := make([]int, len(maddrs))
	for i, addr := range maddrs {
		dbma, err := ma.UpsertMultiAddress(ctx, exec, addr)
		if err != nil {
			return nil, errors.Wrap(err, "upsert multi address")
		}
		maids[i] = int(dbma.ID)
	}

	sort.Ints(maids)
	maids64 := make([]int64, len(maids))
	for i, x := range maids {
		maids64[i] = int64(x)
	}

	return maids64, nil
}
