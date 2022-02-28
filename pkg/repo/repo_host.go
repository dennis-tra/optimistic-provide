package repo

import (
	"context"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type HostRepo interface {
	FindByID(ctx context.Context, hostID int) (*models.Host, error)
	FindByPeerID(ctx context.Context, peerID peer.ID) (*models.Host, error)
	FindAllUnarchived(ctx context.Context) (models.HostSlice, error)
	ArchiveHost(ctx context.Context, host *models.Host) error
}

var _ HostRepo = &Host{}

type Host struct {
	dbc *db.Client
}

func NewHostRepo(dbc *db.Client) HostRepo {
	return &Host{
		dbc: dbc,
	}
}

func (h *Host) FindByPeerID(ctx context.Context, peerID peer.ID) (*models.Host, error) {
	return models.Hosts(
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.HostColumns.PeerID),
		models.PeerWhere.MultiHash.EQ(peerID.String()),
		qm.Load(models.HostRels.Peer),
	).One(ctx, h.dbc)
}

func (h *Host) FindByID(ctx context.Context, hostID int) (*models.Host, error) {
	return models.Hosts(models.HostWhere.ID.EQ(hostID)).One(ctx, h.dbc)
}

func (h *Host) FindAllUnarchived(ctx context.Context) (models.HostSlice, error) {
	return models.Hosts(
		models.HostWhere.ArchivedAt.IsNull(),
		qm.Load(models.HostRels.Peer),
	).All(ctx, h.dbc)
}

func (h *Host) ArchiveHost(ctx context.Context, host *models.Host) error {
	host.ArchivedAt = null.TimeFrom(time.Now())
	_, err := host.Update(ctx, h.dbc, boil.Infer())
	return err
}
