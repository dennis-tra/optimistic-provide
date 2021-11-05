package main

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"

	u "github.com/ipfs/go-ipfs-util"
	kbucket "github.com/libp2p/go-libp2p-kbucket"

	"github.com/stretchr/testify/assert"

	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/require"
)

func TestRun_IsBootstrapPeer(t *testing.T) {
	ctx := context.Background()
	mnet := mocknet.New(ctx)

	local, err := mnet.GenPeer()
	require.NoError(t, err)

	remote, err := mnet.GenPeer()
	require.NoError(t, err)

	other, err := mnet.GenPeer()
	require.NoError(t, err)

	start := time.Now()
	tests := []struct {
		name  string
		spans []Span
		want  bool
	}{
		{
			name: "single dial",
			spans: []Span{
				&DialSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote.ID(), nil),
				},
			},
			want: false,
		},
		{
			name: "single send request",
			spans: []Span{
				&SendRequestSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote.ID(), nil),
				},
			},
			want: true,
		},
		{
			name: "dial after send request",
			spans: []Span{
				&SendRequestSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote.ID(), nil),
				},
				&DialSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote.ID(), nil),
				},
			},
			want: true,
		},
		{
			name: "dial before send request",
			spans: []Span{
				&DialSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote.ID(), nil),
				},
				&SendRequestSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote.ID(), nil),
				},
			},
			want: false,
		},
		{
			name: "dial other peer and after send request",
			spans: []Span{
				&DialSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), other.ID(), nil),
				},
				&SendRequestSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote.ID(), nil),
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Run{
				LocalID:   local.ID(),
				Spans:     tt.spans,
				StartedAt: start,
			}
			assert.Equal(t, r.IsBootstrapPeer(remote.ID()), tt.want)
		})
	}
}

