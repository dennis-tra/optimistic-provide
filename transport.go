package main

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	tptu "github.com/libp2p/go-libp2p-transport-upgrader"
	"github.com/libp2p/go-tcp-transport"
	websocket "github.com/libp2p/go-ws-transport"
	ma "github.com/multiformats/go-multiaddr"
)

// TCPTransport is a thin wrapper around the actual *tcp.TcpTransport implementation.
// It intercepts calls to Dial to track when which peer was dialed.
type TCPTransport struct {
	// The peer ID of the dial-initiating peer
	local peer.ID

	// Span channel to publish dial spans on
	sc chan<- Span

	// The original TCP transport implementation
	trpt *tcp.TcpTransport
}

func NewTCPTransport(local peer.ID, ec chan<- Span) func(*tptu.Upgrader) *TCPTransport {
	return func(u *tptu.Upgrader) *TCPTransport {
		return &TCPTransport{
			local: local,
			sc:    ec,
			trpt:  tcp.NewTCPTransport(u),
		}
	}
}

func (t *TCPTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	start := time.Now()
	dial, err := t.trpt.Dial(ctx, raddr, p)
	t.sc <- &DialSpan{
		BaseSpan:  NewBaseSpan(start, t.local, p, err),
		Transport: "tcp",
		Maddr:     raddr,
	}
	return dial, err
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

// WSTransport is a thin wrapper around the actual *websocket.WebsocketTransport
// implementation. It intercepts calls to Dial to track when which peer is dialed.
type WSTransport struct {
	// The peer ID of the dial-initiating peer
	local peer.ID

	// Span channel to publish dial spans on
	sc chan<- Span

	// The original websocket transport implementation
	trpt *websocket.WebsocketTransport
}

func NewWSTransport(local peer.ID, ec chan<- Span) func(u *tptu.Upgrader) *WSTransport {
	return func(u *tptu.Upgrader) *WSTransport {
		return &WSTransport{
			local: local,
			sc:    ec,
			trpt:  websocket.New(u),
		}
	}
}

func (ws *WSTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	start := time.Now()
	dial, err := ws.trpt.Dial(ctx, raddr, p)
	ws.sc <- &DialSpan{
		BaseSpan:  NewBaseSpan(start, ws.local, p, err),
		Transport: "tcp",
		Maddr:     raddr,
	}
	return dial, err
}

func (ws *WSTransport) CanDial(addr ma.Multiaddr) bool {
	return ws.trpt.CanDial(addr)
}

func (ws *WSTransport) Listen(laddr ma.Multiaddr) (transport.Listener, error) {
	return ws.trpt.Listen(laddr)
}

func (ws *WSTransport) Protocols() []int {
	return ws.trpt.Protocols()
}

func (ws *WSTransport) Proxy() bool {
	return ws.trpt.Proxy()
}
