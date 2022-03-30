package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type AddProvidersRepo interface {
	List(ctx context.Context, provide *models.Provide) ([]*models.AddProviderRPC, error)
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

func (ap AddProviders) List(ctx context.Context, provide *models.Provide) ([]*models.AddProviderRPC, error) {
	return provide.AddProviderRPCS(qm.Load(models.AddProviderRPCRels.Remote)).All(ctx, ap.dbc)
}