func TestRun_GatherPeerInfos(t *testing.T) {
	ctx := context.Background()
	agent := "test-agent"

	mnet := mocknet.New(ctx)

	local, err := mnet.GenPeer()
	require.NoError(t, err)

	remote1, err := mnet.GenPeer()
	require.NoError(t, err)

	remote2, err := mnet.GenPeer()
	require.NoError(t, err)

	remote3, err := mnet.GenPeer()
	require.NoError(t, err)

	content, err := NewRandomContent()
	require.NoError(t, err)

	start := time.Now()
	distance := hex.EncodeToString(u.XOR(kbucket.ConvertPeerID(remote1.ID()), kbucket.ConvertKey(string(content.cid.Hash()))))

	tests := []struct {
		name     string
		involved map[peer.ID]struct{}
		spans    []Span
		want     PeerInfo
	}{
		{
			name: "standard",
			involved: map[peer.ID]struct{}{
				remote1.ID(): {},
			},
			spans: []Span{
				&DialSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote1.ID(), nil),
				},
				&SendRequestSpan{
					BaseSpan:     NewBaseSpan(start.Add(time.Second), local.ID(), remote1.ID(), nil),
					AgentVersion: agent,
				},
			},
			want: PeerInfo{
				ID:              remote1.ID(),
				AgentVersion:    agent,
				Distance:        distance,
				RelDiscoveredAt: 0,
				DiscoveredAt:    time.Time{},
				DiscoveredFrom:  "",
				IsBootstrap:     false,
				FirstSpanAt:     start,
			},
		},
		{
			name: "standard",
			involved: map[peer.ID]struct{}{
				remote1.ID(): {},
				remote2.ID(): {},
			},
			spans: []Span{
				&DialSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote2.ID(), nil),
				},
				&SendRequestSpan{
					BaseSpan:     NewBaseSpan(start.Add(time.Second), local.ID(), remote2.ID(), nil),
					AgentVersion: agent,
					Response: &pb.Message{
						Type:        pb.Message_FIND_NODE,
						CloserPeers: pb.RawPeerInfosToPBPeers([]peer.AddrInfo{{ID: remote1.ID()}}),
					},
				},
				&DialSpan{
					BaseSpan: NewBaseSpan(start.Add(2*time.Second), local.ID(), remote1.ID(), nil),
				},
			},
			want: PeerInfo{
				ID:              remote1.ID(),
				AgentVersion:    "",
				Distance:        distance,
				RelDiscoveredAt: time.Second.Seconds(),
				DiscoveredAt:    start.Add(time.Second),
				DiscoveredFrom:  remote2.ID(),
				IsBootstrap:     false,
				FirstSpanAt:     start.Add(2 * time.Second),
			},
		},
		{
			name: "discovered later",
			involved: map[peer.ID]struct{}{
				remote1.ID(): {},
				remote2.ID(): {},
			},
			spans: []Span{
				&SendRequestSpan{
					BaseSpan:     NewBaseSpan(start, local.ID(), remote1.ID(), nil),
					AgentVersion: agent,
					Response: &pb.Message{
						Type:        pb.Message_FIND_NODE,
						CloserPeers: pb.RawPeerInfosToPBPeers([]peer.AddrInfo{}),
					},
				},
				&SendRequestSpan{
					BaseSpan:     NewBaseSpan(start.Add(time.Second), local.ID(), remote2.ID(), nil),
					AgentVersion: agent,
					Response: &pb.Message{
						Type:        pb.Message_FIND_NODE,
						CloserPeers: pb.RawPeerInfosToPBPeers([]peer.AddrInfo{{ID: remote1.ID()}}),
					},
				},
			},
			want: PeerInfo{
				ID:              remote1.ID(),
				AgentVersion:    agent,
				Distance:        distance,
				RelDiscoveredAt: 0,
				DiscoveredAt:    time.Time{},
				DiscoveredFrom:  "",
				IsBootstrap:     true,
				FirstSpanAt:     start,
			},
		},
		{
			name: "discovered twice",
			involved: map[peer.ID]struct{}{
				remote1.ID(): {},
				remote2.ID(): {},
				remote3.ID(): {},
			},
			spans: []Span{
				&SendRequestSpan{
					BaseSpan: BaseSpan{
						Local:  local.ID(),
						Remote: remote2.ID(),
						Start:  start,
						End:    start.Add(10 * time.Second),
					},
					AgentVersion: agent,
					Response: &pb.Message{
						Type:        pb.Message_FIND_NODE,
						CloserPeers: pb.RawPeerInfosToPBPeers([]peer.AddrInfo{{ID: remote1.ID()}}),
					},
				},
				&SendRequestSpan{
					BaseSpan: BaseSpan{
						Local:  local.ID(),
						Remote: remote3.ID(),
						Start:  start.Add(time.Second),
						End:    start.Add(5 * time.Second),
					},
					AgentVersion: agent,
					Response: &pb.Message{
						Type:        pb.Message_FIND_NODE,
						CloserPeers: pb.RawPeerInfosToPBPeers([]peer.AddrInfo{{ID: remote1.ID()}}),
					},
				},
				&DialSpan{
					BaseSpan: NewBaseSpan(start.Add(20*time.Second), local.ID(), remote1.ID(), nil),
				},
			},
			want: PeerInfo{
				ID:              remote1.ID(),
				AgentVersion:    "",
				Distance:        distance,
				RelDiscoveredAt: 0,
				DiscoveredAt:    start.Add(5 * time.Second),
				DiscoveredFrom:  remote3.ID(),
				IsBootstrap:     false,
				FirstSpanAt:     start.Add(20 * time.Second),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Run{
				StartedAt: start,
				EndedAt:   start.Add(10 * time.Second),
				LocalID:   local.ID(),
				Spans:     tt.spans,
				Involved:  tt.involved,
			}
			got := r.GatherPeerInfos(content)

			info, found := got[remote1.ID().Pretty()]
			require.True(t, found)

			info.RelDiscoveredAt = tt.want.RelDiscoveredAt
			info.DiscoveredAt = tt.want.DiscoveredAt
			assert.Equal(t, tt.want, info)
		})
	}
}

func TestRun_PeerOrder(t *testing.T) {
	ctx := context.Background()
	agent := "test-agent"

	mnet := mocknet.New(ctx)

	local, err := mnet.GenPeer()
	require.NoError(t, err)

	remote1, err := mnet.GenPeer()
	require.NoError(t, err)

	remote2, err := mnet.GenPeer()
	require.NoError(t, err)

	content, err := NewRandomContent()
	require.NoError(t, err)

	start := time.Now()

	tests := []struct {
		name     string
		involved map[peer.ID]struct{}
		spans    []Span
		want     []peer.ID
	}{
		{
			name:     "bootstrap after discovered",
			involved: map[peer.ID]struct{}{remote1.ID(): {}, remote2.ID(): {}},
			spans: []Span{
				&DialSpan{
					BaseSpan: NewBaseSpan(start, local.ID(), remote2.ID(), nil),
				},
				&SendRequestSpan{
					BaseSpan:     NewBaseSpan(start.Add(time.Second), local.ID(), remote1.ID(), nil),
					AgentVersion: agent,
				},
			},
			want: []peer.ID{remote1.ID(), remote2.ID()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Run{
				StartedAt: start,
				EndedAt:   start.Add(10 * time.Second),
				LocalID:   local.ID(),
				Spans:     tt.spans,
				Involved:  tt.involved,
			}
			got := r.PeerOrder(r.GatherPeerInfos(content))
			assert.Equal(t, tt.want, got)
		})
	}
}
