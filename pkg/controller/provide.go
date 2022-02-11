package controller

import (
	"context"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/host"
	"github.com/dennis-tra/optimistic-provide/pkg/services"
)

type ProvideController struct {
	ctx context.Context
	dbc *db.Client

	hostsLk sync.RWMutex
	hosts   map[string]*host.Host

	rts *services.RoutingTableService
	hs  *services.HostService
}

func NewProvideController(ctx context.Context, dbc *db.Client) *ProvideController {
	return &ProvideController{
		ctx:   ctx,
		dbc:   dbc,
		rts:   services.NewRoutingTableService(dbc),
		hosts: map[string]*host.Host{},
	}
}

func (pc *ProvideController) SetRoutingTableService(rts *services.RoutingTableService) {
	pc.rts = rts
}

func (pc *ProvideController) SetHostService(hs *services.HostService) {
	pc.hs = hs
}

func (pc *ProvideController) Create(c *gin.Context) {
	pc.hostsLk.Lock()
	defer pc.hostsLk.Unlock()

	h, err := host.New(pc.ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pc.hosts[h.Host.ID().String()] = h

	c.JSON(http.StatusOK, h)
}

func (pc *ProvideController) Provide(c *gin.Context) {
	//h, code := pc.getHost(c)
	//if code != 0 {
	//	c.Status(code)
	//	return
	//}
	//
	//rts, err := pc.rts.SaveRoutingTable(context.Background(), h)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//dbPeer, err := pc.dbc.UpsertLocalPeer(h.PeerID)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//content, err := services.NewRandomContent()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//provide := &models.Provide{
	//	ProviderID:            dbPeer.ID,
	//	ContentID:             content.CID.String(),
	//	InitialRoutingTableID: rts.ID,
	//	StartedAt:             time.Now(),
	//}
	//if err = provide.Insert(context.Background(), pc.dbc, boil.Infer()); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//go func() {
	//	h.DHT.Provide(context.Background(), content.CID, false)
	//}()
	//
	//type ProvideResponse struct {
	//	ProvideID   int    `json:"provideId"`
	//	ContentID   string `json:"contentId"`
	//	ProviderID  string `json:"providerId"`
	//	RtInitialID int    `json:"rtInitialId,omitempty"`
	//	RtFinalID   int    `json:"rtFinalId,omitempty"`
	//}

	c.JSON(http.StatusOK, map[string]string{})
}
