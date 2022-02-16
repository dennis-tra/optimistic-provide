package wrap

import (
	"context"
	"time"

	"github.com/dennis-tra/optimistic-provide/pkg/db/models"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/transport"
	websocket "github.com/libp2p/go-ws-transport"
	ma "github.com/multiformats/go-multiaddr"
)

// WSTransport is a thin wrapper around the actual *websocket.WebsocketTransport
// implementation. It intercepts calls to Dial to track when which peer is dialed.
type WSTransport struct {
	*Notifier

	// The original websocket transport implementation
	trpt *websocket.WebsocketTransport
}

func NewWSTransport() (*WSTransport, func(transport.Upgrader, network.ResourceManager) *WSTransport) {
	ws := &WSTransport{
		Notifier: newNotifier(models.DialTransportWS),
	}
	return ws, func(upgrader transport.Upgrader, rcmgr network.ResourceManager) *WSTransport {
		ws.trpt = websocket.New(upgrader, rcmgr)
		return ws
	}
}

func (ws *WSTransport) Dial(ctx context.Context, raddr ma.Multiaddr, p peer.ID) (transport.CapableConn, error) {
	start := time.Now()
	ws.notifyDialStarted(raddr, p, start)
	c, err := ws.trpt.Dial(ctx, raddr, p)
	ws.notifyDialEnded(raddr, p, start, time.Now(), err)
	return c, err
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
