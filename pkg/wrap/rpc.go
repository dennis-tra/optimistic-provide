package wrap

import (
	"context"
	"time"

	"github.com/google/uuid"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"

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

var (
	_ Disconnector     = (*MessageSenderImpl)(nil)
	_ pb.MessageSender = (*MessageSenderImpl)(nil)
)

func NewMessageSenderImpl() *MessageSenderImpl {
	return &MessageSenderImpl{}
}

func (m *MessageSenderImpl) Init(h host.Host, protos []protocol.ID) pb.MessageSender {
	m.messenger = net.NewMessageSenderImpl(h, protos)
	return m
}

func (m *MessageSenderImpl) SendRequest(ctx context.Context, p peer.ID, pmes *pb.Message) (*pb.Message, error) {
	var qid uuid.UUID
	if val := ctx.Value(kaddht.QueryIdCtxKey{}); val != nil {
		qid = val.(uuid.UUID)
	}
	started := &RPCSendRequestStartedEvent{
		QueryID:    qid,
		RemotePeer: p,
		StartedAt:  time.Now(),
		Request:    pmes,
	}
	PublishRPCEvent(ctx, started)
	resp, err := m.messenger.SendRequest(ctx, p, pmes)
	ended := &RPCSendRequestEndedEvent{
		QueryID:    qid,
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
	var qid uuid.UUID
	if val := ctx.Value(kaddht.QueryIdCtxKey{}); val != nil {
		qid = val.(uuid.UUID)
	}
	started := &RPCSendMessageStartedEvent{
		QueryID:    qid,
		RemotePeer: p,
		StartedAt:  time.Now(),
		Message:    pmes,
	}
	PublishRPCEvent(ctx, started)
	err := m.messenger.SendMessage(ctx, p, pmes)
	ended := &RPCSendMessageEndedEvent{
		QueryID:    qid,
		RemotePeer: p,
		StartedAt:  started.StartedAt,
		EndedAt:    time.Now(),
		Message:    pmes,
		Error:      err,
	}
	PublishRPCEvent(ctx, ended)
	return err
}

func (m *MessageSenderImpl) OnDisconnect(ctx context.Context, p peer.ID) {
	m.messenger.(Disconnector).OnDisconnect(ctx, p)
}
