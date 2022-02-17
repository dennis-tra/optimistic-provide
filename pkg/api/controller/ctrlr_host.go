package controller

import (
	"context"
	"net/http"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/api/render"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

// HostController holds the API logic for all routes in /hosts
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
	h, err := hc.hs.Create(hc.ctx)
	if err != nil {
		render.InternalServerError(c, err, "Could not start libp2p host")
		return
	}

	render.OK(c, h)
}

func (hc *HostController) List(c *gin.Context) {
	render.OK(c, hc.hs.Hosts())
}

func (hc *HostController) Get(c *gin.Context) {
	h, code := hc.getHost(c)
	if code != 0 {
		c.Status(code)
		return
	}

	render.OK(c, h)
}

func (hc *HostController) Stop(c *gin.Context) {
	peerID, code := hc.getPeerID(c)
	if code != 0 {
		c.Status(code)
		return
	}

	if err := hc.hs.Stop(peerID); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (hc *HostController) Bootstrap(c *gin.Context) {
	h, code := hc.getHost(c)
	if code != 0 {
		c.Status(code)
		return
	}

	if err := h.Bootstrap(hc.ctx); err != nil {
		render.InternalServerError(c, err, "Bootstrapping failed")
		return
	}

	render.OK(c, h)
}

func (hc *HostController) getPeerID(c *gin.Context) (peer.ID, int) {
	param, ok := c.Params.Get("peerID")
	if !ok {
		return "", http.StatusInternalServerError
	}

	peerID, err := peer.Decode(param)
	if err != nil {
		return "", http.StatusBadRequest
	}
	return peerID, 0
}

func (hc *HostController) getHost(c *gin.Context) (*dht.Host, int) {
	peerID, code := hc.getPeerID(c)
	if code != 0 {
		return nil, code
	}

	h, found := hc.hs.Host(peerID)
	if !found {
		return nil, http.StatusNotFound
	}
	return h, 0
}

func (hc *HostController) RefreshRoutingTable(c *gin.Context) {
	peerID, code := hc.getPeerID(c)
	if code != 0 {
		c.Status(code)
		return
	}

	if err := hc.hs.RefreshRoutingTableAsync(hc.ctx, peerID); err != nil {
		render.InternalServerError(c, err, "")
		return
	}

	c.Status(http.StatusOK)
}

func (hc *HostController) SaveRoutingTable(c *gin.Context) {
	peerID, code := hc.getPeerID(c)
	if code != 0 {
		c.Status(code)
		return
	}

	rts, err := hc.hs.SaveRoutingTable(hc.ctx, peerID)
	if err != nil {
		render.InternalServerError(c, err, "Saving routing table failed")
		return
	}

	render.OK(c, rts)
}