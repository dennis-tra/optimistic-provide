package dht

import (
	"context"
	"sync"
	"time"

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

	Name         string
	DBPeer       *models.Peer
	DHT          *kaddht.IpfsDHT
	Bootstrapped *time.Time
	CreatedAt    time.Time
	Transports   []*wrap.Notifier
	MsgSender    *wrap.MessageSenderImpl

	rtPeerAdded   func(peer.ID)
	rtPeerRemoved func(peer.ID)
	rtListenerslk sync.RWMutex
	rtListeners   map[RoutingTableListener]struct{}
}

func New(ctx context.Context, name string) (*Host, error) {
	key, _, err := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	if err != nil {
		return nil, errors.Wrap(err, "generate key pair")
	}

	tcp, tcpTrpt := wrap.NewTCPTransport()
	ws, wsTrpt := wrap.NewWSTransport()
	quic, quicTrpt := wrap.NewQuicTransport()
	msgSender := wrap.NewMessageSenderImpl()

	var dht *kaddht.IpfsDHT
	h, err := libp2p.New(
		libp2p.Identity(key),
		libp2p.Transport(tcpTrpt),
		libp2p.Transport(wsTrpt),
		libp2p.Transport(quicTrpt),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			dht, err = kaddht.New(ctx, h,
				kaddht.Mode(kaddht.ModeClient),
				kaddht.MessageSenderImpl(msgSender.Init),
			)
			return dht, err
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host")
	}

	newHost := &Host{
		Name:          name,
		Host:          h,
		DHT:           dht,
		MsgSender:     msgSender,
		CreatedAt:     time.Now(),
		Transports:    []*wrap.Notifier{tcp.Notifier, ws.Notifier, quic.Notifier},
		rtListeners:   map[RoutingTableListener]struct{}{},
		rtPeerAdded:   dht.RoutingTable().PeerAdded,
		rtPeerRemoved: dht.RoutingTable().PeerRemoved,
	}

	dht.RoutingTable().PeerAdded = newHost.peerAdded
	dht.RoutingTable().PeerRemoved = newHost.peerRemoved

	log.Infow("Initialized new libp2p host", "localID", util.FmtPeerID(h.ID()))
	return newHost, nil
}

func (h *Host) Close() error {
	h.rtListenerslk.RLock()
	for listener := range h.rtListeners {
		go listener.OnClose()
	}
	h.rtListenerslk.RUnlock()
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

func (h *Host) RegisterRoutingTableListener(listener RoutingTableListener) {
	h.rtListenerslk.Lock()
	defer h.rtListenerslk.Unlock()
	h.rtListeners[listener] = struct{}{}
}

func (h *Host) UnregisterRoutingTableListener(listener RoutingTableListener) {
	h.rtListenerslk.Lock()
	defer h.rtListenerslk.Unlock()
	delete(h.rtListeners, listener)
}

func (h *Host) peerAdded(p peer.ID) {
	h.rtListenerslk.RLock()
	for listener := range h.rtListeners {
		go listener.PeerAdded(p)
	}
	h.rtListenerslk.RUnlock()
	h.rtPeerAdded(p)
}

func (h *Host) peerRemoved(p peer.ID) {
	h.rtListenerslk.RLock()
	for listener := range h.rtListeners {
		go listener.PeerRemoved(p)
	}
	h.rtListenerslk.RUnlock()
	h.rtPeerRemoved(p)
}
