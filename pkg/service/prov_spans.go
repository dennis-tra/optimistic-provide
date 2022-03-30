package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"

	"github.com/dennis-tra/optimistic-provide/pkg/util"
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

type FindNodesSpan struct {
	QueryID      uuid.UUID
	RemotePeerID peer.ID
	Start        time.Time
	End          time.Time
	CloserPeers  []*peer.AddrInfo
	Error        error
}

type GetProvidersSpan struct {
	QueryID      *uuid.UUID
	RemotePeerID peer.ID
	Start        time.Time
	End          time.Time
	Providers    []*peer.AddrInfo
	CloserPeers  []*peer.AddrInfo
	Error        error
}

type AddProvidersSpan struct {
	QueryID       uuid.UUID
	RemotePeerID  peer.ID
	Content       *util.Content
	Start         time.Time
	ProviderAddrs []ma.Multiaddr
	End           time.Time
	Error         error
}
