package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/dennis-tra/optimistic-provide/pkg/repo"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
)

type HostService interface {
	Create(ctx context.Context, name string) (*dht.Host, error)
	Hosts(ctx context.Context) (map[string]*dht.Host, models.HostSlice, error)
	Host(ctx context.Context, p peer.ID) (*dht.Host, error)
	Start(ctx context.Context, h *dht.Host) (*dht.Host, error)
	Stop(p peer.ID) error
	Archive(ctx context.Context, dbHost *models.Host) error
}

var _ HostService = &Host{}

type Host struct {
	peerService PeerService
	rtService   RoutingTableService
	hostRepo    repo.HostRepo
	hostsLk     sync.RWMutex
	hosts       map[string]*dht.Host
}

func NewHostService(peerService PeerService, rtService RoutingTableService, hostRepo repo.HostRepo) HostService {
	return &Host{
		peerService: peerService,
		rtService:   rtService,
		hostRepo:    hostRepo,
		hostsLk:     sync.RWMutex{},
		hosts:       map[string]*dht.Host{},
	}
}

func (hs *Host) Create(ctx context.Context, name string) (*dht.Host, error) {
	key, _, err := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	if err != nil {
		return nil, errors.Wrap(err, "generate key pair")
	}

	keyDat, err := crypto.MarshalPrivateKey(key)
	if err != nil {
		return nil, errors.Wrap(err, "marshal private key")
	}

	h, err := dht.New(ctx, key)
	if err != nil {
		return nil, errors.Wrap(err, "new dht host")
	}

	hs.hostsLk.Lock()
	hs.hosts[h.Host.ID().String()] = h
	hs.hostsLk.Unlock()

	txn, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}
	defer deferTxRollback(txn)

	dbPeer, err := hs.peerService.UpsertLocalPeer(ctx, txn, h.Host)
	if err != nil {
		return nil, errors.Wrap(err, "upsert new libp2p host peer")
	}

	dbHost := &models.Host{
		Name:       name,
		PrivateKey: keyDat,
	}

	if err = dbHost.SetPeer(ctx, txn, false, dbPeer); err != nil {
		return nil, errors.Wrap(err, "set peer to db host")
	}

	if err = dbHost.Insert(ctx, txn, boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "insert host")
	}

	h.DBHost = dbHost

	if err = txn.Commit(); err != nil {
		return nil, errors.Wrap(err, "commit transaction")
	}

	return h, nil
}

func (hs *Host) Start(ctx context.Context, h *dht.Host) (*dht.Host, error) {
	hs.hostsLk.Lock()
	defer hs.hostsLk.Unlock()

	if h.StartedAt != nil {
		return h, nil
	}

	keyDat, err := crypto.UnmarshalPrivateKey(h.DBHost.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "marshal private key")
	}

	startedHost, err := dht.New(ctx, keyDat)
	if err != nil {
		return nil, errors.Wrap(err, "new dht host")
	}
	startedHost.DBHost = h.DBHost

	hs.hosts[startedHost.Host.ID().String()] = startedHost

	return startedHost, nil
}

func (hs *Host) Hosts(ctx context.Context) (map[string]*dht.Host, models.HostSlice, error) {
	hs.hostsLk.RLock()
	newMap := map[string]*dht.Host{}
	for k, v := range hs.hosts {
		newMap[k] = v
	}
	hs.hostsLk.RUnlock()

	dbHosts, err := hs.hostRepo.FindAllUnarchived(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting all hosts from DB")
	}

	return newMap, dbHosts, nil
}

func (hs *Host) Host(ctx context.Context, p peer.ID) (*dht.Host, error) {
	hs.hostsLk.RLock()
	h, ok := hs.hosts[p.String()]
	hs.hostsLk.RUnlock()
	if !ok {
		dbHost, err := hs.hostRepo.FindByPeerID(ctx, p)
		if err != nil {
			return nil, errors.Wrap(err, "find host with peer ID in DB "+p.String())
		}
		return &dht.Host{DBHost: dbHost}, nil
	}
	return h, nil
}

func (hs *Host) Stop(p peer.ID) error {
	hs.hostsLk.Lock()
	defer hs.hostsLk.Unlock()

	h, ok := hs.hosts[p.String()]
	if !ok {
		return fmt.Errorf("not found")
	}

	if err := h.Close(); err != nil {
		return err
	}

	delete(hs.hosts, p.String())

	return nil
}

func (hs *Host) Archive(ctx context.Context, dbHost *models.Host) error {
	hs.hostsLk.Lock()
	defer hs.hostsLk.Unlock()

	if err := hs.hostRepo.ArchiveHost(ctx, dbHost); err != nil {
		return errors.Wrap(err, "archive host")
	}

	for _, h := range hs.hosts {
		if h.DBHost.ID == dbHost.ID {
			if err := h.Close(); err != nil {
				return errors.Wrap(err, "close host")
			}

			delete(hs.hosts, h.ID().String())
		}
	}

	return nil
}
