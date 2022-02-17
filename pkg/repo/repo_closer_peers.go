package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CloserPeersRepo interface {
	Save(ctx context.Context, cp *models.CloserPeer) (*models.CloserPeer, error)
}

var _ CloserPeersRepo = &CloserPeers{}

type CloserPeers struct {
	dbc *db.Client
}

func NewCloserPeersRepo(dbc *db.Client) CloserPeersRepo {
	return &CloserPeers{
		dbc: dbc,
	}
}

func (c CloserPeers) Save(ctx context.Context, cp *models.CloserPeer) (*models.CloserPeer, error) {
	return cp, cp.Insert(ctx, c.dbc, boil.Infer())
}
