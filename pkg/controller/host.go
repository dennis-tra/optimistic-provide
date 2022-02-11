package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/host"
	"github.com/dennis-tra/optimistic-provide/pkg/services"
)

type HostController struct {
	ctx context.Context
	dbc *db.Client

	rts *services.RoutingTableService
	hs  *services.HostService
}

func NewHostController(ctx context.Context, dbc *db.Client) *HostController {
	return &HostController{
		ctx: ctx,
		dbc: dbc,
	}
}

func (hc *HostController) SetRoutingTableService(rts *services.RoutingTableService) {
	hc.rts = rts
}

func (hc *HostController) SetHostService(hs *services.HostService) {
	hc.hs = hs
}

func (hc *HostController) Create(c *gin.Context) {
	h, err := hc.hs.Create(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h)
}

func (hc *HostController) List(c *gin.Context) {
	c.JSON(http.StatusOK, hc.hs.Hosts())
}

func (hc *HostController) Get(c *gin.Context) {
	h, code := hc.getHost(c)
	if code != 0 {
		c.Status(code)
		return
	}

	c.JSON(http.StatusOK, h)
}

func (hc *HostController) Stop(c *gin.Context) {
	param, ok := c.Params.Get("peerID")
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	peerID, err := peer.Decode(param)
	if err != nil {
		c.Status(http.StatusBadRequest)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h)
}

func (hc *HostController) getHost(c *gin.Context) (*host.Host, int) {
	param, ok := c.Params.Get("peerID")
	if !ok {
		return nil, http.StatusInternalServerError
	}

	peerID, err := peer.Decode(param)
	if err != nil {
		return nil, http.StatusBadRequest
	}

	h, found := hc.hs.Host(peerID)
	if !found {
		return nil, http.StatusNotFound
	}
	return h, 0
}

func (hc *HostController) RefreshRoutingTable(c *gin.Context) {
	h, code := hc.getHost(c)
	if code != 0 {
		c.Status(code)
		return
	}

	go h.RefreshRoutingTable(hc.ctx)

	c.Status(http.StatusOK)
}

func (hc *HostController) SaveRoutingTable(c *gin.Context) {
	h, code := hc.getHost(c)
	if code != 0 {
		c.Status(code)
		return
	}

	if _, err := hc.rts.SaveRoutingTable(context.Background(), h); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
