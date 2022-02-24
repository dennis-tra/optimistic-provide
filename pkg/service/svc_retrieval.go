package service

import (
	"context"
	"sync"
	"time"

	"github.com/volatiletech/null/v8"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-kad-dht/qpeerset"

	ks "github.com/whyrusleeping/go-keyspace"

	"github.com/ipfs/go-cid"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type RetrievalService interface {
	Retrieve(ctx context.Context, h *dht.Host, id cid.Cid, count int) (*models.Retrieval, error)
}

var _ RetrievalService = &Retrieval{}

type Retrieval struct {
	peerService  PeerService
	hostService  HostService
	rtService    RoutingTableService
	maService    MultiAddressService
	dialService  DialService
	connService  ConnectionService
	gpService    GetProvidersService
	psService    PeerStateService
	retrieveRepo repo.RetrievalRepo
}

func NewRetrievalService(peerService PeerService, hostService HostService, rtService RoutingTableService, maService MultiAddressService, dialService DialService, connService ConnectionService, gpService GetProvidersService, psService PeerStateService, apService AddRetrievalrsService, retrieveRepo repo.RetrievalRepo) *Retrieval {
	return &Retrieval{
		peerService:  peerService,
		hostService:  hostService,
		rtService:    rtService,
		maService:    maService,
		dialService:  dialService,
		connService:  connService,
		gpService:    gpService,
		psService:    psService,
		retrieveRepo: retrieveRepo,
	}
}

func (rs *Retrieval) Retrieve(ctx context.Context, h *dht.Host, contentID cid.Cid, count int) (*models.Retrieval, error) {
	rts, err := rs.rtService.Save(ctx, h)
	if err != nil {
		return nil, err
	}

	retrieval := &models.Retrieval{
		RetrieverID:           h.DBPeer.ID,
		ContentID:             contentID.String(),
		Distance:              ks.XORKeySpace.Key([]byte(h.ID())).Distance(ks.XORKeySpace.Key(contentID.Hash())).Bytes(),
		InitialRoutingTableID: rts.ID,
		StartedAt:             time.Now(),
	}
	go rs.startRetrieving(h, retrieval, contentID, count)

	return retrieval, nil
}

func (rs *Retrieval) startRetrieving(h *dht.Host, retrieval *models.Retrieval, contentID cid.Cid, count int) {
	ctx := context.Background()

	state := &RetrievalState{
		h:                    h,
		content:              contentID,
		dialsLk:              sync.RWMutex{},
		dials:                []*DialSpan{},
		getProvidersLk:       sync.RWMutex{},
		getProviders:         []*GetProvidersSpan{},
		connectionsStartedLk: sync.RWMutex{},
		connectionsStarted:   map[peer.ID]time.Time{},
		connectionsLk:        sync.RWMutex{},
		connections:          []*ConnectionSpan{},
		relevantPeers:        sync.Map{},
		peerSet:              qpeerset.NewQueryPeerset(string(contentID.Hash())),
	}

	ctx = state.Register(ctx)
	h.DHT.FindProvidersAsync(ctx, contentID, count)
	end := time.Now()

	rts, err := rs.rtService.Save(context.Background(), h)
	if err != nil {
		log.Warn(err)
	}

	retrieval.Error = nullStringFromErr(err)
	retrieval.EndedAt = null.TimeFrom(end)
	retrieval.FinalRoutingTableID = null.IntFrom(rts.ID)
	if retrieval, err = rs.retrieveRepo.Update(context.Background(), retrieval); err != nil {
		log.Warn(err)
	}

	if err = rs.dialService.Save(context.Background(), h.Host, retrieval.ID, state.dials); err != nil {
		log.Warn(err)
	}

	if err = rs.connService.Save(context.Background(), h.Host, retrieval.ID, state.connections); err != nil {
		log.Warn(err)
	}

	if err = rs.gpService.Save(context.Background(), h.Host, retrieval.ID, state.getProviders); err != nil {
		log.Warn(err)
	}

	if err = rs.psService.Save(context.Background(), h.Host, retrieval.ID, state.peerSet.AllStates()); err != nil {
		log.Warn(err)
	}
}
