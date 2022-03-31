package controller

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"sort"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
	ks "github.com/whyrusleeping/go-keyspace"

	"golang.org/x/sync/errgroup"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

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

	if h.Host == nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTSTOPPED,
			Message: "Host is stopped. Start it first to bootstrap",
		})
		return
	}

	pr := &types.ProvideRequest{}
	if err := c.BindJSON(pr); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could start provide because of a malformed JSON request",
			Details: types.ErrDetails(err),
		})
		return
	}

	provide, err := pc.ps.Provide(pc.ctx, h, service.ProvideType(pr.Type))
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
		HostId:                h.PeerID(),
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
			HostId:                h.PeerID(),
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

	c.JSON(http.StatusCreated, types.ProvideDetails{
		Provide: types.Provide{
			ContentId:             dbProvide.ContentID,
			EndedAt:               util.TimeToStr(dbProvide.EndedAt.Ptr()),
			Error:                 dbProvide.Error.Ptr(),
			FinalRoutingTableId:   dbProvide.FinalRoutingTableID.Ptr(),
			HostId:                h.PeerID(),
			InitialRoutingTableId: dbProvide.InitialRoutingTableID,
			ProvideId:             dbProvide.ID,
			StartedAt:             dbProvide.StartedAt.Format(time.RFC3339Nano),
		},
		AddProvidersCount: len(dbProvide.R.AddProviderRPCS),
		ConnectionsCount:  len(dbProvide.R.Connections),
		DialsCount:        len(dbProvide.R.Dials),
		Distance:          fmt.Sprintf("0x%x", new(big.Int).SetBytes(dbProvide.Distance)),
		FindNodesCount:    len(dbProvide.R.FindNodesRPCS),
	})
}

