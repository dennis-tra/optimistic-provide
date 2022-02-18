package controller

import (
	"context"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/api/render"
	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
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
	h, err := hc.hs.Create(hc.ctx)
	if err != nil {
		render.ErrInternal(c, "Could not create libp2p host", err)
		return
	}

	render.OK(c, &types.Host{
		HostID:         h.ID().String(),
		BootstrappedAt: nil,
		CreatedAt:      h.CreatedAt.Format(time.RFC3339),
	})
}

// List lists returns all running libp2p hosts.
func (hc *HostController) List(c *gin.Context) {
	hosts := []*types.Host{}
	for _, h := range hc.hs.Hosts() {
		hosts = append(hosts, &types.Host{
			HostID:         h.ID().String(),
			BootstrappedAt: util.TimeToStr(h.Bootstrapped),
			CreatedAt:      h.CreatedAt.Format(time.RFC3339),
		})
	}

	sort.Slice(hosts, func(i, j int) bool {
		ti, _ := time.Parse(time.RFC3339, hosts[i].CreatedAt)
		tj, _ := time.Parse(time.RFC3339, hosts[j].CreatedAt)
		return ti.Before(tj)
	})

	render.OK(c, hosts)
}

func (hc *HostController) Get(c *gin.Context) {
	h, err := hc.getHost(c)
	if err != nil {
		render.Err(c, err)
		return
	}

	render.OK(c, &types.Host{
		HostID:         h.ID().String(),
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		CreatedAt:      h.CreatedAt.Format(time.RFC3339),
	})
}

func (hc *HostController) Stop(c *gin.Context) {
	peerID, err := getHostID(c)
	if err != nil {
		render.Err(c, err)
		return
	}

	if err := hc.hs.Stop(peerID); err != nil {
		render.ErrInternal(c, "Could not stop libp2p host", err)
		return
	}
	c.Status(http.StatusOK)
}

func (hc *HostController) Bootstrap(c *gin.Context) {
	h, err := hc.getHost(c)
	if err != nil {
		render.Err(c, err)
		return
	}

	if err := h.Bootstrap(hc.ctx); err != nil {
		render.ErrInternal(c, "Could not bootstrap host", err)
		return
	}

	render.OK(c, &types.Host{
		HostID:         h.ID().String(),
		BootstrappedAt: util.TimeToStr(h.Bootstrapped),
		CreatedAt:      h.CreatedAt.Format(time.RFC3339),
	})
}

func getHostID(c *gin.Context) (peer.ID, *render.Error) {
	param, ok := c.Params.Get("hostID")
	if !ok {
		return "", render.NewErrInternalServerError(render.ErrorCodeGetPeerFromPath, "Could not get host multi hash from endpoint path", nil)
	}

	peerID, err := peer.Decode(param)
	if err != nil {
		return "", render.NewErrBadRequest(render.ErrorCodeMalformedPeerID, "Could not decode host peer ID: "+param, err)
	}

	return peerID, nil
}

func (hc *HostController) getHost(c *gin.Context) (*dht.Host, *render.Error) {
	peerID, err := getHostID(c)
	if err != nil {
		return nil, err
	}

	h, found := hc.hs.Host(peerID)
	if !found {
		return nil, render.NewErrNotFound(render.ErrorCodeHostNotFound, "Host with ID "+peerID.String()+" was not found.", nil)
	}

	return h, nil
}
