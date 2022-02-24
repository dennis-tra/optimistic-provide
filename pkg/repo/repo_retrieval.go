package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type RetrievalRepo interface {
	Save(ctx context.Context, provide *models.Retrieval) (*models.Retrieval, error)
	Update(ctx context.Context, provide *models.Retrieval) (*models.Retrieval, error)
	List(ctx context.Context, hostID string) ([]*models.Retrieval, error)
	Get(ctx context.Context, hostID string, provideID int) (*models.Retrieval, error)
}

var _ RetrievalRepo = &Retrieval{}

type Retrieval struct {
	dbc *db.Client
}

func NewRetrievalRepo(dbc *db.Client) RetrievalRepo {
	return &Retrieval{
		dbc: dbc,
	}
}

func (p Retrieval) Save(ctx context.Context, provide *models.Retrieval) (*models.Retrieval, error) {
	return provide, provide.Insert(ctx, p.dbc, boil.Infer())
}

func (p Retrieval) Update(ctx context.Context, provide *models.Retrieval) (*models.Retrieval, error) {
	_, err := provide.Update(ctx, p.dbc, boil.Infer())
	return provide, err
}

func (p Retrieval) List(ctx context.Context, hostID string) ([]*models.Retrieval, error) {
	return models.Retrievals(
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.RetrievalColumns.RetrieverID),
		models.PeerWhere.MultiHash.EQ(hostID),
		qm.OrderBy(models.RetrievalColumns.CreatedAt),
	).All(ctx, p.dbc)
}

func (p Retrieval) Get(ctx context.Context, hostID string, provideID int) (*models.Retrieval, error) {
	return models.Retrievals(
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.RetrievalColumns.RetrieverID),
		models.PeerWhere.MultiHash.EQ(hostID),
		models.RetrievalWhere.ID.EQ(provideID),
	).One(ctx, p.dbc)
}
