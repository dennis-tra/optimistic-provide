package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type FindNodesRepo interface {
	Save(ctx context.Context, dial *models.FindNode) (*models.FindNode, error)
}

var _ FindNodesRepo = &FindNodes{}

type FindNodes struct {
	dbc *db.Client
}

func NewFindNodesRepo(dbc *db.Client) FindNodesRepo {
	return &FindNodes{
		dbc: dbc,
	}
}

func (fn FindNodes) Save(ctx context.Context, dial *models.FindNode) (*models.FindNode, error) {
	return dial, dial.Insert(ctx, fn.dbc, boil.Infer())
}
