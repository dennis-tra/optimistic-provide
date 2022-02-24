package service

import (
	"context"
	"sync"
	"time"

	ks "github.com/whyrusleeping/go-keyspace"

	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-kad-dht/qpeerset"
	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
)

var log = logging.Logger("optprov")

type ProvideService interface {
	Provide(ctx context.Context, h *dht.Host) (*models.Provide, error)
}

var _ ProvideService = &Provide{}

type Provide struct {
	peerService PeerService
	hostService HostService
	rtService   RoutingTableService
	maService   MultiAddressService
	dialService DialService
	connService ConnectionService
	fnService   FindNodesService
	psService   PeerStateService
	apService   AddProvidersService
	provideRepo repo.ProvideRepo
}

func NewProvideService(peerService PeerService, hostService HostService, rtService RoutingTableService, maService MultiAddressService, dialService DialService, connService ConnectionService, fnService FindNodesService, psService PeerStateService, apService AddProvidersService, provideRepo repo.ProvideRepo) *Provide {
	return &Provide{
		peerService: peerService,
		hostService: hostService,
		rtService:   rtService,
		maService:   maService,
		dialService: dialService,
		connService: connService,
		fnService:   fnService,
		psService:   psService,
		apService:   apService,
		provideRepo: provideRepo,
	}
}

func (ps *Provide) Provide(ctx context.Context, h *dht.Host) (*models.Provide, error) {
	rts, err := ps.rtService.Save(ctx, h)
	if err != nil {
		return nil, err
	}

	content, err := util.NewRandomContent()
	if err != nil {
		return nil, err
	}

	provide := &models.Provide{
		ProviderID:            h.DBPeer.ID,
		ContentID:             content.CID.String(),
		Distance:              ks.XORKeySpace.Key([]byte(h.ID())).Distance(ks.XORKeySpace.Key(content.CID.Hash())).Bytes(),
		InitialRoutingTableID: rts.ID,
		StartedAt:             time.Now(),
	}

	provide, err = ps.provideRepo.Save(ctx, provide)
	if err != nil {
		return nil, err
	}

	go ps.startProviding(h, provide, content)

	return provide, nil
}

func (ps *Provide) startProviding(h *dht.Host, provide *models.Provide, content *util.Content) {
	ctx := context.Background()

	state := &ProvideState{
		h:                    h,
		content:              content,
		dialsLk:              sync.RWMutex{},
		dials:                []*DialSpan{},
		findNodesLk:          sync.RWMutex{},
		findNodes:            []*FindNodesSpan{},
		addProvidersLk:       sync.RWMutex{},
		addProviders:         []*AddProvidersSpan{},
		connectionsStartedLk: sync.RWMutex{},
		connectionsStarted:   map[peer.ID]time.Time{},
		connectionsLk:        sync.RWMutex{},
		connections:          []*ConnectionSpan{},
		relevantPeers:        sync.Map{},
		peerSet:              qpeerset.NewQueryPeerset(string(content.CID.Hash())),
	}

	ctx = state.Register(ctx)
	err := h.DHT.Provide(ctx, content.CID, true)
	end := time.Now()
	state.Unregister()

	rts, err := ps.rtService.Save(context.Background(), h)
	if err != nil {
		log.Warn(err)
	}

	provide.Error = nullStringFromErr(err)
	provide.EndedAt = null.TimeFrom(end)
	provide.FinalRoutingTableID = null.IntFrom(rts.ID)
	if provide, err = ps.provideRepo.Update(context.Background(), provide); err != nil {
		log.Warn(err)
	}

	if err = ps.dialService.Save(context.Background(), h.Host, provide.ID, state.dials); err != nil {
		log.Warn(err)
	}

	if err = ps.connService.Save(context.Background(), h.Host, provide.ID, state.connections); err != nil {
		log.Warn(err)
	}

	if err = ps.fnService.Save(context.Background(), h.Host, provide.ID, state.findNodes); err != nil {
		log.Warn(err)
	}

	if err = ps.apService.Save(context.Background(), h.Host, provide.ID, state.addProviders); err != nil {
		log.Warn(err)
	}

	if err = ps.psService.Save(context.Background(), h.Host, provide.ID, state.peerSet.AllStates()); err != nil {
		log.Warn(err)
	}
}

func nullStringFromErr(err error) null.String {
	if err == nil {
		return null.NewString("", false)
	}
	return null.StringFrom(err.Error())
}
