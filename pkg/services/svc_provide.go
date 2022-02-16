package services

import (
	"context"
	"sync"
	"time"

	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/db/models"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/lib"
)

var log = logging.Logger("optprov")

type ProvideService struct {
	dbc *db.Client
}

func NewProvideService(dbc *db.Client) *ProvideService {
	return &ProvideService{
		dbc: dbc,
	}
}

func (ps *ProvideService) Provide(h *dht.Host, dbProvide *models.Provide, content *lib.Content) {
	ctx := context.Background()

	state := &ProvideState{
		h:                    h,
		dialsLk:              sync.RWMutex{},
		dials:                []*DialSpan{},
		connectionsStartedLk: sync.RWMutex{},
		connectionsStarted:   map[peer.ID]time.Time{},
		connectionsLk:        sync.RWMutex{},
		connections:          []*ConnectionSpan{},
	}

	ctx, cancel := state.Register(ctx)
	err := h.DHT.Provide(ctx, content.CID, true)
	state.Unregister(cancel)

	if err != nil {
		dbProvide.Error = null.StringFrom(err.Error())
		if _, err = dbProvide.Update(ctx, ps.dbc, boil.Infer()); err != nil {
			log.Warn("Could not update provide", err)
		}
	}

	if err = ps.SaveDials(context.Background(), h.Host, dbProvide, state.dials); err != nil {
		log.Error(err)
	}
}

func (ps *ProvideService) SaveDials(ctx context.Context, h host.Host, dbProvide *models.Provide, dials []*DialSpan) error {
	log.Info("Saving dials...")
	defer log.Info("Done saving dials")

	localPeer, err := ps.dbc.UpsertLocalPeer(h.ID())
	if err != nil {
		return err
	}

	for _, dial := range dials {
		remotePeer, err := ps.dbc.UpsertPeer(ps.dbc, h, dial.RemotePeerID)
		if err != nil {
			return err
		}

		maddr, err := ps.dbc.UpsertMultiAddress(ctx, ps.dbc, dial.Maddr)
		if err != nil {
			return err
		}

		errStr := ""
		if dial.Error != nil {
			errStr = dial.Error.Error()
		}
		dbDial := &models.Dial{
			ProvideID:      dbProvide.ID,
			LocalID:        localPeer.ID,
			RemoteID:       remotePeer.ID,
			Transport:      dial.Trpt,
			MultiAddressID: maddr.ID,
			StartedAt:      dial.Start,
			EndedAt:        dial.End,
			Error:          null.NewString(errStr, errStr != ""),
		}
		if err := dbDial.Insert(ctx, ps.dbc, boil.Infer()); err != nil {
			log.Warnw("Could not insert dial", "error", err)
		}
	}

	return nil
}
