package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type DialRepo interface {
	Save(ctx context.Context, dial *models.Dial) (*models.Dial, error)
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

func (d Dial) Save(ctx context.Context, dial *models.Dial) (*models.Dial, error) {
	return dial, dial.Insert(ctx, d.dbc, boil.Infer())
}
