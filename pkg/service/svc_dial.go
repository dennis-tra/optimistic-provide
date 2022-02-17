package service

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/repo"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/volatiletech/null/v8"
)

type DialService interface {
	Save(ctx context.Context, h host.Host, provideID int, dials []*DialSpan) error
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

func (d *Dial) Save(ctx context.Context, h host.Host, provideID int, dials []*DialSpan) error {
	log.Info("Saving dials...")
	defer log.Info("Done saving dials")

	localPeer, err := d.peerService.UpsertLocalPeer(h)
	if err != nil {
		return err
	}

	for _, dial := range dials {
		remotePeer, err := d.peerService.UpsertPeer(h, dial.RemotePeerID)
		if err != nil {
			return err
		}

		maddr, err := d.maService.UpsertMultiAddress(ctx, dial.Maddr)
		if err != nil {
			return err
		}

		errStr := ""
		if dial.Error != nil {
			errStr = dial.Error.Error()
		}
		dbDial := &models.Dial{
			ProvideID:      provideID,
			LocalID:        localPeer.ID,
			RemoteID:       remotePeer.ID,
			Transport:      dial.Trpt,
			MultiAddressID: maddr.ID,
			StartedAt:      dial.Start,
			EndedAt:        dial.End,
			Error:          null.NewString(errStr, errStr != ""),
		}
		if dbDial, err = d.dialRepo.Save(ctx, dbDial); err != nil {
			return err
		}
	}

	return nil
}
