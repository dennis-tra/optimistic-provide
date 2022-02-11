package wrap

import (
	"context"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	websocket "github.com/libp2p/go-ws-transport"
	ma "github.com/multiformats/go-multiaddr"
)

// WSTransport is a thin wrapper around the actual *websocket.WebsocketTransport
// implementation. It intercepts calls to Dial to track when which peer is dialed.
type WSTransport struct {
	// The peer ID of the dial-initiating peer
	local peer.ID

	// The original websocket transport implementation
	trpt *websocket.WebsocketTransport
}

func NewWSTransport(local peer.ID) func(upgrader transport.Upgrader, rcmgr network.ResourceManager) *WSTransport {
	return func(upgrader transport.Upgrader, rcmgr network.ResourceManager) *WSTransport {
		return &WSTransport{
			local: local,
			trpt:  websocket.New(upgrader, rcmgr),
		}
	}
}

func (ws *WSTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	return ws.trpt.Dial(ctx, raddr, p)
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
