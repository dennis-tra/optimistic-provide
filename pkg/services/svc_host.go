package services

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/host"
)

type HostService struct {
	dbc     *db.Client
	hostsLk sync.RWMutex
	hosts   map[string]*host.Host
}

func NewHostService(dbc *db.Client) *HostService {
	return &HostService{
		dbc:   dbc,
		hosts: map[string]*host.Host{},
	}
}

func (hs *HostService) Create(ctx context.Context) (*host.Host, error) {
	hs.hostsLk.Lock()
	defer hs.hostsLk.Unlock()

	h, err := host.New(ctx)
	if err != nil {
		return nil, err
	}

	hs.hosts[h.Host.ID().String()] = h

	return h, err
}

func (hs *HostService) Hosts() map[string]*host.Host {
	hs.hostsLk.RLock()
	defer hs.hostsLk.RUnlock()

	return hs.hosts
}

func (hs *HostService) Host(p peer.ID) (*host.Host, bool) {
	hs.hostsLk.RLock()
	defer hs.hostsLk.RUnlock()

	h, ok := hs.hosts[p.String()]

	return h, ok
}

func (hs *HostService) Stop(p peer.ID) error {
	hs.hostsLk.Lock()
	defer hs.hostsLk.Unlock()

	h, ok := hs.hosts[p.String()]
	if !ok {
		return fmt.Errorf("not found")
	}

	if err := h.Host.Close(); err != nil {
		return err
	}

	delete(hs.hosts, p.String())

	return nil
}
