package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type GetProvidersRepo interface {
	List(ctx context.Context, retrieval *models.Retrieval) ([]*models.GetProvidersRPC, error)
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

func (gp GetProviders) List(ctx context.Context, retrieval *models.Retrieval) ([]*models.GetProvidersRPC, error) {
	return retrieval.GetProvidersRPCS(
		qm.Load(models.GetProvidersRPCRels.Remote),
	).All(ctx, gp.dbc)
}
