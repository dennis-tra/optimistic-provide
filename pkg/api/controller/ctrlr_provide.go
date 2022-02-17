package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/api/entities"
	"github.com/dennis-tra/optimistic-provide/pkg/api/render"
	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

type ProvideController struct {
	ctx context.Context
	dbc *db.Client
	ps  service.ProvideService
}

func NewProvideController(ctx context.Context, ps service.ProvideService) *ProvideController {
	return &ProvideController{
		ctx: ctx,
		ps:  ps,
	}
}

func (pc *ProvideController) Create(c *gin.Context) {
	// TODO: check pending provide
	req := entities.ProvideCreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		render.NewErrBadRequest(render.ErrorCodeMalformedJSON, "Could not parse JSON body", err)
		return
	}

	hostID, err := peer.Decode(req.HostID)
	if err != nil {
		render.Err(c, render.NewErrBadRequest(render.ErrorCodeMalformedPeerID, "Could not decode peer ID: "+req.HostID, err))
		return
	}

	provide, err := pc.ps.Provide(pc.ctx, hostID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	resp := entities.ProvideResponse{
		ProvideID:             provide.ID,
		HostID:                hostID.String(),
		ContentID:             provide.ContentID,
		InitialRoutingTableID: provide.InitialRoutingTableID,
		StartedAt:             time.Now(),
	}

	c.JSON(http.StatusOK, resp)
}
