package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
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

func (m *Measurement) Save(filename string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return errors.Wrap(err, "marshal measurement")
	}

	f, err := os.Create(filename)
	if err != nil {
		return errors.Wrap(err, "creating measurement")
	}
	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return errors.Wrap(err, "writing measurement")
	}

	return nil
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

	// Won't be rendered to JSON
	firstSpan Span
}
