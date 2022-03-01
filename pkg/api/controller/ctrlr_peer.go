package controller

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

// PeerController holds the API logic for all routes in /hosts
type PeerController struct {
	ctx context.Context
	ps  service.PeerService
}

// NewPeerController initializes a new host controller with the provided services.
func NewPeerController(ctx context.Context, ps service.PeerService) *PeerController {
	return &PeerController{
		ctx: ctx,
		ps:  ps,
	}
}

func (pc *PeerController) Get(c *gin.Context) {
	pid := c.MustGet("peerID").(peer.ID)

	dbPeer, err := pc.ps.Find(c.Request.Context(), pid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, types.ErrorResponse{
				Code:    types.ErrorCodePEERNOTFOUND,
				Message: "Peer with ID " + pid.String() + " was not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Code:    types.ErrorCodeINTERNAL,
				Message: "Error retrieving peer with ID " + pid.String(),
				Details: types.ErrDetails(err),
			})
		}
		return
	}

	c.JSON(http.StatusOK, types.Peer{
		AgentVersion: dbPeer.AgentVersion.Ptr(),
		CreatedAt:    dbPeer.CreatedAt.Format(time.RFC3339Nano),
		PeerId:       dbPeer.MultiHash,
		Protocols:    dbPeer.Protocols,
	})
}
