package service

import (
	"context"

	"github.com/volatiletech/null/v8"

	"github.com/libp2p/go-libp2p-kad-dht/qpeerset"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/libp2p/go-libp2p-core/host"
)

type PeerStateService interface {
	Save(ctx context.Context, h host.Host, op HostOperation, id int, states []qpeerset.QueryPeerState) error
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

func (ps *PeerState) Save(ctx context.Context, h host.Host, op HostOperation, id int, states []qpeerset.QueryPeerState) error {
	log.Info("Saving connections...")
	defer log.Info("Done saving connections")

	for _, state := range states {
		remotePeer, err := ps.peerService.UpsertPeer(h, state.ID)
		if err != nil {
			return err
		}
		referrerPeer, err := ps.peerService.UpsertPeer(h, state.ReferredBy)
		if err != nil {
			return err
		}
		pState := &models.PeerState{
			PeerID:     remotePeer.ID,
			ReferrerID: referrerPeer.ID,
			State:      ps.mapState(state.State),
			Distance:   state.Distance.Bytes(),
		}

		switch op {
		case HostOperationProvide:
			pState.ProvideID = null.IntFrom(id)
		case HostOperationRetrieval:
			pState.RetrievalID = null.IntFrom(id)
		default:
			panic(op)
		}

		if pState, err = ps.psRepo.Save(ctx, pState); err != nil {
			return err
		}
	}

	return nil
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
