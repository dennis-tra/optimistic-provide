package repo

import (
	"github.com/dennis-tra/optimistic-provide/pkg/db"
)

type PeerStateRepo interface{}

var _ PeerStateRepo = &PeerState{}

type PeerState struct {
	dbc *db.Client
}

func NewPeerStateRepo(dbc *db.Client) PeerStateRepo {
	return &PeerState{
		dbc: dbc,
	}
}
