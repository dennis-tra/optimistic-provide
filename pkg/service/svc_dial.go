package service

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/dht"

	"github.com/dennis-tra/optimistic-provide/pkg/util"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/volatiletech/null/v8"
)

type HostOperation string

const (
	HostOperationProvide   HostOperation = "PROVIDE"
	HostOperationRetrieval HostOperation = "RETRIEVAL"
)

type DialService interface {
	Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, dials []*DialSpan) (models.DialSlice, error)
}

var _ DialService = &Dial{}

type Dial struct {
	peerService PeerService
	maService   MultiAddressService
	dialRepo    repo.DialRepo
}

func NewDialService(peerService PeerService, maService MultiAddressService, dialRepo repo.DialRepo) DialService {
	return &Dial{
		peerService: peerService,
		maService:   maService,
		dialRepo:    dialRepo,
	}
}

func (d *Dial) Save(ctx context.Context, exec boil.ContextExecutor, h *dht.Host, dials []*DialSpan) (models.DialSlice, error) {
	log.Info("Saving Dials")

	dbDials := make([]*models.Dial, len(dials))
	for i, dial := range dials {
		remotePeer, err := d.peerService.UpsertPeer(ctx, exec, h, dial.RemotePeerID)
		if err != nil {
			return nil, err
		}

		maddr, err := d.maService.UpsertMultiAddress(ctx, exec, dial.Maddr)
		if err != nil {
			return nil, err
		}

		dbDial := &models.Dial{
			LocalID:        h.DBHost.PeerID,
			RemoteID:       remotePeer.ID,
			Transport:      dial.Trpt,
			MultiAddressID: maddr.ID,
			StartedAt:      dial.Start,
			EndedAt:        dial.End,
			Error:          null.StringFromPtr(util.ErrorStr(dial.Error)),
		}

		if dbDial, err = d.dialRepo.Save(ctx, exec, dbDial); err != nil {
			return nil, err
		}

		dbDials[i] = dbDial
	}

	return dbDials, nil
}
