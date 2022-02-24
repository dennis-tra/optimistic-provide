package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ProviderPeersRepo interface {
	Save(ctx context.Context, provider *models.ProviderPeer) (*models.ProviderPeer, error)
}

var _ ProviderPeersRepo = &ProviderPeers{}

type ProviderPeers struct {
	dbc *db.Client
}

func NewProviderPeersRepo(dbc *db.Client) ProviderPeersRepo {
	return &ProviderPeers{
		dbc: dbc,
	}
}

func (p ProviderPeers) Save(ctx context.Context, provider *models.ProviderPeer) (*models.ProviderPeer, error) {
	return provider, provider.Insert(ctx, p.dbc, boil.Infer())
}
