package main

import (
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
)

// Measurement keeps track of the starting conditions and results of an experiment.
type Measurement struct {
	StartedAt  time.Time
	EndedAt    time.Time
	ContentID  string
	Provider   RunData
	Requesters map[string]RunData
	InitRT     bool
}

type RunData struct {
	StartedAt time.Time
	EndedAt   time.Time
	LocalID   peer.ID
	Distance  string
	Spans     []SpanData
	PeerInfos map[string]PeerInfo
	PeerOrder []peer.ID
}

type SpanData struct {
	RelStart  float64
	DurationS float64
	Start     time.Time
	End       time.Time
	PeerID    peer.ID
	Operation SpanOperation
	Type      string
	Error     string
}

type PeerInfo struct {
	ID              peer.ID
	AgentVersion    string
	Distance        string
	RelDiscoveredAt float64
	DiscoveredAt    time.Time
	DiscoveredFrom  peer.ID
	IsBootstrap     bool
}
