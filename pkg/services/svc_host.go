package services

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
)

type HostService struct {
	dbc     *db.Client
	hostsLk sync.RWMutex
	hosts   map[string]*dht.Host
}

func NewHostService(dbc *db.Client) *HostService {
	return &HostService{
		dbc:   dbc,
		hosts: map[string]*dht.Host{},
	}
}

func (hs *HostService) Create(ctx context.Context) (*dht.Host, error) {
	hs.hostsLk.Lock()
	defer hs.hostsLk.Unlock()

	h, err := dht.New(ctx)
	if err != nil {
		return nil, err
	}

	hs.hosts[h.Host.ID().String()] = h

	dbPeer, err := hs.dbc.UpsertLocalPeer(h.Host.ID())
	if err != nil{
		return nil , err
	}
	h.DBPeer = dbPeer

	return h, err
}

func (hs *HostService) Hosts() map[string]*dht.Host {
	hs.hostsLk.RLock()
	defer hs.hostsLk.RUnlock()

	newMap := map[string]*dht.Host{}
	for k, v := range hs.hosts {
		newMap[k] = v
	}

	return newMap
}

func (hs *HostService) Host(p peer.ID) (*dht.Host, bool) {
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