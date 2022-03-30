package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	logging "github.com/ipfs/go-log"
	"github.com/pkg/errors"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

var log = logging.Logger("optprov")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

// RoutingTableController holds the API logic for all routes under /hosts
type RoutingTableController struct {
	ctx context.Context
	rts service.RoutingTableService
	hs  service.HostService
}

// NewRoutingTableController initializes a new host controller with the provided services.
func NewRoutingTableController(ctx context.Context, rts service.RoutingTableService, hs service.HostService) *RoutingTableController {
	return &RoutingTableController{
		ctx: ctx,
		rts: rts,
		hs:  hs,
	}
}

// Create starts a new libp2p host.
func (rtc *RoutingTableController) Create(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if h.StartedAt == nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTSTOPPED,
			Message: "Host is stopped. Start it first to save a snapshot",
		})
		return
	}

	rts, err := rtc.rts.SaveTxn(rtc.ctx, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeSAVINGROUTINGTABLE,
			Message: "Saving routing table for host " + h.PeerID() + "failed",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusCreated, types.RoutingTable{
		Id:         rts.ID,
		HostId:     h.PeerID(),
		BucketSize: rts.BucketSize,
		CreatedAt:  rts.CreatedAt.Format(time.RFC3339Nano),
		EntryCount: rts.EntryCount,
	})
}

func (rtc *RoutingTableController) Get(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)
	routingTableID := c.MustGet("routingTableID").(int)

	rts, err := rtc.rts.FindByIDAndHostID(rtc.ctx, routingTableID, h.ID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, types.ErrorResponse{
				Code:    types.ErrorCodeROUTINGTABLENOTFOUND,
				Message: fmt.Sprintf("Routing table with ID %d for host %s was not found", routingTableID, h.ID()),
			})
		} else {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Code:    types.ErrorCodeINTERNAL,
				Message: fmt.Sprintf("Error retrieving routing table snapshot for ID %d and host %s", routingTableID, h.ID()),
				Details: types.ErrDetails(err),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, types.RoutingTable{
		Id:         rts.ID,
		HostId:     h.PeerID(),
		BucketSize: rts.BucketSize,
		CreatedAt:  rts.CreatedAt.Format(time.RFC3339Nano),
		EntryCount: rts.EntryCount,
	})
}

func (rtc *RoutingTableController) List(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	rts, err := rtc.rts.FindAll(rtc.ctx, h.ID())
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: fmt.Sprintf("Error retrieving routing table snapshots for host %s", h.ID()),
			Details: types.ErrDetails(err),
		})
		return
	}

	snapshots := make([]types.RoutingTable, len(rts))
	for i, r := range rts {
		snapshots[i] = types.RoutingTable{
			BucketSize: r.BucketSize,
			CreatedAt:  r.CreatedAt.Format(time.RFC3339Nano),
			EntryCount: r.EntryCount,
			Id:         r.ID,
			HostId:     r.R.Peer.MultiHash,
		}
	}

	c.JSON(http.StatusOK, snapshots)
}

func (rtc *RoutingTableController) Current(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)
	if h.StartedAt == nil {
		c.JSON(http.StatusOK, []types.RoutingTablePeer{})
		return
	}
	rtl := service.NewRoutingTableListener(h)
	c.JSON(http.StatusOK, rtl.BuildUpdate())
}

func (rtc *RoutingTableController) Listen(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if h.StartedAt == nil {
		c.Status(http.StatusPreconditionFailed)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: fmt.Sprintf("Error could not upgrade connection"),
			Details: types.ErrDetails(err),
		})
		return
	}

	rtl := service.NewRoutingTableListener(h)
	h.RegisterRoutingTableListener(rtl)

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Infof("Stopped reading websocket")
				h.UnregisterRoutingTableListener(rtl)
				rtl.Stop()
				break
			}
			log.Infof("Received websocket message %s", msg)
		}
	}()

	go func() {
		for update := range rtl.Updates() {
			data, err := json.Marshal(update)
			if err != nil {
				log.Warnf("Could not marshal routing table update %v: %s", update, err)
				continue
			}
			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Warn("Could not write websocket message", err)
				continue
			}
		}
		if err = conn.Close(); err != nil {
			log.Warnf("Could not close websocket connection: %s", err)
		}
	}()

	rtl.SendFullUpdate()
}

func (rtc *RoutingTableController) Refresh(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	go func() {
		err := <-h.DHT.RefreshRoutingTable()
		log.Warn("Routing Table refresh", err)
	}()

	c.Status(http.StatusNoContent)
}
