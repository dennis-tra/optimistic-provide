package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

// MeasurementController holds the API logic for all routes under /hosts
type MeasurementController struct {
	ctx context.Context
	hs  service.HostService
	ms  service.MeasurementService
}

// NewProvideMeasurementController initializes a new host controller with the provided services.
func NewProvideMeasurementController(ctx context.Context, hs service.HostService, ms service.MeasurementService) *MeasurementController {
	return &MeasurementController{
		ctx: ctx,
		hs:  hs,
		ms:  ms,
	}
}

// Create starts a new provide measurement
func (rpc *MeasurementController) Create(c *gin.Context) {
	cmr := &types.CreateMeasurementRequest{}
	if err := c.BindJSON(cmr); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could not start provide measurement due to malformed request.",
			Details: types.ErrDetails(err),
		})
		return
	}

	config := types.ProvideMeasurementConfiguration{}
	configDat, err := json.Marshal(cmr.Configuration)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could not marshal configuration to JSON",
			Details: types.ErrDetails(err),
		})
		return
	}
	if err = json.Unmarshal(configDat, &config); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could not unmarshal configuration from JSON",
			Details: types.ErrDetails(err),
		})
		return
	}

	hostID, err := peer.Decode(cmr.HostId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDPEERID,
			Message: "Host ID " + cmr.HostId + " could not be decoded.",
			Details: types.ErrDetails(err),
		})
		return
	}

	h, err := rpc.hs.Host(c.Request.Context(), hostID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTNOTFOUND,
			Message: "Host with ID " + hostID.String() + " was not found.",
			Details: types.ErrDetails(err),
		})
		return
	}

	if h.StartedAt == nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTSTOPPED,
			Message: "Host is stopped. Start it first to save a snapshot",
		})
		return
	}

	if h.Bootstrapped == nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTSTOPPED,
			Message: "Host is not bootstrapped. Bootstrap it first to start the measurement",
		})
		return
	}

	dbMeasurement, err := rpc.ms.StartProvide(c.Request.Context(), h, config)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTNOTFOUND,
			Message: "Could not start provide measurement",
			Details: types.ErrDetails(err),
		})
		return
	}

	resp := &types.CreateMeasurementResponse{
		MeasurementId: dbMeasurement.ID,
		StartedAt:     dbMeasurement.StartedAt.Format(time.RFC3339Nano),
	}

	c.JSON(http.StatusOK, resp)
}

// Stop stops the given provide measurement
func (rpc *MeasurementController) Stop(c *gin.Context) {
	measurementID := c.MustGet("measurementID").(int)

	err := rpc.ms.Stop(measurementID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTNOTFOUND,
			Message: "Measurement with ID " + strconv.Itoa(measurementID) + " is not running.",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
