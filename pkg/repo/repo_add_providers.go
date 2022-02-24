package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type AddProvidersRepo interface {
	Save(ctx context.Context, dial *models.AddProvider) (*models.AddProvider, error)
	List(ctx context.Context, provideID int) ([]*models.AddProvider, error)
}

var _ AddProvidersRepo = &AddProviders{}

type AddProviders struct {
	dbc *db.Client
}

func NewAddProvidersRepo(dbc *db.Client) AddProvidersRepo {
	return &AddProviders{
		dbc: dbc,
	}
}

func (fn AddProviders) List(ctx context.Context, provideID int) ([]*models.AddProvider, error) {
	return models.AddProviders(
		models.AddProviderWhere.ProvideID.EQ(provideID),
		qm.Load(models.AddProviderRels.Remote),
	).All(ctx, fn.dbc)
}

func (fn AddProviders) Save(ctx context.Context, addProvider *models.AddProvider) (*models.AddProvider, error) {
	return addProvider, addProvider.Insert(ctx, fn.dbc, boil.Infer())
}
