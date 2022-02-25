package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/libp2p/go-libp2p-kad-dht/qpeerset"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type PeerStateService interface {
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, states []qpeerset.QueryPeerState) (models.PeerStateSlice, error)
}

var _ PeerStateService = &PeerState{}

type PeerState struct {
	peerService PeerService
	psRepo      repo.PeerStateRepo
}

func NewPeerStateService(peerService PeerService, psRepo repo.PeerStateRepo) PeerStateService {
	return &PeerState{
		peerService: peerService,
		psRepo:      psRepo,
	}
}

func (ps *PeerState) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, states []qpeerset.QueryPeerState) (models.PeerStateSlice, error) {
	log.Info("Saving Peer State")

	dbStates := make([]*models.PeerState, len(states))
	for i, state := range states {
		remotePeer, err := ps.peerService.UpsertPeer(ctx, exec, h, state.ID)
		if err != nil {
			return nil, err
		}

		referrerPeer, err := ps.peerService.UpsertPeer(ctx, exec, h, state.ReferredBy)
		if err != nil {
			return nil, err
		}

		pState := &models.PeerState{
			PeerID:     remotePeer.ID,
			ReferrerID: referrerPeer.ID,
			State:      ps.mapState(state.State),
			Distance:   state.Distance.Bytes(),
		}

		if err = pState.Insert(ctx, exec, boil.Infer()); err != nil {
			return nil, err
		}

		dbStates[i] = pState
	}

	return dbStates, nil
}

func (ps *PeerState) mapState(state qpeerset.PeerState) string {
	switch state {
	case qpeerset.PeerHeard:
		return models.PeerStateHEARD
	case qpeerset.PeerWaiting:
		return models.PeerStateWAITING
	case qpeerset.PeerQueried:
		return models.PeerStateQUERIED
	case qpeerset.PeerUnreachable:
		return models.PeerStateUNREACHABLE
	default:
		panic(state)
	}
}
