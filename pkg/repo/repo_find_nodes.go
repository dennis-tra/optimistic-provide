package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type FindNodesRepo interface {
	Save(ctx context.Context, dial *models.FindNode) (*models.FindNode, error)
	List(ctx context.Context, provideID int) ([]*models.FindNode, error)
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

func (fn FindNodes) List(ctx context.Context, provideID int) ([]*models.FindNode, error) {
	return models.FindNodes(
		models.FindNodeWhere.ProvideID.EQ(provideID),
		qm.Load(models.FindNodeRels.Remote),
	).All(ctx, fn.dbc)
}

func (fn FindNodes) Save(ctx context.Context, dial *models.FindNode) (*models.FindNode, error) {
	return dial, dial.Insert(ctx, fn.dbc, boil.Infer())
}
