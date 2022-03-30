package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/util"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type FindNodesService interface {
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, fnReqs []*FindNodesSpan) (models.FindNodesRPCSlice, error)
}

var _ FindNodesService = &FindNodes{}

type FindNodes struct {
	peerService PeerService
	maService   MultiAddressService
	fnRepo      repo.FindNodesRPCRepo
	cpRepo      repo.CloserPeersRepo
}

func NewFindNodesService(peerService PeerService, maService MultiAddressService, fnRepo repo.FindNodesRPCRepo, cpRepo repo.CloserPeersRepo) FindNodesService {
	return &FindNodes{
		peerService: peerService,
		maService:   maService,
		fnRepo:      fnRepo,
		cpRepo:      cpRepo,
	}
}

func (fn *FindNodes) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, fnReqs []*FindNodesSpan) (models.FindNodesRPCSlice, error) {
	log.Info("Saving Find Nodes RPCs")

	dbFindNodesRPCs := make([]*models.FindNodesRPC, len(fnReqs))
	for i, fnReq := range fnReqs {

		remotePeer, err := fn.peerService.UpsertPeer(ctx, exec, h, fnReq.RemotePeerID)
		if err != nil {
			return nil, err
		}

		dbfn := &models.FindNodesRPC{
			QueryID:          fnReq.QueryID.String(),
			LocalID:          h.DBHost.PeerID,
			RemoteID:         remotePeer.ID,
			StartedAt:        fnReq.Start,
			EndedAt:          fnReq.End,
			Error:            null.StringFromPtr(util.ErrorStr(fnReq.Error)),
			CloserPeersCount: null.NewInt(len(fnReq.CloserPeers), fnReq.Error == nil),
		}

		if err = dbfn.Insert(ctx, exec, boil.Infer()); err != nil {
			return nil, err
		}

		dbcps := make([]*models.CloserPeer, len(fnReq.CloserPeers))
		for j, closerPeer := range fnReq.CloserPeers {
			cp, err := fn.peerService.UpsertPeer(ctx, exec, h, closerPeer.ID)
			if err != nil {
				return nil, err
			}

			maids, err := fn.maService.UpsertMultiAddresses(ctx, exec, closerPeer.Addrs)
			if err != nil {
				return nil, err
			}

			dbcp := &models.CloserPeer{
				PeerID:          cp.ID,
				MultiAddressIds: maids,
			}

			dbcps[j] = dbcp
		}

		if err = dbfn.AddFindNodeRPCCloserPeers(ctx, exec, true, dbcps...); err != nil {
			return nil, err
		}

		dbFindNodesRPCs[i] = dbfn
	}

	return dbFindNodesRPCs, nil
}
