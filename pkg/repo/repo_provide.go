package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ProvideRepo interface {
	Save(ctx context.Context, provide *models.Provide) (*models.Provide, error)
	Update(ctx context.Context, provide *models.Provide) (*models.Provide, error)
}

var _ ProvideRepo = &Provide{}

type Provide struct {
	dbc *db.Client
}

func NewProvideRepo(dbc *db.Client) ProvideRepo {
	return &Provide{
		dbc: dbc,
	}
}

func (p Provide) Save(ctx context.Context, provide *models.Provide) (*models.Provide, error) {
	return provide, provide.Insert(ctx, p.dbc, boil.Infer())
}

func (p Provide) Update(ctx context.Context, provide *models.Provide) (*models.Provide, error) {
	_, err := provide.Update(ctx, p.dbc, boil.Infer())
	return provide, err
}
