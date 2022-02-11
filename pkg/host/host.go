package host

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
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	"github.com/pkg/errors"

	"github.com/dennis-tra/optimistic-provide/pkg/utils"
	"github.com/dennis-tra/optimistic-provide/pkg/wrap"
)

var log = logging.Logger("optprov")

type Host struct {
	dbID         int             `json:"-"`
	PeerID       peer.ID         `json:"peerId"`
	Host         host.Host       `json:"-"`
	DHT          *kaddht.IpfsDHT `json:"-"`
	Bootstrapped *time.Time      `json:"bootstrappedAt"`
	CreatedAt    time.Time       `json:"createdAt"`

	RoutingTableRefresh RoutingTableRefresh `json:"routingTableRefresh"`
}

type RoutingTableRefresh struct {
	StartedAt *time.Time `json:"startedAt"`
	EndedAt   *time.Time `json:"endedAt"`
	Error     *string    `json:"error"`
}

func New(ctx context.Context) (*Host, error) {
	key, _, err := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	if err != nil {
		return nil, errors.Wrap(err, "generate key pair")
	}

	localID, err := peer.IDFromPublicKey(key.GetPublic())
	if err != nil {
		return nil, errors.Wrap(err, "id from public key")
	}

	var dht *kaddht.IpfsDHT
	h, err := libp2p.New(
		libp2p.Identity(key),
		libp2p.Transport(wrap.NewTCPTransport(localID)),
		libp2p.Transport(wrap.NewWSTransport(localID)),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			dht, err = kaddht.New(ctx, h, kaddht.Mode(kaddht.ModeClient))
			return dht, err
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "new libp2p host")
	}

	newHost := &Host{
		PeerID:    h.ID(),
		Host:      h,
		DHT:       dht,
		CreatedAt: time.Now(),
	}

	log.Infow("Initialized new libp2p host", "localID", utils.FmtPeerID(h.ID()))
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
		log.Infow("Connecting to bootstrap peer", "remoteID", utils.FmtPeerID(bp.ID))
		if err := h.Host.Connect(ctx, bp); err != nil {
			return errors.Wrap(err, "connecting to bootstrap peer")
		}
	}
	h.Bootstrapped = now()
	return nil
}

func (h *Host) RefreshRoutingTable(ctx context.Context) {
	log.Infow("Start refreshing routing table")
	defer log.Infow("Done refreshing routing table")

	h.RoutingTableRefresh.StartedAt = now()
	defer func() { h.RoutingTableRefresh.EndedAt = now() }()

	h.RoutingTableRefresh.Error = nil
	h.RoutingTableRefresh.EndedAt = nil

	select {
	case err := <-h.DHT.RefreshRoutingTable():
		if err != nil {
			h.RoutingTableRefresh.Error = strPtr(err.Error())
		}
	case <-ctx.Done():
		if ctx.Err() != nil {
			h.RoutingTableRefresh.Error = strPtr(ctx.Err().Error())
		}
	}
}

func (h *Host) BucketIdForPeer(p peer.ID) int {
	peerID := kbucket.ConvertPeerID(p)
	cpl := kbucket.CommonPrefixLen(peerID, kbucket.ConvertPeerID(h.PeerID))
	bucketID := cpl
	if bucketID >= 20 {
		bucketID = 20 - 1
	}
	return bucketID
}

func now() *time.Time {
	now := time.Now()
	return &now
}

func strPtr(str string) *string {
	return &str
}
