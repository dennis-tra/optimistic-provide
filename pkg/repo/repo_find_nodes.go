package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type FindNodesRPCRepo interface {
	List(ctx context.Context, provide *models.Provide) ([]*models.FindNodesRPC, error)
}

var _ FindNodesRPCRepo = &FindNodesRPC{}

type FindNodesRPC struct {
	dbc *db.Client
}

func NewFindNodesRPCRepo(dbc *db.Client) FindNodesRPCRepo {
	return &FindNodesRPC{
		dbc: dbc,
	}
}

func (fn FindNodesRPC) List(ctx context.Context, provide *models.Provide) ([]*models.FindNodesRPC, error) {
	return provide.FindNodesRPCS(
		qm.Load(models.FindNodesRPCRels.Remote),
	).All(ctx, fn.dbc.DB)
}
