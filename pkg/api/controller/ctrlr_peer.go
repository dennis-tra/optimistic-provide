package controller

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/api/render"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/service"
	"github.com/gin-gonic/gin"
)

// PeerController holds the API logic for all routes in /hosts
type PeerController struct {
	ctx context.Context
	svc service.PeerService
}

// NewPeerController initializes a new host controller with the provided services.
func NewPeerController(ctx context.Context, svc service.PeerService) *PeerController {
	return &PeerController{
		ctx: ctx,
		svc: svc,
	}
}

func (pc *PeerController) Get(c *gin.Context) {
	pidStr, ok := c.Params.Get("peerID")
	if !ok {
		render.BadRequest(c, nil, "peer ID not found in query")
		return
	}

	pid, err := peer.Decode(pidStr)
	if err != nil {
		render.BadRequest(c, err, "could not decode peer ID")
		return
	}

	dbPeer, err := pc.svc.Find(c.Request.Context(), pid)
	if err != nil {
		render.InternalServerError(c, err, "could not find peer ID")
		return
	}

	render.OK(c, dbPeer)
}
