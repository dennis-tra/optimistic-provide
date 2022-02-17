package wrap

import (
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-kad-dht/net"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"
)

type MessageSenderImpl struct {
	messenger pb.MessageSender
}

type Disconnector interface {
	OnDisconnect(ctx context.Context, p peer.ID)
}

var _ pb.MessageSender = (*MessageSenderImpl)(nil)

func NewMessageSenderImpl() *MessageSenderImpl {
	return &MessageSenderImpl{}
}

func (m *MessageSenderImpl) Init(h host.Host, protos []protocol.ID) pb.MessageSender {
	m.messenger = net.NewMessageSenderImpl(h, protos)
	return m
}

func (m *MessageSenderImpl) SendRequest(ctx context.Context, p peer.ID, pmes *pb.Message) (*pb.Message, error) {
	started := &RPCSendRequestStartedEvent{
		RemotePeer: p,
		StartedAt:  time.Now(),
		Request:    pmes,
	}
	PublishRPCEvent(ctx, started)
	resp, err := m.messenger.SendRequest(ctx, p, pmes)
	ended := &RPCSendRequestEndedEvent{
		RemotePeer: p,
		StartedAt:  started.StartedAt,
		EndedAt:    time.Now(),
		Request:    pmes,
		Response:   resp,
		Error:      err,
	}
	PublishRPCEvent(ctx, ended)
	return resp, err
}

func (m *MessageSenderImpl) SendMessage(ctx context.Context, p peer.ID, pmes *pb.Message) error {
	return m.messenger.SendMessage(ctx, p, pmes)
}

func (m *MessageSenderImpl) OnDisconnect(ctx context.Context, p peer.ID) {
	m.messenger.(Disconnector).OnDisconnect(ctx, p)
}
