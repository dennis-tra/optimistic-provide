package wrap

import (
	"context"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	"github.com/libp2p/go-tcp-transport"
	ma "github.com/multiformats/go-multiaddr"
)

// TCPTransport is a thin wrapper around the actual *tcp.TcpTransport implementation.
// It intercepts calls to Dial to track when which peer was dialed.
type TCPTransport struct {
	// The peer ID of the dial-initiating peer
	local peer.ID

	// The original TCP transport implementation
	trpt *tcp.TcpTransport
}

func NewTCPTransport(local peer.ID) func(transport.Upgrader, network.ResourceManager, ...tcp.Option) (*TCPTransport, error) {
	return func(upgrader transport.Upgrader, rcmgr network.ResourceManager, opts ...tcp.Option) (*TCPTransport, error) {
		trpt, err := tcp.NewTCPTransport(upgrader, rcmgr, opts...)
		if err != nil {
			return nil, err
		}
		return &TCPTransport{
			local: local,
			trpt:  trpt,
		}, nil
	}
}

func (t *TCPTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	return t.trpt.Dial(ctx, raddr, p)
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
