package controller

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/dennis-tra/optimistic-provide/pkg/lib"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/db/models"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/services"
)

type ProvideController struct {
	ctx context.Context
	dbc *db.Client

	hostsLk sync.RWMutex
	hosts   map[string]*dht.Host

	ps  *services.ProvideService
	rts *services.RoutingTableService
	hs  *services.HostService
}

func NewProvideController(ctx context.Context, dbc *db.Client, ps *services.ProvideService, rts *services.RoutingTableService, hs *services.HostService) *ProvideController {
	return &ProvideController{
		ctx:   ctx,
		dbc:   dbc,
		ps:    ps,
		rts:   rts,
		hs:    hs,
		hosts: map[string]*dht.Host{},
	}
}

type CreateProvideRequest struct {
	HostID string `json:"hostId"`
}

type CreateProvideResponse struct {
	ProviderID            string    `json:"providerId"`
	ProvideID             int       `json:"provideId"`
	ContentID             string    `json:"contentId"`
	InitialRoutingTableID int       `json:"initialRoutingTableId"`
	StartedAt             time.Time `json:"startedAt"`
}

func (pc *ProvideController) Create(c *gin.Context) {
	// TODO: check pending provide
	req := CreateProvideRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	hostID, err := peer.Decode(req.HostID)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	h, found := pc.hs.Host(hostID)
	if !found {
		c.Status(http.StatusNotFound)
		return
	}

	rts, err := pc.rts.SaveRoutingTable(context.Background(), h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	content, err := lib.NewRandomContent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	provide := &models.Provide{
		ProviderID:            h.DBPeer.ID,
		ContentID:             content.CID.String(),
		InitialRoutingTableID: rts.ID,
		StartedAt:             time.Now(),
	}

	if err := provide.Insert(context.Background(), pc.dbc, boil.Infer()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	go pc.ps.Provide(h, provide, content)

	resp := CreateProvideResponse{
		ProviderID:            h.PeerID.String(),
		ProvideID:             provide.ID,
		ContentID:             content.CID.String(),
		InitialRoutingTableID: rts.ID,
		StartedAt:             time.Now(),
	}

	c.JSON(http.StatusOK, resp)
}

func (pc *ProvideController) SetProvideService(ps *services.ProvideService) {
	pc.ps = ps
}
