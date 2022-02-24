package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type ConnectionRepo interface {
	Save(ctx context.Context, conn *models.Connection) (*models.Connection, error)
	List(ctx context.Context, provideID int) ([]*models.Connection, error)
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

func (c Connection) List(ctx context.Context, provideID int) ([]*models.Connection, error) {
	return models.Connections(
		models.ConnectionWhere.ProvideID.EQ(provideID),
		qm.Load(models.ConnectionRels.Remote),
		qm.Load(models.ConnectionRels.MultiAddress),
	).All(ctx, c.dbc)
}

func (c Connection) Save(ctx context.Context, conn *models.Connection) (*models.Connection, error) {
	return conn, conn.Insert(ctx, c.dbc, boil.Infer())
}
