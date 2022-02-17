package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PeerStateRepo interface {
	Save(ctx context.Context, dial *models.PeerState) (*models.PeerState, error)
}

var _ PeerStateRepo = &PeerState{}

type PeerState struct {
	dbc *db.Client
}

func NewPeerStateRepo(dbc *db.Client) PeerStateRepo {
	return &PeerState{
		dbc: dbc,
	}
}

func (fn PeerState) Save(ctx context.Context, dial *models.PeerState) (*models.PeerState, error) {
	return dial, dial.Insert(ctx, fn.dbc, boil.Infer())
}
