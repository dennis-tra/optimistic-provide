package service

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

type DialSpan struct {
	RemotePeerID peer.ID
	Maddr        ma.Multiaddr
	Start        time.Time
	End          time.Time
	Trpt         string
	Error        error
}

type ConnectionSpan struct {
	RemotePeerID peer.ID
	Maddr        ma.Multiaddr
	Start        time.Time
	End          time.Time
}

type QuerySpan struct {
	RemotePeerID peer.ID
	Start        time.Time
	End          time.Time
	Error        error
}

type FindNodesSpan struct {
	RemotePeerID peer.ID
	Start        time.Time
	End          time.Time
	CloserPeers  []*peer.AddrInfo
	Error        error
}
