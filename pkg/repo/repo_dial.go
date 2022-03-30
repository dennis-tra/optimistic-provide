package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type DialRepo interface {
	ListFromProvide(ctx context.Context, provide *models.Provide) ([]*models.Dial, error)
	ListFromRetrieval(ctx context.Context, retrieval *models.Retrieval) ([]*models.Dial, error)
	Save(ctx context.Context, exec boil.ContextExecutor, dial *models.Dial) (*models.Dial, error)
}

var _ DialRepo = &Dial{}

type Dial struct {
	dbc *db.Client
}

func NewDialRepo(dbc *db.Client) DialRepo {
	return &Dial{
		dbc: dbc,
	}
}

func (d Dial) ListFromProvide(ctx context.Context, provide *models.Provide) ([]*models.Dial, error) {
	return provide.Dials(
		qm.Load(models.DialRels.Remote),
		qm.Load(models.DialRels.MultiAddress),
	).All(ctx, d.dbc)
}

func (d Dial) ListFromRetrieval(ctx context.Context, retrieval *models.Retrieval) ([]*models.Dial, error) {
	return retrieval.Dials(
		qm.Load(models.DialRels.Remote),
		qm.Load(models.DialRels.MultiAddress),
	).All(ctx, d.dbc)
}

func (d Dial) Save(ctx context.Context, exec boil.ContextExecutor, dial *models.Dial) (*models.Dial, error) {
	return dial, dial.Insert(ctx, exec, boil.Infer())
}
