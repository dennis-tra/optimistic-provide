package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

// ConnectionController holds the API logic for all routes under /hosts
type ConnectionController struct {
	ctx context.Context
	cs  service.ConnectionService
	ps  service.ProvideService
}

// NewConnectionController initializes a new host controller with the provided services.
func NewConnectionController(ctx context.Context, ds service.ConnectionService, ps service.ProvideService) *ConnectionController {
	return &ConnectionController{
		ctx: ctx,
		cs:  ds,
		ps:  ps,
	}
}

// List lists returns all running libp2p hosts.
func (cc *ConnectionController) List(c *gin.Context) {
	provideID := c.MustGet("provideID").(int)

	provide, err := cc.ps.GetByID(c.Request.Context(), provideID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not get provide record for ID " + strconv.Itoa(provideID),
			Details: types.ErrDetails(err),
		})
		return
	}

	dbConnections, err := cc.cs.List(c.Request.Context(), provide)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not get dials from DB",
			Details: types.ErrDetails(err),
		})
		return
	}

	dials := make([]*types.Connection, len(dbConnections))
	for i, dbConnection := range dbConnections {
		dials[i] = &types.Connection{
			Id:           dbConnection.ID,
			MultiAddress: dbConnection.R.MultiAddress.Maddr,
			RemoteId:     dbConnection.R.Remote.MultiHash,
			DurationInS:  float32(dbConnection.EndedAt.Sub(dbConnection.StartedAt).Seconds()),
			EndedAt:      dbConnection.EndedAt.Format(time.RFC3339Nano),
			StartedAt:    dbConnection.StartedAt.Format(time.RFC3339Nano),
		}
	}

	c.JSON(http.StatusOK, dials)
}
