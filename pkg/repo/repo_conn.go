package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ConnectionRepo interface {
	Save(ctx context.Context, conn *models.Connection) (*models.Connection, error)
}

var _ ConnectionRepo = &Connection{}

type Connection struct {
	dbc *db.Client
}

func NewConnectionRepo(dbc *db.Client) ConnectionRepo {
	return &Connection{
		dbc: dbc,
	}
}

func (c Connection) Save(ctx context.Context, conn *models.Connection) (*models.Connection, error) {
	return conn, conn.Insert(ctx, c.dbc, boil.Infer())
}
