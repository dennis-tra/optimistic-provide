package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type ProvideRepo interface {
	Save(ctx context.Context, provide *models.Provide) (*models.Provide, error)
	Update(ctx context.Context, provide *models.Provide) (*models.Provide, error)
	List(ctx context.Context, hostID string) ([]*models.Provide, error)
	Get(ctx context.Context, hostID string, provideID int) (*models.Provide, error)
}

var _ ProvideRepo = &Provide{}

type Provide struct {
	dbc *db.Client
}

func NewProvideRepo(dbc *db.Client) ProvideRepo {
	return &Provide{
		dbc: dbc,
	}
}

func (p Provide) Save(ctx context.Context, provide *models.Provide) (*models.Provide, error) {
	return provide, provide.Insert(ctx, p.dbc, boil.Infer())
}

func (p Provide) Update(ctx context.Context, provide *models.Provide) (*models.Provide, error) {
	_, err := provide.Update(ctx, p.dbc, boil.Infer())
	return provide, err
}

func (p Provide) List(ctx context.Context, hostID string) ([]*models.Provide, error) {
	return models.Provides(
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.ProvideColumns.ProviderID),
		models.PeerWhere.MultiHash.EQ(hostID),
		qm.OrderBy(models.ProvideColumns.CreatedAt),
	).All(ctx, p.dbc)
}

func (p Provide) Get(ctx context.Context, hostID string, provideID int) (*models.Provide, error) {
	return models.Provides(
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.ProvideColumns.ProviderID),
		models.PeerWhere.MultiHash.EQ(hostID),
		models.ProvideWhere.ID.EQ(provideID),
		qm.Load(models.ProvideRels.Dials),
		qm.Load(models.ProvideRels.Connections),
		qm.Load(models.ProvideRels.AddProviderRPCS),
		qm.Load(models.ProvideRels.FindNodesRPCS),
		qm.Load(models.ProvideRels.PeerStates),
	).One(ctx, p.dbc)
}
