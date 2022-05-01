package dht

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/libp2p/go-libp2p-core/peer"

	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/pkg/errors"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
	"github.com/dennis-tra/optimistic-provide/pkg/wrap"
)

var log = logging.Logger("optprov")

type RoutingTableListener interface {
	PeerAdded(p peer.ID)
	PeerRemoved(p peer.ID)
	OnClose()
}

type Host struct {
	host.Host

	DBHost       *models.Host
	DHT          *kaddht.IpfsDHT
	Bootstrapped *time.Time
	StartedAt    *time.Time
	Transports   []*wrap.Notifier
	MsgSender    *wrap.MessageSenderImpl

	rtPeerAdded   func(peer.ID)
	rtPeerRemoved func(peer.ID)
	rtListenerslk sync.RWMutex
	rtListeners   map[RoutingTableListener]*sync.WaitGroup
}

func New(ctx context.Context, key crypto.PrivKey) (*Host, error) {
	tcp, tcpTrpt := wrap.NewTCPTransport()
	ws, wsTrpt := wrap.NewWSTransport()
	quic, quicTrpt := wrap.NewQuicTransport()
	msgSender := wrap.NewMessageSenderImpl()

	newHost := &Host{
		MsgSender:   msgSender,
		Transports:  []*wrap.Notifier{tcp.Notifier, ws.Notifier, quic.Notifier},
		rtListeners: map[RoutingTableListener]*sync.WaitGroup{},
	}

	var dht *kaddht.IpfsDHT
	h, err := libp2p.New(
		libp2p.DefaultListenAddrs,
		libp2p.Identity(key),
		libp2p.Transport(tcpTrpt),
		libp2p.Transport(wsTrpt),
		libp2p.Transport(quicTrpt),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			var err error
			dht, err = kaddht.New(ctx, h,
				kaddht.Mode(kaddht.ModeClient),
				kaddht.MessageSenderImpl(msgSender.Init),
				kaddht.NetworkSizeHook(newHost.SaveNetworkSizeEstimate),
			)
			return dht, err
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host")
	}

	now := time.Now()
	newHost.Host = h
	newHost.DHT = dht
	newHost.StartedAt = &now
	newHost.rtPeerAdded = dht.RoutingTable().PeerAdded
	newHost.rtPeerRemoved = dht.RoutingTable().PeerRemoved

	dht.RoutingTable().PeerAdded = newHost.peerAdded
	dht.RoutingTable().PeerRemoved = newHost.peerRemoved

	log.Infow("Initialized new libp2p host", "localID", util.FmtPeerID(h.ID()))
	return newHost, nil
}

func (h *Host) SaveNetworkSizeEstimate(avg float64, std float64, r2 float64, i int) {
	estimate := models.NetworkSizeEstimate{
		PeerID:         h.DBHost.ID,
		NetworkSize:    avg,
		NetworkSizeErr: std,
		RSquared:       r2,
		Extra:          null.StringFrom("linear-regression"),
	}
	if err := estimate.Insert(context.Background(), boil.GetContextDB(), boil.Infer()); err != nil {
		fmt.Printf("Error inserting network size estimate: %s", err)
	}
}

func (h *Host) PeerID() string {
	return h.DBHost.R.Peer.MultiHash
}

func (h *Host) Close() error {
	h.rtListenerslk.Lock()
	for listener := range h.rtListeners {
		wg, ok := h.rtListeners[listener]
		if !ok {
			continue
		}
		// Wait for all in-flight go-routines to finish before returning
		wg.Wait()
		listener.OnClose()
		delete(h.rtListeners, listener)
	}
	h.rtListenerslk.Unlock()
	return h.Host.Close()
}

func (h *Host) Bootstrap(ctx context.Context) error {
	for _, bp := range kaddht.GetDefaultBootstrapPeerAddrInfos() {
		log.Infow("Connecting to bootstrap peer", "remoteID", util.FmtPeerID(bp.ID))
		if err := h.Host.Connect(ctx, bp); err != nil {
			return errors.Wrap(err, "connecting to bootstrap peer")
		}
	}
	h.Bootstrapped = util.NowPtr()
	return nil
}

func (h *Host) IsRoutingTableListenerRegistered(listener RoutingTableListener) bool {
	h.rtListenerslk.RLock()
	defer h.rtListenerslk.RUnlock()
	_, ok := h.rtListeners[listener]
	return ok
}

func (h *Host) RegisterRoutingTableListener(listener RoutingTableListener) {
	h.rtListenerslk.Lock()
	defer h.rtListenerslk.Unlock()
	h.rtListeners[listener] = &sync.WaitGroup{}
}

func (h *Host) UnregisterRoutingTableListener(listener RoutingTableListener) {
	h.rtListenerslk.Lock()
	defer h.rtListenerslk.Unlock()
	wg, ok := h.rtListeners[listener]
	if !ok {
		return
	}
	// Wait for all in-flight go-routines to finish before returning
	defer wg.Wait()

	delete(h.rtListeners, listener)
}

func (h *Host) peerAdded(p peer.ID) {
	h.rtPeerAdded(p)

	h.rtListenerslk.RLock()
	defer h.rtListenerslk.RUnlock()

	for listener, wg := range h.rtListeners {
		wgCpy := wg
		listenerCpy := listener
		wg.Add(1)
		go func() {
			listenerCpy.PeerAdded(p)
			wgCpy.Done()
		}()
	}
}

func (h *Host) peerRemoved(p peer.ID) {
	h.rtPeerRemoved(p)

	h.rtListenerslk.RLock()
	defer h.rtListenerslk.RUnlock()

	for listener, wg := range h.rtListeners {
		wgCpy := wg
		listenerCpy := listener
		wg.Add(1)
		go func() {
			listenerCpy.PeerRemoved(p)
			wgCpy.Done()
		}()
	}
}
