package controller

import (
	"context"
	"net/http"
	"sort"
	"time"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
	"github.com/gin-gonic/gin"
)

// HostController holds the API logic for all routes under /hosts
type HostController struct {
	ctx context.Context
	hs  service.HostService
}

// NewHostController initializes a new host controller with the provided services.
func NewHostController(ctx context.Context, hs service.HostService) *HostController {
	return &HostController{
		ctx: ctx,
		hs:  hs,
	}
}

// Create starts a new libp2p host.
func (hc *HostController) Create(c *gin.Context) {
	chr := &types.CreateHostRequest{}
	if err := c.BindJSON(chr); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could not create libp2p host because of a malformed JSON request",
			Details: types.ErrDetails(err),
		})
		return
	}

	h, err := hc.hs.Create(hc.ctx, chr.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not create libp2p host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusCreated, &types.Host{
		Name:           h.DBHost.Name,
		HostId:         h.ID().String(),
		BootstrappedAt: nil,
		StartedAt:      util.StrPtr(h.StartedAt.Format(time.RFC3339Nano)),
		CreatedAt:      h.DBHost.CreatedAt.Format(time.RFC3339Nano),
	})
}

// List lists returns all running libp2p hosts.
func (hc *HostController) List(c *gin.Context) {
	hosts := []*types.Host{}

	runningHosts, dbHosts, err := hc.hs.Hosts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not get running and DB hosts",
			Details: types.ErrDetails(err),
		})
		return
	}

	for _, dbHost := range dbHosts {
		respHost := &types.Host{
			Name:      dbHost.Name,
			HostId:    dbHost.R.Peer.MultiHash,
			CreatedAt: dbHost.CreatedAt.Format(time.RFC3339Nano),
		}

		for _, runningHost := range runningHosts {
			if runningHost.ID().String() == dbHost.R.Peer.MultiHash {
				respHost.StartedAt = util.StrPtr(runningHost.StartedAt.Format(time.RFC3339Nano))
				respHost.BootstrappedAt = util.TimeToStr(runningHost.Bootstrapped)
				break
			}
		}

		hosts = append(hosts, respHost)
	}

	sort.Slice(hosts, func(i, j int) bool {
		ti, _ := time.Parse(time.RFC3339Nano, hosts[i].CreatedAt)
		tj, _ := time.Parse(time.RFC3339Nano, hosts[j].CreatedAt)
		return ti.Before(tj)
	})

	c.JSON(http.StatusOK, hosts)
}

func (hc *HostController) Get(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)
	c.JSON(http.StatusOK, &types.Host{
		Name:           h.DBHost.Name,
		HostId:         h.DBHost.R.Peer.MultiHash,
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		StartedAt:      util.TimeToStr(h.StartedAt),
		CreatedAt:      h.DBHost.CreatedAt.Format(time.RFC3339Nano),
	})
}

func (hc *HostController) Start(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	h, err := hc.hs.Start(hc.ctx, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not stop libp2p host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusOK, &types.Host{
		Name:           h.DBHost.Name,
		HostId:         h.ID().String(),
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		StartedAt:      util.StrPtr(h.StartedAt.Format(time.RFC3339Nano)),
		CreatedAt:      h.DBHost.CreatedAt.Format(time.RFC3339Nano),
	})
}

func (hc *HostController) Stop(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if h.Host == nil {
		c.JSON(http.StatusOK, &types.Host{
			Name:           h.DBHost.Name,
			HostId:         h.DBHost.R.Peer.MultiHash,
			BootstrappedAt: nil,
			StartedAt:      nil,
			CreatedAt:      h.DBHost.CreatedAt.Format(time.RFC3339Nano),
		})
		return
	}

	if err := hc.hs.Stop(h.ID()); err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not stop libp2p host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusOK, &types.Host{
		Name:           h.DBHost.Name,
		HostId:         h.ID().String(),
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		StartedAt:      nil,
		CreatedAt:      h.DBHost.CreatedAt.Format(time.RFC3339Nano),
	})
}

func (hc *HostController) Bootstrap(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if h.Host == nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTSTOPPED,
			Message: "Host is stopped. Start it first to bootstrap",
		})
		return
	}

	if err := h.Bootstrap(hc.ctx); err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not bootstrap host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusOK, &types.Host{
		Name:           h.DBHost.Name,
		HostId:         h.ID().String(),
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		StartedAt:      util.StrPtr(h.StartedAt.Format(time.RFC3339Nano)),
		CreatedAt:      h.DBHost.CreatedAt.Format(time.RFC3339Nano),
	})
}

func (hc *HostController) Archive(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if err := hc.hs.Archive(c.Request.Context(), h.DBHost); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not archive host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
