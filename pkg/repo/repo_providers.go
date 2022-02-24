package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ProvidersRepo interface {
	Save(ctx context.Context, provider *models.Provider) (*models.Provider, error)
}

var _ ProvidersRepo = &Providers{}

type Providers struct {
	dbc *db.Client
}

func NewProvidersRepo(dbc *db.Client) ProvidersRepo {
	return &Providers{
		dbc: dbc,
	}
}

func (p Providers) Save(ctx context.Context, provider *models.Provider) (*models.Provider, error) {
	return provider, provider.Insert(ctx, p.dbc, boil.Infer())
}
