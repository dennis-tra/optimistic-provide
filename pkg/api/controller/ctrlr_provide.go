package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

type ProvideController struct {
	ctx context.Context
	dbc *db.Client

	ps service.ProvideService
}

func NewProvideController(ctx context.Context, ps service.ProvideService) *ProvideController {
	return &ProvideController{
		ctx: ctx,
		ps:  ps,
	}
}

type CreateProvideRequest struct {
	HostID string `json:"hostId"`
}

type CreateProvideResponse struct {
	ProviderID            string    `json:"providerId"`
	ProvideID             int       `json:"provideId"`
	ContentID             string    `json:"contentId"`
	InitialRoutingTableID int       `json:"initialRoutingTableId"`
	StartedAt             time.Time `json:"startedAt"`
}

func (pc *ProvideController) Create(c *gin.Context) {
	// TODO: check pending provide
	req := CreateProvideRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	hostID, err := peer.Decode(req.HostID)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	provide, err := pc.ps.Provide(pc.ctx, hostID)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	resp := CreateProvideResponse{
		ProviderID:            hostID.String(),
		ProvideID:             provide.ID,
		ContentID:             provide.ContentID,
		InitialRoutingTableID: provide.InitialRoutingTableID,
		StartedAt:             time.Now(),
	}

	c.JSON(http.StatusOK, resp)
}
