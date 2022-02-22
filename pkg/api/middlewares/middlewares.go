package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

func HostID(hs service.HostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		param, ok := c.Params.Get("hostID")
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.Error{
				Code:    types.ErrorCodeINTERNAL,
				Message: "Could not get host peer ID from endpoint path.",
			})
			return
		}

		peerID, err := peer.Decode(param)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, types.Error{
				Code:    types.ErrorCodeMALFORMEDPEERID,
				Message: "Could not decode host peer ID: " + param,
				Details: types.ErrDetails(err),
			})
			return
		}

		h, found := hs.Host(peerID)
		if !found {
			c.AbortWithStatusJSON(http.StatusNotFound, types.Error{
				Code:    types.ErrorCodeHOSTNOTFOUND,
				Message: "Host with ID " + peerID.String() + " was not found.",
			})
			return
		}

		c.Set("host", h)

		c.Next()
	}
}

func RoutingTableID(c *gin.Context) {
	param, ok := c.Params.Get("routingTableID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.Error{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not get routing table ID from endpoint path.",
		})
		return
	}

	routingTableID, err := strconv.Atoi(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.Error{
			Code:    types.ErrorCodeROUTINGTABLENOTFOUND,
			Message: "Could not convert " + param + "to integer",
		})
		return
	}

	c.Set("routingTableID", routingTableID)

	c.Next()
}

func PeerID(c *gin.Context) {
	param, ok := c.Params.Get("peerID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, types.Error{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Could not get peer ID from endpoint path.",
		})
		return
	}

	peerID, err := peer.Decode(param)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.Error{
			Code:    types.ErrorCodeMALFORMEDPEERID,
			Message: "Could not decode peer ID: " + param,
			Details: types.ErrDetails(err),
		})
		return
	}

	c.Set("peerID", peerID)

	c.Next()
}
