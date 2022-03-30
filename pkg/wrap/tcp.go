package wrap

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	"github.com/libp2p/go-tcp-transport"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

// TCPTransport is a thin wrapper around the actual *tcp.TcpTransport implementation.
// It intercepts calls to Dial to track when which peer was dialed.
type TCPTransport struct {
	*Notifier

	// The original TCP transport implementation
	trpt *tcp.TcpTransport
}

func NewTCPTransport() (*TCPTransport, func(transport.Upgrader, network.ResourceManager, ...tcp.Option) (*TCPTransport, error)) {
	t := &TCPTransport{
		Notifier: newNotifier(models.DialTransportTCP),
	}
	return t, func(upgrader transport.Upgrader, rcmgr network.ResourceManager, opts ...tcp.Option) (*TCPTransport, error) {
		trpt, err := tcp.NewTCPTransport(upgrader, rcmgr, opts...)
		if err != nil {
			return nil, err
		}
		t.trpt = trpt
		return t, nil
	}
}

func (t *TCPTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	start := time.Now()
	t.notifyDialStarted(raddr, p, start)
	c, err := t.trpt.Dial(ctx, raddr, p)
	t.notifyDialEnded(raddr, p, start, time.Now(), err)
	return c, err
}

func (t *TCPTransport) CanDial(addr ma.Multiaddr) bool {
	return t.trpt.CanDial(addr)
}

func (t *TCPTransport) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	return t.trpt.Listen(laddr)
}

func (t *TCPTransport) Protocols() []int {
	return t.trpt.Protocols()
}

func (t *TCPTransport) Proxy() bool {
	return t.trpt.Proxy()
}
