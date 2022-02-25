package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type ConnectionRepo interface {
	ListFromProvide(ctx context.Context, provide *models.Provide) ([]*models.Connection, error)
	ListFromRetrieval(ctx context.Context, retrieval *models.Retrieval) ([]*models.Connection, error)
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

func (c Connection) ListFromProvide(ctx context.Context, provide *models.Provide) ([]*models.Connection, error) {
	return provide.Connections(
		qm.Load(models.ConnectionRels.Remote),
		qm.Load(models.ConnectionRels.MultiAddress),
	).All(ctx, c.dbc)
}

func (c Connection) ListFromRetrieval(ctx context.Context, retrieval *models.Retrieval) ([]*models.Connection, error) {
	return retrieval.Connections(
		qm.Load(models.ConnectionRels.Remote),
		qm.Load(models.ConnectionRels.MultiAddress),
	).All(ctx, c.dbc)
}
