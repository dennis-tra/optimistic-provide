package dht

import (
	"context"
	"time"

	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/pkg/errors"

	"github.com/dennis-tra/optimistic-provide/pkg/lib"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/wrap"
)

var log = logging.Logger("optprov")

type Host struct {
	DBPeer       *models.Peer
	PeerID       peer.ID
	Host         host.Host
	DHT          *kaddht.IpfsDHT
	Bootstrapped *time.Time
	CreatedAt    time.Time
	Transports   []*wrap.Notifier
	MsgSender    *wrap.MessageSenderImpl

	RoutingTableRefresh RoutingTableRefresh
}

type RoutingTableRefresh struct {
	StartedAt *time.Time
	EndedAt   *time.Time
	Error     *string
}

func New(ctx context.Context) (*Host, error) {
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
		PeerID:     h.ID(),
		Host:       h,
		DHT:        dht,
		MsgSender:  msgSender,
		CreatedAt:  time.Now(),
		Transports: []*wrap.Notifier{tcp.Notifier, ws.Notifier, quic.Notifier},
	}

	log.Infow("Initialized new libp2p host", "localID", lib.FmtPeerID(h.ID()))
	return newHost, nil
}

func (h *Host) RoutingTableRefreshDuration() *time.Duration {
	if h.RoutingTableRefresh.StartedAt == nil || h.RoutingTableRefresh.EndedAt == nil {
		return nil
	}
	dur := h.RoutingTableRefresh.EndedAt.Sub(*h.RoutingTableRefresh.StartedAt)
	return &dur
}

func (h *Host) Bootstrap(ctx context.Context) error {
	for _, bp := range kaddht.GetDefaultBootstrapPeerAddrInfos() {
		log.Infow("Connecting to bootstrap peer", "remoteID", lib.FmtPeerID(bp.ID))
		if err := h.Host.Connect(ctx, bp); err != nil {
			return errors.Wrap(err, "connecting to bootstrap peer")
		}
	}
	h.Bootstrapped = lib.NowPtr()
	return nil
}

func (h *Host) RefreshRoutingTable(ctx context.Context) {
	log.Infow("Start refreshing routing table")
	defer log.Infow("Done refreshing routing table")

	h.RoutingTableRefresh.StartedAt = lib.NowPtr()
	defer func() { h.RoutingTableRefresh.EndedAt = lib.NowPtr() }()

	h.RoutingTableRefresh.Error = nil
	h.RoutingTableRefresh.EndedAt = nil

	select {
	case err := <-h.DHT.RefreshRoutingTable():
		if err != nil {
			h.RoutingTableRefresh.Error = lib.StrPtr(err.Error())
		}
	case <-ctx.Done():
		if ctx.Err() != nil {
			h.RoutingTableRefresh.Error = lib.StrPtr(ctx.Err().Error())
		}
	}
}
