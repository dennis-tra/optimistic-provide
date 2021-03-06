package service

import (
	"context"
	"sort"

	lru "github.com/hashicorp/golang-lru"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type PeerService interface {
	Find(ctx context.Context, p peer.ID) (*models.Peer, error)
	UpsertLocalPeer(ctx context.Context, exec boil.ContextExecutor, h host.Host) (*models.Peer, error)
	UpsertLocalPeerTxn(ctx context.Context, h host.Host) (*models.Peer, error)
	UpsertPeer(ctx context.Context, exec boil.ContextExecutor, h host.Host, pid peer.ID) (*models.Peer, error)
	UpsertPeerForInfo(ctx context.Context, exec boil.ContextExecutor, h host.Host, pid peer.ID, pi *PeerInfo) (*models.Peer, error)
}

var _ PeerService = &Peer{}

type Peer struct {
	repo  repo.PeerRepo
	cache *lru.Cache
}

func (ps *Peer) Find(ctx context.Context, p peer.ID) (*models.Peer, error) {
	return ps.repo.Find(ctx, p)
}

func NewPeerService(repo repo.PeerRepo) PeerService {
	cache, err := lru.New(1000)
	if err != nil {
		panic(err)
	}

	return &Peer{
		repo:  repo,
		cache: cache,
	}
}

func (ps *Peer) UpsertLocalPeerTxn(ctx context.Context, h host.Host) (*models.Peer, error) {
	txn, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}

	snapshot, err := ps.UpsertLocalPeer(ctx, txn, h)
	if err != nil {
		if err2 := txn.Rollback(); err != nil {
			log.Warn("error rolling back transaction", err2)
		}
		return nil, err
	}
	if err := txn.Commit(); err != nil {
		if err2 := txn.Rollback(); err != nil {
			log.Warn("error rolling back transaction", err2)
		}
		return nil, errors.Wrap(err, "committing transaction")
	}

	return snapshot, nil
}

func (ps *Peer) UpsertLocalPeer(ctx context.Context, exec boil.ContextExecutor, h host.Host) (*models.Peer, error) {
	protocols := []string{}
	for _, prot := range kaddht.DefaultProtocols {
		protocols = append(protocols, string(prot))
	}
	return ps.repo.UpsertPeer(ctx, exec, h.ID(), "optprov", protocols)
}

func (ps *Peer) UpsertPeer(ctx context.Context, exec boil.ContextExecutor, h host.Host, pid peer.ID) (*models.Peer, error) {
	av := ""
	if agent, err := h.Peerstore().Get(pid, "AgentVersion"); err == nil {
		av = agent.(string)
	}

	protocols := []string{}
	if prots, err := h.Peerstore().GetProtocols(pid); err == nil {
		protocols = prots
	}

	if cachedDbPeer := ps.isCached(pid, av, protocols); cachedDbPeer != nil {
		return cachedDbPeer, nil
	}

	dbPeer, err := ps.repo.UpsertPeer(ctx, exec, pid, av, protocols)
	if err != nil {
		return nil, err
	}

	ps.cache.Add(dbPeer.MultiHash, dbPeer)

	return dbPeer, nil
}

func (ps *Peer) UpsertPeerForInfo(ctx context.Context, exec boil.ContextExecutor, h host.Host, pid peer.ID, pi *PeerInfo) (*models.Peer, error) {
	if pi == nil {
		return ps.UpsertPeer(ctx, exec, h, pid)
	}

	pi.SetFromPeerstore(h.Peerstore())

	if cachedDbPeer := ps.isCached(pi.PeerID, pi.AgentVersion, pi.Protocols); cachedDbPeer != nil {
		return cachedDbPeer, nil
	}

	dbPeer, err := ps.repo.UpsertPeer(ctx, exec, pi.PeerID, pi.AgentVersion, pi.Protocols)
	if err != nil {
		return nil, err
	}

	ps.cache.Add(dbPeer.MultiHash, dbPeer)

	return dbPeer, nil
}

func (ps *Peer) isCached(pid peer.ID, av string, protocols []string) *models.Peer {
	cached, found := ps.cache.Get(pid.String())
	if !found {
		return nil
	}

	cachedDbPeer, ok := cached.(*models.Peer)
	if !ok {
		ps.cache.Remove(pid.String())
		return nil
	}

	if av != "" && cachedDbPeer.AgentVersion.String != av {
		return nil
	}

	if len(protocols) != 0 && len(cachedDbPeer.Protocols) != len(protocols) {
		return nil
	}

	sort.Strings(cachedDbPeer.Protocols)
	sort.Strings(protocols)

	for i := 0; i < len(protocols); i++ {
		if cachedDbPeer.Protocols[i] != protocols[i] {
			return nil
		}
	}

	return cachedDbPeer
}
