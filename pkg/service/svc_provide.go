package service

import (
	"context"
	"errors"
	"sync"
	"time"

	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/lib"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

var log = logging.Logger("optprov")

type ProvideService interface {
	Provide(ctx context.Context, hostID peer.ID) (*models.Provide, error)
}

var _ ProvideService = &Provide{}

type Provide struct {
	peerService PeerService
	rtService   RoutingTableService
	hostService HostService
	maService   MultiAddressService
	dialService DialService
	connService ConnectionService
	provideRepo repo.ProvideRepo
}

func NewProvideService(peerService PeerService, hostService HostService, rtService RoutingTableService, maService MultiAddressService, dialService DialService, connService ConnectionService, provideRepo repo.ProvideRepo) *Provide {
	return &Provide{
		peerService: peerService,
		rtService:   rtService,
		hostService: hostService,
		maService:   maService,
		dialService: dialService,
		connService: connService,
		provideRepo: provideRepo,
	}
}

func (ps *Provide) Provide(ctx context.Context, hostID peer.ID) (*models.Provide, error) {
	h, found := ps.hostService.Host(hostID)
	if !found {
		return nil, errors.New("host not found")
	}

	rts, err := ps.rtService.SaveRoutingTable(context.Background(), h)
	if err != nil {
		return nil, err
	}

	content, err := lib.NewRandomContent()
	if err != nil {
		return nil, err
	}

	provide := &models.Provide{
		ProviderID:            h.DBPeer.ID,
		ContentID:             content.CID.String(),
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

func (ps *Provide) startProviding(h *dht.Host, dbProvide *models.Provide, content *lib.Content) {
	ctx := context.Background()

	state := &ProvideState{
		h:                    h,
		dialsLk:              sync.RWMutex{},
		dials:                []*DialSpan{},
		connectionsStartedLk: sync.RWMutex{},
		connectionsStarted:   map[peer.ID]time.Time{},
		connectionsLk:        sync.RWMutex{},
		connections:          []*ConnectionSpan{},
		queriesStartedLk:     sync.RWMutex{},
		queriesStarted:       map[peer.ID]time.Time{},
		queriesLk:            sync.RWMutex{},
		queries:              []*QuerySpan{},
	}

	ctx, cancel := state.Register(ctx)
	err := h.DHT.Provide(ctx, content.CID, true)
	state.Unregister(cancel)

	if err != nil {
		dbProvide.Error = null.StringFrom(err.Error())
		if dbProvide, err = ps.provideRepo.Update(ctx, dbProvide); err != nil {
			log.Warn("Could not update provide", err)
		}
	}

	if err = ps.dialService.Save(context.Background(), h.Host, dbProvide.ID, state.dials); err != nil {
		log.Error(err)
	}

	if err = ps.connService.Save(context.Background(), h.Host, dbProvide.ID, state.connections); err != nil {
		log.Error(err)
	}
}
