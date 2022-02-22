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
		c.JSON(http.StatusBadRequest, types.Error{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could not create libp2p host because of a malformed JSON request",
			Details: types.ErrDetails(err),
		})
		return
	}

	h, err := hc.hs.Create(hc.ctx, chr.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not create libp2p host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusCreated, &types.Host{
		Name:           h.Name,
		HostId:         h.ID().String(),
		BootstrappedAt: nil,
		CreatedAt:      h.CreatedAt.Format(time.RFC3339),
	})
}

// List lists returns all running libp2p hosts.
func (hc *HostController) List(c *gin.Context) {
	hosts := []*types.Host{}
	for _, h := range hc.hs.Hosts() {
		hosts = append(hosts, &types.Host{
			Name:           h.Name,
			HostId:         h.ID().String(),
			BootstrappedAt: util.TimeToStr(h.Bootstrapped),
			CreatedAt:      h.CreatedAt.Format(time.RFC3339),
		})
	}

	sort.Slice(hosts, func(i, j int) bool {
		ti, _ := time.Parse(time.RFC3339, hosts[i].CreatedAt)
		tj, _ := time.Parse(time.RFC3339, hosts[j].CreatedAt)
		return ti.Before(tj)
	})

	c.JSON(http.StatusOK, hosts)
}

func (hc *HostController) Get(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)
	c.JSON(http.StatusOK, &types.Host{
		Name:           h.Name,
		HostId:         h.ID().String(),
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		CreatedAt:      h.CreatedAt.Format(time.RFC3339),
	})
}

func (hc *HostController) Stop(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if err := hc.hs.Stop(h.ID()); err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not stop libp2p host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (hc *HostController) Bootstrap(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if err := h.Bootstrap(hc.ctx); err != nil {
		c.JSON(http.StatusInternalServerError, types.Error{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not bootstrap host",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusOK, &types.Host{
		Name:           h.Name,
		HostId:         h.ID().String(),
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		CreatedAt:      h.CreatedAt.Format(time.RFC3339),
	})
}
