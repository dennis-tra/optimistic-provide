package repo

import (
	"context"

	"github.com/volatiletech/null/v8"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type DialRepo interface {
	Save(ctx context.Context, dial *models.Dial) (*models.Dial, error)
	ListFromProvide(ctx context.Context, provideID int) ([]*models.Dial, error)
	ListFromRetrieval(ctx context.Context, retrievalID int) ([]*models.Dial, error)
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

func (d Dial) ListFromProvide(ctx context.Context, provideID int) ([]*models.Dial, error) {
	return models.Dials(
		models.DialWhere.ProvideID.EQ(null.IntFrom(provideID)),
		qm.Load(models.DialRels.Remote),
		qm.Load(models.DialRels.MultiAddress),
	).All(ctx, d.dbc)
}

func (d Dial) ListFromRetrieval(ctx context.Context, retrievalID int) ([]*models.Dial, error) {
	return models.Dials(
		models.DialWhere.RetrievalID.EQ(null.IntFrom(retrievalID)),
		qm.Load(models.DialRels.Remote),
		qm.Load(models.DialRels.MultiAddress),
	).All(ctx, d.dbc)
}

func (d Dial) Save(ctx context.Context, dial *models.Dial) (*models.Dial, error) {
	return dial, dial.Insert(ctx, d.dbc, boil.Infer())
}
