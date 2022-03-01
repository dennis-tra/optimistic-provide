package controller

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/dennis-tra/optimistic-provide/pkg/util"

	"github.com/gin-gonic/gin"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

type ProvideController struct {
	ctx context.Context
	dbc *db.Client
	ps  service.ProvideService
	hs  service.HostService
}

func NewProvideController(ctx context.Context, ps service.ProvideService, hs service.HostService) *ProvideController {
	return &ProvideController{
		ctx: ctx,
		ps:  ps,
		hs:  hs,
	}
}

func (pc *ProvideController) Create(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	pr := &types.ProvideRequest{}
	if err := c.BindJSON(pr); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could start provide because of a malformed JSON request",
			Details: types.ErrDetails(err),
		})
		return
	}

	provide, err := pc.ps.Provide(pc.ctx, h, pr.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Error starting provide operation",
			Details: types.ErrDetails(err),
		})
		return
	}

	c.JSON(http.StatusCreated, types.Provide{
		ContentId:             provide.ContentID,
		EndedAt:               nil,
		Error:                 nil,
		FinalRoutingTableId:   nil,
		HostId:                h.ID().String(),
		InitialRoutingTableId: provide.InitialRoutingTableID,
		ProvideId:             provide.ID,
		StartedAt:             provide.StartedAt.Format(time.RFC3339Nano),
	})
}

func (pc *ProvideController) List(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	dbProvides, err := pc.ps.List(pc.ctx, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Error listing provide operations",
			Details: types.ErrDetails(err),
		})
		return
	}

	provides := make([]types.Provide, len(dbProvides))
	for i, dbProvide := range dbProvides {
		provides[i] = types.Provide{
			ContentId:             dbProvide.ContentID,
			EndedAt:               util.TimeToStr(dbProvide.EndedAt.Ptr()),
			Error:                 dbProvide.Error.Ptr(),
			FinalRoutingTableId:   dbProvide.FinalRoutingTableID.Ptr(),
			HostId:                h.ID().String(),
			InitialRoutingTableId: dbProvide.InitialRoutingTableID,
			ProvideId:             dbProvide.ID,
			StartedAt:             dbProvide.StartedAt.Format(time.RFC3339Nano),
		}
	}

	c.JSON(http.StatusCreated, provides)
}

func (pc *ProvideController) Get(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)
	provideID := c.MustGet("provideID").(int)

	dbProvide, err := pc.ps.Get(c.Request.Context(), h, provideID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Error getting provide",
			Details: types.ErrDetails(err),
		})
		return
	}

	connections := make([]types.Connection, len(dbProvide.R.Connections))
	for i, dbConn := range dbProvide.R.Connections {
		connections[i] = types.Connection{
			DurationInS: float32(dbConn.EndedAt.Sub(dbConn.StartedAt).Seconds()),
			EndedAt:     dbConn.EndedAt.Format(time.RFC3339Nano),
			Id:          dbConn.ID,
			// MultiAddress: dbConn.R.MultiAddress.Maddr,
			// RemoteId:     dbConn.R.Remote.MultiHash,
			StartedAt: dbConn.StartedAt.Format(time.RFC3339Nano),
		}
	}

	findNodes := make([]types.FindNode, len(dbProvide.R.FindNodesRPCS))
	for i, dbFindNode := range dbProvide.R.FindNodesRPCS {
		findNodes[i] = types.FindNode{
			CloserPeersCount: dbFindNode.CloserPeersCount.Ptr(),
			DurationInS:      float32(dbFindNode.EndedAt.Sub(dbFindNode.StartedAt).Seconds()),
			EndedAt:          dbFindNode.EndedAt.Format(time.RFC3339Nano),
			Error:            dbFindNode.Error.Ptr(),
			Id:               dbFindNode.ID,
			// RemoteId:         dbFindNode.R.Remote.MultiHash,
			StartedAt: dbFindNode.StartedAt.Format(time.RFC3339Nano),
		}
	}

	dials := make([]types.Dial, len(dbProvide.R.Dials))
	for i, dbDial := range dbProvide.R.Dials {
		dials[i] = types.Dial{
			DurationInS: float32(dbDial.EndedAt.Sub(dbDial.StartedAt).Seconds()),
			EndedAt:     dbDial.EndedAt.Format(time.RFC3339Nano),
			Error:       dbDial.Error.Ptr(),
			Id:          dbDial.ID,
			// MultiAddress: dbDial.R.MultiAddress.Maddr,
			// RemoteId:     dbDial.R.Remote.MultiHash,
			StartedAt: dbDial.StartedAt.Format(time.RFC3339Nano),
			Transport: dbDial.Transport,
		}
	}

	addProviders := make([]types.AddProvider, len(dbProvide.R.AddProviderRPCS))
	for i, dbAddProvider := range dbProvide.R.AddProviderRPCS {
		addProviders[i] = types.AddProvider{
			Distance:    base64.RawStdEncoding.EncodeToString(dbAddProvider.Distance),
			DurationInS: float32(dbAddProvider.EndedAt.Sub(dbAddProvider.StartedAt).Seconds()),
			EndedAt:     dbAddProvider.EndedAt.Format(time.RFC3339Nano),
			Error:       dbAddProvider.Error.Ptr(),
			Id:          dbAddProvider.ID,
			// RemoteId:    dbAddProvider.R.Remote.MultiHash,
			StartedAt: dbAddProvider.StartedAt.Format(time.RFC3339Nano),
		}
	}

	c.JSON(http.StatusCreated, types.ProvideDetails{
		Provide: types.Provide{
			ContentId:             dbProvide.ContentID,
			EndedAt:               util.TimeToStr(dbProvide.EndedAt.Ptr()),
			Error:                 dbProvide.Error.Ptr(),
			FinalRoutingTableId:   dbProvide.FinalRoutingTableID.Ptr(),
			HostId:                h.DBHost.R.Peer.MultiHash,
			InitialRoutingTableId: dbProvide.InitialRoutingTableID,
			ProvideId:             dbProvide.ID,
			StartedAt:             dbProvide.StartedAt.Format(time.RFC3339Nano),
		},
		Distance:     fmt.Sprintf("0x%x", new(big.Int).SetBytes(dbProvide.Distance)),
		Connections:  connections,
		FindNodes:    findNodes,
		Dials:        dials,
		AddProviders: addProviders,
	})
}
