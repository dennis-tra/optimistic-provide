package controller

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/ipfs/go-cid"

	"github.com/gin-gonic/gin"
	"github.com/libp2p/go-libp2p-core/peer"
	pb "github.com/libp2p/go-libp2p-kad-dht/pb"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/service"
)

// RPCController holds the API logic for all routes under /hosts
type RPCController struct {
	ctx       context.Context
	hs        service.HostService
	gpService service.GetProvidersService
}

// NewRPCController initializes a new host controller with the provided services.
func NewRPCController(ctx context.Context, hs service.HostService, gpService service.GetProvidersService) *RPCController {
	return &RPCController{
		ctx:       ctx,
		hs:        hs,
		gpService: gpService,
	}
}

// GetProviders runs a GET_PROVIDERS RPC against the given host for the given target key.
func (rpc *RPCController) GetProviders(c *gin.Context) {
	h := c.MustGet("host").(*dht.Host)

	if h.StartedAt == nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeHOSTSTOPPED,
			Message: "Host is stopped. Start it first to save a snapshot",
		})
		return
	}

	gpr := &types.GetProvidersRequest{}
	if err := c.BindJSON(gpr); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "Could not issue GetProviders RPC because of a malformed JSON request",
			Details: types.ErrDetails(err),
		})
		return
	}

	remoteID, err := peer.Decode(gpr.RemoteId)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "RemoteID is malformed and cannot be decoded to a PeerID",
			Details: types.ErrDetails(err),
		})
		return
	}

	contentID, err := cid.Decode(gpr.ContentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Code:    types.ErrorCodeMALFORMEDREQUEST,
			Message: "ContentID is malformed and cannot be decoded to a CID",
			Details: types.ErrDetails(err),
		})
		return
	}

	pm, err := pb.NewProtocolMessenger(h.MsgSender)
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: fmt.Sprintf("could not initialize protocol messanger"),
			Details: types.ErrDetails(err),
		})
		return
	}

	start := time.Now()
	providers, closerPeers, err := pm.GetProviders(c.Request.Context(), remoteID, contentID.Hash())
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: fmt.Sprintf("could not initialize protocol messenger"),
			Details: types.ErrDetails(err),
		})
		return
	}

	gps := &service.GetProvidersSpan{
		RemotePeerID: remoteID,
		Start:        start,
		End:          time.Now(),
		Providers:    providers,
		Error:        err,
	}

	if _, err = rpc.gpService.Save(c.Request.Context(), boil.GetContextDB(), h, []*service.GetProvidersSpan{gps}); err != nil {
		c.JSON(http.StatusPreconditionFailed, types.ErrorResponse{
			Code:    types.ErrorCodeINTERNAL,
			Message: fmt.Sprintf("could not save get providers to db"),
			Details: types.ErrDetails(err),
		})
		return
	}

	resp := &types.GetProvidersResponse{
		Providers:   make([]types.AddrInfo, len(providers)),
		CloserPeers: make([]types.AddrInfo, len(closerPeers)),
	}

	for i, p := range providers {
		maddrs := make([]string, len(p.Addrs))
		for j, m := range p.Addrs {
			maddrs[j] = m.String()
		}
		resp.Providers[i] = types.AddrInfo{
			PeerID:         p.ID.String(),
			MultiAddresses: maddrs,
		}
	}

	for i, p := range closerPeers {
		maddrs := make([]string, len(p.Addrs))
		for j, m := range p.Addrs {
			maddrs[j] = m.String()
		}
		resp.CloserPeers[i] = types.AddrInfo{
			PeerID:         p.ID.String(),
			MultiAddresses: maddrs,
		}
	}

	c.JSON(http.StatusOK, resp)
}
