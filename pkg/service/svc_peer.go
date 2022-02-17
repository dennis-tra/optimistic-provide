package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
)

type PeerService interface {
	Find(ctx context.Context, p peer.ID) (*models.Peer, error)
	UpsertLocalPeer(h host.Host) (*models.Peer, error)
	UpsertPeer(h host.Host, pid peer.ID) (*models.Peer, error)
}

var _ PeerService = &Peer{}

type Peer struct {
	repo repo.PeerRepo
}

func (ps *Peer) Find(ctx context.Context, p peer.ID) (*models.Peer, error) {
	return ps.repo.Find(ctx, p)
}

func NewPeerService(repo repo.PeerRepo) PeerService {
	return &Peer{
		repo: repo,
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

	return ps.repo.UpsertPeer(pid, av, protocols)
}
