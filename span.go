package main

import (
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
	ma "github.com/multiformats/go-multiaddr"
)

type SpanOperation string

const (
	SpanOpDial        SpanOperation = "dial"
	SpanOpSendRequest SpanOperation = "send_request"
	SpanOpSendMessage SpanOperation = "send_message"
)

// The Span interface must be implemented by all sc
// that need to be tracked.
type Span interface {
	LocalID() peer.ID
	RemoteID() peer.ID
	StartedAt() time.Time
	Duration() time.Duration
	EndedAt() time.Time
	Operation() SpanOperation
	Type() string
	Error() error
}

type BaseSpan struct {
	Start  time.Time
	End    time.Time
	Local  peer.ID
	Remote peer.ID
	Err    error
}

func NewBaseSpan(start time.Time, local peer.ID, remote peer.ID, err error) BaseSpan {
	return BaseSpan{
		Local:  local,
		Remote: remote,
		Start:  start,
		End:    time.Now(),
		Err:    err,
	}
}

func (bs *BaseSpan) LocalID() peer.ID {
	return bs.Local
}

func (bs *BaseSpan) RemoteID() peer.ID {
	return bs.Remote
}

func (bs *BaseSpan) StartedAt() time.Time {
	return bs.Start
}

func (bs *BaseSpan) Duration() time.Duration {
	return bs.End.Sub(bs.Start)
}

func (bs *BaseSpan) EndedAt() time.Time {
	return bs.End
}

func (bs *BaseSpan) Error() error {
	return bs.Err
}

type DialSpan struct {
	BaseSpan
	Transport string
	Maddr     ma.Multiaddr
}

func (ds *DialSpan) Operation() SpanOperation {
	return SpanOpDial
}

func (ds *DialSpan) Type() string {
	return ds.Transport
}

type SendRequestSpan struct {
	BaseSpan
	Request      *pb.Message
	Response     *pb.Message
	AgentVersion string
	Protocols    []string
}

func NewSendRequestSpan(h host.Host, start time.Time, local peer.ID, remote peer.ID, req *pb.Message, resp *pb.Message, err error) *SendRequestSpan {
	span := &SendRequestSpan{
		BaseSpan:  NewBaseSpan(start, local, remote, err),
		Request:   req,
		Response:  resp,
		Protocols: []string{},
	}

	// Extract agent
	if agent, err := h.Peerstore().Get(remote, "AgentVersion"); err == nil {
		span.AgentVersion = agent.(string)
	}

	// Extract protocols
	if protocols, err := h.Peerstore().GetProtocols(remote); err == nil {
		span.Protocols = protocols
	}

	return span
}

func (srs *SendRequestSpan) ReturnedCloserPeer(peerID peer.ID) bool {
	if srs.Response == nil {
		return false
	}

	for _, closer := range pb.PBPeersToPeerInfos(srs.Response.CloserPeers) {
		if closer.ID == peerID {
			return true
		}
	}

	return false
}

func (srs *SendRequestSpan) ReturnedProvider(peerID peer.ID) bool {
	if srs.Response == nil {
		return false
	}

	for _, provider := range pb.PBPeersToPeerInfos(srs.Response.ProviderPeers) {
		if provider.ID == peerID {
			return true
		}
	}

	return false
}

func (srs *SendRequestSpan) Operation() SpanOperation {
	return SpanOpSendRequest
}

func (srs *SendRequestSpan) Type() string {
	return srs.Request.Type.String()
}

type SendMessageSpan struct {
	BaseSpan
	Message      *pb.Message
	AgentVersion string
	Protocols    []string
}

func NewSendMessageSpan(h host.Host, start time.Time, local peer.ID, remote peer.ID, pmes *pb.Message, err error) *SendMessageSpan {
	span := &SendMessageSpan{
		BaseSpan:  NewBaseSpan(start, local, remote, err),
		Message:   pmes,
		Protocols: []string{},
	}

	// Extract agent
	if agent, err := h.Peerstore().Get(remote, "AgentVersion"); err == nil {
		span.AgentVersion = agent.(string)
	}

	// Extract protocols
	if protocols, err := h.Peerstore().GetProtocols(remote); err == nil {
		span.Protocols = protocols
	}

	return span
}

func (sms *SendMessageSpan) Operation() SpanOperation {
	return SpanOpSendMessage
}

func (sms *SendMessageSpan) Type() string {
	return sms.Message.Type.String()
}