func (pc *ProvideController) Graph(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)
	provideID := c.MustGet("provideID").(int)

	dbProvide, err := pc.ps.GetByID(c.Request.Context(), provideID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Error getting provide",
			Details: types.ErrDetails(err),
		})
		return
	}

	exec := boil.GetContextDB()

	errg, ectx := errgroup.WithContext(c.Request.Context())
	var dbDials models.DialSlice
	var dbConns models.ConnectionSlice
	var dbFindNodesRPCs models.FindNodesRPCSlice
	var dbAddProviderRPCs models.AddProviderRPCSlice
	var dbPeerStates models.PeerStateSlice

	errg.Go(func() error {
		var innerErr error
		dbDials, innerErr = dbProvide.Dials(
			qm.Load(models.DialRels.Remote),
			qm.Load(models.DialRels.MultiAddress),
		).All(ectx, exec)
		return innerErr
	})
	errg.Go(func() error {
		var innerErr error
		dbConns, innerErr = dbProvide.Connections(
			qm.Load(models.ConnectionRels.Remote),
			qm.Load(models.ConnectionRels.MultiAddress),
		).All(ectx, exec)
		return innerErr
	})
	errg.Go(func() error {
		var innerErr error
		dbFindNodesRPCs, innerErr = dbProvide.FindNodesRPCS(
			qm.Load(models.FindNodesRPCRels.Remote),
			qm.Load(models.FindNodesRPCRels.FindNodeRPCCloserPeers),
		).All(ectx, exec)
		return innerErr
	})
	errg.Go(func() error {
		var innerErr error
		dbAddProviderRPCs, innerErr = dbProvide.AddProviderRPCS(
			qm.Load(models.AddProviderRPCRels.Remote),
		).All(ectx, exec)
		return innerErr
	})
	errg.Go(func() error {
		var innerErr error
		dbPeerStates, innerErr = dbProvide.PeerStates(
			qm.Load(models.PeerStateRels.Peer),
			qm.Load(models.PeerStateRels.Referrer),
		).All(ectx, exec)
		return innerErr
	})

	if err = errg.Wait(); err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: "Error getting provide details",
			Details: types.ErrDetails(err),
		})
		return
	}
	allPeers := map[string]time.Time{}

	dials := make([]types.Dial, len(dbDials))
	for i, dbDial := range dbDials {
		if t, found := allPeers[dbDial.R.Remote.MultiHash]; !found || dbDial.StartedAt.Before(t) {
			allPeers[dbDial.R.Remote.MultiHash] = dbDial.StartedAt
		}

		dials[i] = types.Dial{
			Id:           dbDial.ID,
			DurationInS:  float32(dbDial.EndedAt.Sub(dbDial.StartedAt).Seconds()),
			EndedAt:      dbDial.EndedAt.Format(time.RFC3339Nano),
			Error:        dbDial.Error.Ptr(),
			MultiAddress: dbDial.R.MultiAddress.Maddr,
			RemoteId:     dbDial.R.Remote.MultiHash,
			StartedAt:    dbDial.StartedAt.Format(time.RFC3339Nano),
			Transport:    dbDial.Transport,
		}
	}

	conns := make([]types.Connection, len(dbConns))
	for i, dbConn := range dbConns {
		if t, found := allPeers[dbConn.R.Remote.MultiHash]; !found || dbConn.StartedAt.Before(t) {
			allPeers[dbConn.R.Remote.MultiHash] = dbConn.StartedAt
		}
		conns[i] = types.Connection{
			Id:           dbConn.ID,
			DurationInS:  float32(dbConn.EndedAt.Sub(dbConn.StartedAt).Seconds()),
			EndedAt:      dbConn.EndedAt.Format(time.RFC3339Nano),
			MultiAddress: dbConn.R.MultiAddress.Maddr,
			RemoteId:     dbConn.R.Remote.MultiHash,
			StartedAt:    dbConn.StartedAt.Format(time.RFC3339Nano),
		}
	}

	addProviders := make([]types.AddProvider, len(dbAddProviderRPCs))
	for i, dbAddProviderRPC := range dbAddProviderRPCs {
		if t, found := allPeers[dbAddProviderRPC.R.Remote.MultiHash]; !found || dbAddProviderRPC.StartedAt.Before(t) {
			allPeers[dbAddProviderRPC.R.Remote.MultiHash] = dbAddProviderRPC.StartedAt
		}
		addProviders[i] = types.AddProvider{
			Distance:    fmt.Sprintf("0x%x", new(big.Int).SetBytes(dbAddProviderRPC.Distance)),
			DurationInS: float32(dbAddProviderRPC.EndedAt.Sub(dbAddProviderRPC.StartedAt).Seconds()),
			EndedAt:     dbAddProviderRPC.EndedAt.Format(time.RFC3339Nano),
			Error:       dbAddProviderRPC.Error.Ptr(),
			Id:          dbAddProviderRPC.ID,
			RemoteId:    dbAddProviderRPC.R.Remote.MultiHash,
			StartedAt:   dbAddProviderRPC.StartedAt.Format(time.RFC3339Nano),
		}
	}

	findNodes := make([]types.FindNode, len(dbFindNodesRPCs))
	for i, dbFindNodesRPC := range dbFindNodesRPCs {
		if t, found := allPeers[dbFindNodesRPC.R.Remote.MultiHash]; !found || dbFindNodesRPC.StartedAt.Before(t) {
			allPeers[dbFindNodesRPC.R.Remote.MultiHash] = dbFindNodesRPC.StartedAt
		}

		dbCloserPeers, err := dbFindNodesRPC.FindNodeRPCCloserPeers(qm.Load(models.CloserPeerRels.Peer)).All(c.Request.Context(), exec)
		if err != nil {
			c.JSON(http.StatusInternalServerError, types.ErrorResponse{
				Code:    types.ErrorCodeINTERNAL,
				Message: "Error getting closer peers",
				Details: types.ErrDetails(err),
			})
			return
		}

		closerPeers := make([]types.FindNodeCloserPeer, len(dbFindNodesRPC.R.FindNodeRPCCloserPeers))
		for j, dbCloserPeer := range dbCloserPeers {
			p1, err1 := peer.Decode(dbFindNodesRPC.R.Remote.MultiHash)
			p2, err2 := peer.Decode(dbCloserPeer.R.Peer.MultiHash)
			if err1 != nil || err2 != nil {
				c.JSON(http.StatusInternalServerError, types.ErrorResponse{
					Code:    types.ErrorCodeINTERNAL,
					Message: fmt.Sprintf("Error decoding multi hashes %q or %q", dbFindNodesRPC.R.Remote.MultiHash, dbCloserPeer.R.Peer.MultiHash),
					Details: types.ErrDetails(err),
				})
				return
			}
			contentID, err := cid.Decode(dbProvide.ContentID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, types.ErrorResponse{
					Code:    types.ErrorCodeINTERNAL,
					Message: "error decoding CID",
					Details: types.ErrDetails(err),
				})
				return
			}
			closerPeers[j] = types.FindNodeCloserPeer{
				Bucket:   int(util.BucketIdForPeer(p1, p2)),
				Distance: fmt.Sprintf("0x%x", new(big.Int).SetBytes(ks.XORKeySpace.Key([]byte(dbCloserPeer.R.Peer.MultiHash)).Distance(ks.XORKeySpace.Key(contentID.Hash())).Bytes())),
				PeerId:   dbCloserPeer.R.Peer.MultiHash,
			}
		}

		findNodes[i] = types.FindNode{
			DurationInS: float32(dbFindNodesRPC.EndedAt.Sub(dbFindNodesRPC.StartedAt).Seconds()),
			EndedAt:     dbFindNodesRPC.EndedAt.Format(time.RFC3339Nano),
			Error:       dbFindNodesRPC.Error.Ptr(),
			Id:          dbFindNodesRPC.ID,
			QueryId:     dbFindNodesRPC.QueryID,
			RemoteId:    dbFindNodesRPC.R.Remote.MultiHash,
			StartedAt:   dbFindNodesRPC.StartedAt.Format(time.RFC3339Nano),
			CloserPeers: closerPeers,
		}
	}

	peerInfos := []types.ProvidePeerInfo{}
	for _, dbPeerState := range dbPeerStates {
		firstInteractedAt := allPeers[dbPeerState.R.Peer.MultiHash]
		peerInfos = append(peerInfos, types.ProvidePeerInfo{
			AgentVersion:      dbPeerState.R.Peer.AgentVersion.Ptr(),
			Distance:          fmt.Sprintf("0x%x", new(big.Int).SetBytes(dbPeerState.Distance)),
			FirstInteractedAt: util.TimeToStr(&firstInteractedAt),
			PeerId:            dbPeerState.R.Peer.MultiHash,
			Protocols:         dbPeerState.R.Peer.Protocols,
			ReferredBy:        dbPeerState.R.Referrer.MultiHash,
			State:             types.QueryPeerState(dbPeerState.State),
		})
	}

	sort.Slice(peerInfos, func(i, j int) bool {
		t1Str, t2Str := peerInfos[i].FirstInteractedAt, peerInfos[j].FirstInteractedAt
		if t1Str == nil && t2Str == nil {
			return false
		} else if t1Str == nil && t2Str != nil {
			return false
		} else if t1Str != nil && t2Str == nil {
			return true
		} else if t1Str != nil && t2Str != nil {
			t1, _ := time.Parse(*t1Str, time.RFC3339Nano)
			t2, _ := time.Parse(*t2Str, time.RFC3339Nano)
			return t1.Before(t2)
		}
		panic("AAAHH")
	})

	c.JSON(http.StatusOK, types.ProvideGraph{
		Provide: types.Provide{
			ContentId:             dbProvide.ContentID,
			EndedAt:               util.TimeToStr(dbProvide.EndedAt.Ptr()),
			Error:                 dbProvide.Error.Ptr(),
			FinalRoutingTableId:   dbProvide.FinalRoutingTableID.Ptr(),
			HostId:                h.PeerID(),
			InitialRoutingTableId: dbProvide.InitialRoutingTableID,
			ProvideId:             dbProvide.ID,
			StartedAt:             dbProvide.StartedAt.Format(time.RFC3339Nano),
		},
		AddProviders: addProviders,
		Connections:  conns,
		Dials:        dials,
		FindNodes:    findNodes,
		Peers:        peerInfos,
	})
}
