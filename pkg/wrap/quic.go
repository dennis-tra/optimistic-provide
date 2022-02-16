package wrap

import (
	"context"
	"time"

	"github.com/dennis-tra/optimistic-provide/pkg/db/models"

	"github.com/libp2p/go-libp2p-core/connmgr"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/pnet"
	quic "github.com/libp2p/go-libp2p-quic-transport"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	ma "github.com/multiformats/go-multiaddr"
)

// QuicTransport is a thin wrapper around the actual *websocket.WebsocketTransport
// implementation. It intercepts calls to Dial to track when which peer is dialed.
type QuicTransport struct {
	*Notifier

	// The original websocket transport implementation
	trpt transport.Transport
}

func NewQuicTransport() (*QuicTransport, func(key ic.PrivKey, psk pnet.PSK, gater connmgr.ConnectionGater, rcmgr network.ResourceManager) (*QuicTransport, error)) {
	q := &QuicTransport{
		Notifier: newNotifier(models.DialTransportQuic),
	}
	return q, func(key ic.PrivKey, psk pnet.PSK, gater connmgr.ConnectionGater, rcmgr network.ResourceManager) (*QuicTransport, error) {
		trpt, err := quic.NewTransport(key, psk, gater, rcmgr)
		q.trpt = trpt
		return q, err
	}
}

func (q *QuicTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	start := time.Now()
	q.notifyDialStarted(raddr, p, start)
	c, err := q.trpt.Dial(ctx, raddr, p)
	q.notifyDialEnded(raddr, p, start, time.Now(), err)
	return c, err
}

func (q *QuicTransport) CanDial(addr ma.Multiaddr) bool {
	return q.trpt.CanDial(addr)
}

func (q *QuicTransport) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	return q.trpt.Listen(laddr)
}

func (q *QuicTransport) Protocols() []int {
	return q.trpt.Protocols()
}

func (q *QuicTransport) Proxy() bool {
	return q.trpt.Proxy()
}
