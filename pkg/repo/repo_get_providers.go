package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type GetProvidersRepo interface {
	Save(ctx context.Context, dial *models.GetProvider) (*models.GetProvider, error)
	List(ctx context.Context, provideID int) ([]*models.GetProvider, error)
}

var _ GetProvidersRepo = &GetProviders{}

type GetProviders struct {
	dbc *db.Client
}

func NewGetProvidersRepo(dbc *db.Client) GetProvidersRepo {
	return &GetProviders{
		dbc: dbc,
	}
}

func (fn GetProviders) List(ctx context.Context, retrievalID int) ([]*models.GetProvider, error) {
	return models.GetProviders(
		models.GetProviderWhere.RetrievalID.EQ(retrievalID),
		qm.Load(models.GetProviderRels.Remote),
	).All(ctx, fn.dbc)
}

func (fn GetProviders) Save(ctx context.Context, dial *models.GetProvider) (*models.GetProvider, error) {
	return dial, dial.Insert(ctx, fn.dbc, boil.Infer())
}
