package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/dennis-tra/optimistic-provide/pkg/models"

	"github.com/pkg/errors"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
)

type HostService interface {
	Create(ctx context.Context) (*dht.Host, error)
	Hosts() map[string]*dht.Host
	Host(p peer.ID) (*dht.Host, bool)
	Stop(p peer.ID) error
	RefreshRoutingTableAsync(ctx context.Context, p peer.ID) error
	SaveRoutingTable(ctx context.Context, p peer.ID) (*models.RoutingTableSnapshot, error)
}

var _ HostService = &Host{}

type Host struct {
	peerService PeerService
	rtService   RoutingTableService
	hostsLk     sync.RWMutex
	hosts       map[string]*dht.Host
}

func NewHostService(peerService PeerService, rtService RoutingTableService) HostService {
	return &Host{
		peerService: peerService,
		rtService:   rtService,
		hostsLk:     sync.RWMutex{},
		hosts:       map[string]*dht.Host{},
	}
}

func (hs *Host) Create(ctx context.Context) (*dht.Host, error) {
	hs.hostsLk.Lock()
	defer hs.hostsLk.Unlock()

	h, err := dht.New(ctx)
	if err != nil {
		return nil, err
	}

	hs.hosts[h.Host.ID().String()] = h

	dbPeer, err := hs.peerService.UpsertLocalPeer(h.Host)
	if err != nil {
		return nil, err
	}
	h.DBPeer = dbPeer

	return h, err
}

func (hs *Host) Hosts() map[string]*dht.Host {
	hs.hostsLk.RLock()
	defer hs.hostsLk.RUnlock()

	newMap := map[string]*dht.Host{}
	for k, v := range hs.hosts {
		newMap[k] = v
	}

	return newMap
}

func (hs *Host) Host(p peer.ID) (*dht.Host, bool) {
	hs.hostsLk.RLock()
	defer hs.hostsLk.RUnlock()

	h, ok := hs.hosts[p.String()]

	return h, ok
}

func (hs *Host) Stop(p peer.ID) error {
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

func (hs *Host) RefreshRoutingTableAsync(ctx context.Context, p peer.ID) error {
	h, found := hs.Host(p)
	if !found {
		return errors.New("peer not found")
	}
	go h.RefreshRoutingTable(ctx)

	return nil
}

func (hs *Host) SaveRoutingTable(ctx context.Context, p peer.ID) (*models.RoutingTableSnapshot, error) {
	h, found := hs.Host(p)
	if !found {
		return nil, errors.New("peer not found")
	}

	return hs.rtService.SaveRoutingTable(ctx, h)
}
