package service

import (
	"context"
	"sort"

	lru "github.com/hashicorp/golang-lru"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type PeerService interface {
	Find(ctx context.Context, p peer.ID) (*models.Peer, error)
	UpsertLocalPeer(h host.Host) (*models.Peer, error)
	UpsertPeer(h host.Host, pid peer.ID) (*models.Peer, error)
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

func (ps *Peer) UpsertLocalPeer(h host.Host) (*models.Peer, error) {
	protocols := []string{}
	for _, prot := range kaddht.DefaultProtocols {
		protocols = append(protocols, string(prot))
	}
	return ps.repo.UpsertPeer(h.ID(), "optprov", protocols)
}

func (ps *Peer) UpsertPeer(h host.Host, pid peer.ID) (*models.Peer, error) {
	av := ""
	if agent, err := h.Peerstore().Get(pid, "AgentVersion"); err == nil {
		av = agent.(string)
	}

	protocols := []string{}
	if prots, err := h.Peerstore().GetProtocols(pid); err == nil {
		protocols = prots
	}

	cached, found := ps.cache.Get(pid.String())
	if found {
		cachedDbPeer := cached.(*models.Peer)
		sort.Strings(cachedDbPeer.Protocols)

		weHaveNewAgent := av != "" && cachedDbPeer.AgentVersion.String != av
		weHaveNewProtocols := false
		if len(cachedDbPeer.Protocols) != len(protocols) {
			weHaveNewProtocols = true
		} else {
			for i := 0; i < len(protocols); i++ {
				if cachedDbPeer.Protocols[i] != protocols[i] {
					weHaveNewProtocols = true
					break
				}
			}
		}
		if !weHaveNewAgent && !weHaveNewProtocols {
			return cachedDbPeer, nil
		}
	}

	dbPeer, err := ps.repo.UpsertPeer(pid, av, protocols)
	if err != nil {
		return nil, err
	}

	ps.cache.Add(dbPeer.MultiHash, dbPeer)

	return dbPeer, nil
}
