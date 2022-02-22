package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

type ProvideController struct {
	ctx context.Context
	dbc *db.Client
	ps  service.ProvideService
	hs  service.HostService
}

func NewProvideController(ctx context.Context, ps service.ProvideService, hs service.HostService) *ProvideController {
	return &ProvideController{
		ctx: ctx,
		ps:  ps,
		hs:  hs,
	}
}

func (pc *ProvideController) Create(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	provide, err := pc.ps.Provide(pc.ctx, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Error starting provide operation",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusCreated, types.Provide{
		ContentId:             provide.ContentID,
		EndedAt:               nil,
		Error:                 nil,
		FinalRoutingTableId:   nil,
		HostId:                h.ID().String(),
		InitialRoutingTableId: provide.InitialRoutingTableID,
		ProvideId:             provide.ID,
		StartedAt:             provide.StartedAt.Format(time.RFC3339),
	})
}
