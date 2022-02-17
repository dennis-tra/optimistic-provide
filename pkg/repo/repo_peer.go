package repo

import (
	"context"
	"sort"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type PeerRepo interface {
	Find(context.Context, peer.ID) (*models.Peer, error)
	UpsertPeer(pid peer.ID, av string, protocols []string) (*models.Peer, error)
}

var _ PeerRepo = &Peer{}

type Peer struct {
	dbc *db.Client
}

func NewPeerRepo(dbc *db.Client) PeerRepo {
	return &Peer{
		dbc: dbc,
	}
}

func (p *Peer) Find(ctx context.Context, pid peer.ID) (*models.Peer, error) {
	return models.Peers(models.PeerWhere.MultiHash.EQ(pid.Pretty())).One(ctx, p.dbc)
}

func (p *Peer) UpsertPeer(pid peer.ID, av string, protocols []string) (*models.Peer, error) {
	sort.Strings(protocols)

	dbPeer := &models.Peer{
		MultiHash:    pid.String(),
		AgentVersion: null.NewString(av, av != ""),
		Protocols:    types.StringArray(protocols),
	}
	var prots interface{} = types.StringArray(protocols)
	if len(protocols) == 0 {
		prots = nil
	}

	rows, err := queries.Raw("SELECT upsert_peer($1, $2, $3)", dbPeer.MultiHash, dbPeer.AgentVersion.Ptr(), prots).Query(p.dbc)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, rows.Err()
	}
	if err = rows.Scan(&dbPeer.ID); err != nil {
		return nil, err
	}
	return dbPeer, rows.Close()
}
