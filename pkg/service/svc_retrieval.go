package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/ipfs/go-cid"
	"github.com/volatiletech/null/v8"
	ks "github.com/whyrusleeping/go-keyspace"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
)

type RetrievalService interface {
	List(ctx context.Context, h *dht.Host) ([]*models.Retrieval, error)
	Retrieve(ctx context.Context, h *dht.Host, id cid.Cid, count int) (*models.Retrieval, error)
}

var _ RetrievalService = &Retrieval{}

type Retrieval struct {
	hostService   HostService
	rtService     RoutingTableService
	dialService   DialService
	connService   ConnectionService
	gpService     GetProvidersService
	psService     PeerStateService
	retrievalRepo repo.RetrievalRepo
}

func NewRetrievalService(hostService HostService, rtService RoutingTableService, dialService DialService, connService ConnectionService, gpService GetProvidersService, psService PeerStateService, retrievalRepo repo.RetrievalRepo) *Retrieval {
	return &Retrieval{
		hostService:   hostService,
		rtService:     rtService,
		dialService:   dialService,
		connService:   connService,
		gpService:     gpService,
		psService:     psService,
		retrievalRepo: retrievalRepo,
	}
}

func (rs *Retrieval) List(ctx context.Context, h *dht.Host) ([]*models.Retrieval, error) {
	return rs.retrievalRepo.List(ctx, h.ID().Pretty())
}

func (rs *Retrieval) Retrieve(ctx context.Context, h *dht.Host, contentID cid.Cid, count int) (*models.Retrieval, error) {
	txn, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "begin transaction")
	}
	defer deferTxRollback(txn)

	rts, err := rs.rtService.Save(ctx, txn, h)
	if err != nil {
		return nil, errors.Wrap(err, "save routing table")
	}

	retrieval := &models.Retrieval{
		RetrieverID:           h.DBPeer.ID,
		ContentID:             contentID.String(),
		Distance:              ks.XORKeySpace.Key([]byte(h.ID())).Distance(ks.XORKeySpace.Key(contentID.Hash())).Bytes(),
		InitialRoutingTableID: rts.ID,
		StartedAt:             time.Now(),
	}

	if err = retrieval.Insert(ctx, txn, boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "inserting retrieval")
	}

	if err := txn.Commit(); err != nil {
		return nil, errors.Wrap(err, "committing transaction")
	}

	go rs.startRetrieving(h, retrieval, contentID, count)

	return retrieval, nil
}

func (rs *Retrieval) startRetrieving(h *dht.Host, retrieval *models.Retrieval, contentID cid.Cid, count int) {
	ctx := context.Background()

	state := NewRetrievalState(h, contentID)

	ctx = state.Register(ctx)
	log.Infow("Start finding providers", "cid", contentID)
	for provider := range h.DHT.FindProvidersAsync(ctx, contentID, count) {
		log.Infow("Found Provider", "providerID", provider.ID.String())
	}
	log.Infow("Start finding providers", "cid", contentID)
	end := time.Now()
	state.Unregister()

	retrieval.EndedAt = null.TimeFrom(end)

	if err := rs.saveRetrieval(h, retrieval, state); err != nil {
		log.Errorw("error saving retrieval operation", "err", err)
	}
}

func (rs *Retrieval) saveRetrieval(h *dht.Host, retrieval *models.Retrieval, state *RetrievalState) error {
	saveCtx := context.Background()

	txn, err := boil.BeginTx(saveCtx, nil)
	if err != nil {
		return errors.Wrap(err, "begin transaction")
	}
	defer deferTxRollback(txn)

	rts, err := rs.rtService.Save(saveCtx, txn, h)
	if err != nil {
		return errors.Wrap(err, "saving final routing table")
	}

	dbDials, err := rs.dialService.Save(saveCtx, txn, h, state.dials)
	if err != nil {
		return errors.Wrap(err, "saving dials")
	}

	dbConns, err := rs.connService.Save(saveCtx, txn, h, state.connections)
	if err != nil {
		log.Warn(err)
	}

	allPeerState := []*models.PeerState{}
	for uuid, set := range state.peerSet {
		peerStates, err := rs.psService.Save(saveCtx, txn, h, uuid, set.AllStates())
		if err != nil {
			return errors.Wrap(err, "saving peer states")
		}
		allPeerState = append(allPeerState, peerStates...)
	}

	getProvidersRPCs, err := rs.gpService.Save(saveCtx, txn, h, state.getProviders)
	if err != nil {
		return errors.Wrap(err, "saving get providers RPCs")
	}

	if err = retrieval.SetFinalRoutingTable(saveCtx, txn, false, rts); err != nil {
		return errors.Wrap(err, "setting final routing table")
	}

	if err = retrieval.SetDials(saveCtx, txn, false, dbDials...); err != nil {
		return errors.Wrap(err, "setting dials")
	}

	if err = retrieval.SetConnections(saveCtx, txn, false, dbConns...); err != nil {
		return errors.Wrap(err, "setting connections")
	}

	if err = retrieval.SetGetProvidersRPCS(saveCtx, txn, false, getProvidersRPCs...); err != nil {
		return errors.Wrap(err, "setting get providers rpcs")
	}

	if err = retrieval.SetPeerStates(saveCtx, txn, false, allPeerState...); err != nil {
		return errors.Wrap(err, "setting peer states")
	}

	retrieval.DoneAt = null.TimeFrom(time.Now())

	if _, err = retrieval.Update(saveCtx, txn, boil.Infer()); err != nil {
		return errors.Wrap(err, "updating provider")
	}

	if err := txn.Commit(); err != nil {
		return errors.Wrap(err, "committing transaction")
	}

	log.Infow("Done Retrieving")
	return nil
}
