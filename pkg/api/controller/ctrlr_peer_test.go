package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
	"github.com/libp2p/go-libp2p-core/host"
)

type TestPeerService struct {
	FindFunc            func(ctx context.Context, p peer.ID) (*models.Peer, error)
	UpsertLocalPeerFunc func(h host.Host) (*models.Peer, error)
	UpsertPeerFunc      func(h host.Host, pid peer.ID) (*models.Peer, error)
}

var _ service.PeerService = (*TestPeerService)(nil)

func TestName(t *testing.T) {
	ctx := context.Background()
	svc := &TestPeerService{}
	r := gin.Default()
	ctrlr := NewPeerController(ctx, svc)

	peers := r.Group("/peers")
	{
		peers.GET("/:id", ctrlr.Get)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/peers/some-peer", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func (tps *TestPeerService) Find(ctx context.Context, p peer.ID) (*models.Peer, error) {
	return tps.FindFunc(ctx, p)
}

func (tps *TestPeerService) UpsertLocalPeer(h host.Host) (*models.Peer, error) {
	return tps.UpsertLocalPeerFunc(h)
}

func (tps *TestPeerService) UpsertPeer(h host.Host, pid peer.ID) (*models.Peer, error) {
	return tps.UpsertPeerFunc(h, pid)
}
