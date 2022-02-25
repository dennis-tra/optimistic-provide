package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/friendsofgo/errors"

	ks "github.com/whyrusleeping/go-keyspace"

	logging "github.com/ipfs/go-log"
	"github.com/volatiletech/null/v8"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
)

var log = logging.Logger("optprov")

type ProvideService interface {
	Provide(ctx context.Context, h *dht.Host) (*models.Provide, error)
	List(ctx context.Context, h *dht.Host) ([]*models.Provide, error)
	Get(ctx context.Context, h *dht.Host, id int) (*models.Provide, error)
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
	txn, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}
	defer deferTxRollback(txn)

	rts, err := ps.rtService.Save(ctx, txn, h)
	if err != nil {
		return nil, errors.Wrap(err, "save routing table")
	}

	content, err := util.NewRandomContent()
	if err != nil {
		return nil, errors.Wrap(err, "new random content")
	}

	provide := &models.Provide{
		ProviderID:            h.DBPeer.ID,
		ContentID:             content.CID.String(),
		Distance:              ks.XORKeySpace.Key([]byte(h.ID())).Distance(ks.XORKeySpace.Key(content.CID.Hash())).Bytes(),
		InitialRoutingTableID: rts.ID,
		StartedAt:             time.Now(),
	}

	if err = provide.Insert(ctx, txn, boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "inserting provide")
	}

	if err := txn.Commit(); err != nil {
		return nil, errors.Wrap(err, "committing transaction")
	}

	go ps.startProviding(h, provide, content)

	return provide, nil
}

func (ps *Provide) List(ctx context.Context, h *dht.Host) ([]*models.Provide, error) {
	return ps.provideRepo.List(ctx, h.ID().Pretty())
}

func (ps *Provide) Get(ctx context.Context, h *dht.Host, provideID int) (*models.Provide, error) {
	return ps.provideRepo.Get(ctx, h.ID().String(), provideID)
}

func (ps *Provide) startProviding(h *dht.Host, provide *models.Provide, content *util.Content) {
	ctx := context.Background()

	state := NewProvideState(h, content)
	ctx = state.Register(ctx)
	log.Infow("Start providing content", "cid", content.CID.String())
	err := h.DHT.Provide(ctx, content.CID, true)
	log.Infow("Done providing content", "cid", content.CID.String())
	end := time.Now()
	state.Unregister()

	provide.Error = null.StringFromPtr(util.ErrorStr(err))
	provide.EndedAt = null.TimeFrom(end)

	if err = ps.saveProvide(h, provide, state); err != nil {
		log.Errorw("error saving provide operation", "err", err)
	}
}

func (ps *Provide) saveProvide(h *dht.Host, provide *models.Provide, state *ProvideState) error {
	saveCtx := context.Background()

	txn, err := boil.BeginTx(saveCtx, nil)
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	defer deferTxRollback(txn)

	rts, err := ps.rtService.Save(saveCtx, txn, h)
	if err != nil {
		return errors.Wrap(err, "saving final routing table")
	}

	dbDials, err := ps.dialService.Save(saveCtx, txn, h, state.dials)
	if err != nil {
		return errors.Wrap(err, "saving dials")
	}

	dbConns, err := ps.connService.Save(saveCtx, txn, h, state.connections)
	if err != nil {
		return errors.Wrap(err, "saving connections")
	}

	findNodesRPCs, err := ps.fnService.Save(saveCtx, txn, h, state.findNodes)
	if err != nil {
		return errors.Wrap(err, "saving find nodes RPCs")
	}

	addProviderRPCs, err := ps.apService.Save(saveCtx, txn, h, state.addProviders)
	if err != nil {
		return errors.Wrap(err, "saving add provider RPCs")
	}

	peerStates, err := ps.psService.Save(saveCtx, txn, h, state.peerSet.AllStates())
	if err != nil {
		return errors.Wrap(err, "saving peer states")
	}

	if err = provide.SetFinalRoutingTable(saveCtx, txn, false, rts); err != nil {
		return errors.Wrap(err, "setting final routing table")
	}

	if err = provide.SetDials(saveCtx, txn, false, dbDials...); err != nil {
		return errors.Wrap(err, "setting dials")
	}

	if err = provide.SetConnections(saveCtx, txn, false, dbConns...); err != nil {
		return errors.Wrap(err, "setting connections")
	}

	if err = provide.SetFindNodesRPCS(saveCtx, txn, false, findNodesRPCs...); err != nil {
		return errors.Wrap(err, "setting find nodes rpcs")
	}

	if err = provide.SetAddProviderRPCS(saveCtx, txn, false, addProviderRPCs...); err != nil {
		return errors.Wrap(err, "setting add provider rpcs")
	}

	if err = provide.SetPeerStates(saveCtx, txn, false, peerStates...); err != nil {
		return errors.Wrap(err, "setting peer states")
	}

	provide.DoneAt = null.TimeFrom(time.Now())

	if _, err = provide.Update(saveCtx, txn, boil.Infer()); err != nil {
		return errors.Wrap(err, "updating provider")
	}

	if err := txn.Commit(); err != nil {
		return errors.Wrap(err, "committing transaction")
	}

	log.Infow("Done Providing")
	return nil
}

func deferTxRollback(tx *sql.Tx) {
	if err := tx.Rollback(); err != nil && !errors.Is(err, sql.ErrTxDone) {
		log.Warnw("error rolling back transaction", "err", err)
	}
}
