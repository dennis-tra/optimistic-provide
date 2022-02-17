package controller

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/render"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

// PeerController holds the API logic for all routes in /hosts
type PeerController struct {
	ctx context.Context
	ps  service.PeerService
}

// NewPeerController initializes a new host controller with the provided services.
func NewPeerController(ctx context.Context, ps service.PeerService) *PeerController {
	return &PeerController{
		ctx: ctx,
		ps:  ps,
	}
}

func (pc *PeerController) Get(c *gin.Context) {
	pid, rerr := getHostID(c)
	if rerr != nil {
		render.Err(c, rerr)
		return
	}

	dbPeer, err := pc.ps.Find(c.Request.Context(), pid)
	if err != nil {
		render.ErrInternal(c, "could not find peer ID", err)
		return
	}

	render.OK(c, dbPeer)
}
