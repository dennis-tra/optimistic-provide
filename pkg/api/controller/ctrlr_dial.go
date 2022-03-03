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

// DialController holds the API logic for all routes under /hosts
type DialController struct {
	ctx context.Context
	ds  service.DialService
	ps  service.ProvideService
}

// NewDialController initializes a new host controller with the provided services.
func NewDialController(ctx context.Context, ds service.DialService, ps service.ProvideService) *DialController {
	return &DialController{
		ctx: ctx,
		ds:  ds,
		ps:  ps,
	}
}

// List lists returns all running libp2p hosts.
func (hc *DialController) List(c *gin.Context) {
	provideID := c.MustGet("provideID").(int)

	provide, err := hc.ps.GetByID(c.Request.Context(), provideID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not get provide record for ID " + strconv.Itoa(provideID),
			Details: types.ErrDetails(err),
		})
		return
	}

	dbDials, err := hc.ds.List(c.Request.Context(), provide)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not get dials from DB",
			Details: types.ErrDetails(err),
		})
		return
	}

	dials := make([]*types.Dial, len(dbDials))
	for i, dbDial := range dbDials {
		dials[i] = &types.Dial{
			Id:           dbDial.ID,
			MultiAddress: dbDial.R.MultiAddress.Maddr,
			RemoteId:     dbDial.R.Remote.MultiHash,
			DurationInS:  float32(dbDial.EndedAt.Sub(dbDial.StartedAt).Seconds()),
			EndedAt:      dbDial.EndedAt.Format(time.RFC3339Nano),
			Error:        dbDial.Error.Ptr(),
			StartedAt:    dbDial.StartedAt.Format(time.RFC3339Nano),
			Transport:    dbDial.Transport,
		}
	}

	c.JSON(http.StatusOK, dials)
}
