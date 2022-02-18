package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/dennis-tra/optimistic-provide/pkg/api/render"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("optprov")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

// RoutingTableController holds the API logic for all routes under /hosts
type RoutingTableController struct {
	ctx    context.Context
	rts    service.RoutingTableService
	hs     service.HostService
	rtRepo repo.RoutingTableRepo
}

// NewRoutingTableController initializes a new host controller with the provided services.
func NewRoutingTableController(ctx context.Context, rts service.RoutingTableService, hs service.HostService, rtRepo repo.RoutingTableRepo) *RoutingTableController {
	return &RoutingTableController{
		ctx:    ctx,
		rts:    rts,
		hs:     hs,
		rtRepo: rtRepo,
	}
}

// Create starts a new libp2p host.
func (rtc *RoutingTableController) Create(c *gin.Context) {
	peerID, rerr := getHostID(c)
	if rerr != nil {
		render.Err(c, rerr)
		return
	}

	h, found := rtc.hs.Host(peerID)
	if !found {
		render.Err(c, render.NewErrNotFound(render.ErrorCodeHostNotFound, "Host with ID "+peerID.String()+" was not found.", nil))
		return
	}

	rts, err := rtc.rts.SaveRoutingTable(rtc.ctx, h)
	if err != nil {
		render.ErrInternal(c, "Saving routing table failed", err)
		return
	}

	render.OK(c, rts)
}

func (rtc *RoutingTableController) Get(c *gin.Context) {
	_, rerr := getHostID(c)
	if rerr != nil {
		render.Err(c, rerr)
		return
	}

	routingTableIDStr, ok := c.Params.Get("routingTableID")
	if !ok {
		render.Err(c, render.NewErrInternalServerError(render.ErrorCodeGetPeerFromPath, "Could not get host multi hash from endpoint path", nil))
		return
	}

	routingTableID, err := strconv.Atoi(routingTableIDStr)
	if rerr != nil {
		render.NewErrInternalServerError(render.ErrorCodeMalformedJSON, "Could not convert "+routingTableIDStr+"to integer", err)
		return
	}

	rts, err := rtc.rtRepo.Find(rtc.ctx, routingTableID)
	if rerr != nil {
		render.NewErrInternalServerError(render.ErrorCodeMalformedJSON, "Could not convert "+routingTableIDStr+"to integer", err)
		return
	}

	// TODO: check host ID
	render.OK(c, rts)
}

func (rtc *RoutingTableController) List(c *gin.Context) {
	hostID, rerr := getHostID(c)
	if rerr != nil {
		render.Err(c, rerr)
		return
	}

	rts, err := rtc.rtRepo.FindAll(rtc.ctx, hostID.String())
	if rerr != nil {
		render.NewErrInternalServerError(render.ErrorCodeMalformedJSON, "Could not convert to integer", err)
		return
	}
	render.OK(c, rts)
}

func (rtc *RoutingTableController) Listen(c *gin.Context) {
	hostID, rerr := getHostID(c)
	if rerr != nil {
		render.Err(c, rerr)
		return
	}

	h, found := rtc.hs.Host(hostID)
	if !found {
		render.Err(c, render.NewErrNotFound(render.ErrorCodeHostNotFound, "Host not found", nil))
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		render.ErrInternal(c, "Could not upgrade connection", err)
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
			data, err := update.Marshal()
			if err != nil {
				log.Warnf("Could not marshal routing table update %v: %s", update, err)
				continue
			}
			log.Infof("Sending %d bytes for %s", len(data), hostID.String())
			if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Warn("Could not write websocket message", err)
				continue
			}
		}
	}()

	rtl.SendUpdate()
}

func (rtc *RoutingTableController) Refresh(c *gin.Context) {
	hostID, rerr := getHostID(c)
	if rerr != nil {
		render.Err(c, rerr)
		return
	}

	h, found := rtc.hs.Host(hostID)
	if !found {
		render.Err(c, render.NewErrNotFound(render.ErrorCodeHostNotFound, "Host not found", nil))
		return
	}

	go func() {
		err := <-h.DHT.RefreshRoutingTable()
		log.Warn("Routing Table refresh", err)
	}()
	c.Status(http.StatusOK)
}
