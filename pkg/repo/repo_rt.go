package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type RoutingTableRepo interface {
	Find(ctx context.Context, routingTableID int) (*models.RoutingTableSnapshot, error)
	FindAll(ctx context.Context, hostID string) ([]*models.RoutingTableSnapshot, error)
	SaveSnapshot(context.Context, int, int, int) (*models.RoutingTableSnapshot, error)
	SaveRoutingTableEntry(context.Context, *models.RoutingTableEntry) (*models.RoutingTableEntry, error)
}

var _ RoutingTableRepo = &RoutingTable{}

type RoutingTable struct {
	dbc *db.Client
}

func NewRoutingTableRepo(dbc *db.Client) RoutingTableRepo {
	return &RoutingTable{
		dbc: dbc,
	}
}

func (r *RoutingTable) SaveSnapshot(ctx context.Context, peerID int, bucketSize int, entryCount int) (*models.RoutingTableSnapshot, error) {
	rts := &models.RoutingTableSnapshot{
		PeerID:     peerID,
		BucketSize: bucketSize,
		EntryCount: entryCount,
	}
	return rts, rts.Insert(ctx, r.dbc, boil.Infer())
}

func (r *RoutingTable) SaveRoutingTableEntry(ctx context.Context, rte *models.RoutingTableEntry) (*models.RoutingTableEntry, error) {
	return rte, rte.Insert(ctx, r.dbc, boil.Infer())
}

func (r *RoutingTable) Find(ctx context.Context, routingTableID int) (*models.RoutingTableSnapshot, error) {
	return models.RoutingTableSnapshots(models.RoutingTableSnapshotWhere.ID.EQ(routingTableID)).One(ctx, r.dbc)
}

func (r *RoutingTable) FindAll(ctx context.Context, hostID string) ([]*models.RoutingTableSnapshot, error) {
	return models.RoutingTableSnapshots(
		qm.Load(models.RoutingTableSnapshotRels.Peer),
		qm.Load(models.RoutingTableSnapshotRels.RoutingTableEntries),
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.RoutingTableSnapshotColumns.PeerID),
		models.PeerWhere.MultiHash.EQ(hostID),
	).All(ctx, r.dbc)
}
