package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/ipfs/go-cid"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"

	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

type RetrievalController struct {
	ctx context.Context
	dbc *db.Client
	rs  service.RetrievalService
	hs  service.HostService
}

func NewRetrievalController(ctx context.Context, rs service.RetrievalService, hs service.HostService) *RetrievalController {
	return &RetrievalController{
		ctx: ctx,
		rs:  rs,
		hs:  hs,
	}
}

func (rc *RetrievalController) Create(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	rr := &types.RetrievalRequest{}

	if err := c.BindJSON(rr); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could start retrieval because of a malformed JSON request",
			Details: types.ErrDetails(err),
		})
		return
	}

	contentID, err := cid.Decode(rr.ContentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could start retrieval because of a malformed content ID: " + rr.ContentId,
			Details: types.ErrDetails(err),
		})
		return
	}

	retrieval, err := rc.rs.Retrieve(rc.ctx, h, contentID, rr.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Error starting retrieval operation",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusCreated, types.Retrieval{
		RetrievalId:           retrieval.ID,
		ContentId:             retrieval.ContentID,
		EndedAt:               nil,
		Error:                 nil,
		FinalRoutingTableId:   nil,
		HostId:                h.ID().String(),
		InitialRoutingTableId: retrieval.InitialRoutingTableID,
		StartedAt:             retrieval.StartedAt.Format(time.RFC3339),
	})
}