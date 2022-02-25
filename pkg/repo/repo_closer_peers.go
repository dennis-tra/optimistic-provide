package repo

import (
	"github.com/dennis-tra/optimistic-provide/pkg/db"
)

type CloserPeersRepo interface{}

var _ CloserPeersRepo = &CloserPeers{}

type CloserPeers struct {
	dbc *db.Client
}

func NewCloserPeersRepo(dbc *db.Client) CloserPeersRepo {
	return &CloserPeers{
		dbc: dbc,
	}
}
