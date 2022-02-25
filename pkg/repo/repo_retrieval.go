package repo

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type RetrievalRepo interface {
	Update(ctx context.Context, retrieval *models.Retrieval) (*models.Retrieval, error)
	List(ctx context.Context, hostID string) ([]*models.Retrieval, error)
	Get(ctx context.Context, hostID string, retrievalID int) (*models.Retrieval, error)
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

func (r Retrieval) Update(ctx context.Context, retrieval *models.Retrieval) (*models.Retrieval, error) {
	_, err := retrieval.Update(ctx, r.dbc, boil.Infer())
	return retrieval, err
}

func (r Retrieval) List(ctx context.Context, hostID string) ([]*models.Retrieval, error) {
	return models.Retrievals(
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.RetrievalColumns.RetrieverID),
		models.PeerWhere.MultiHash.EQ(hostID),
		qm.OrderBy(models.RetrievalColumns.CreatedAt),
	).All(ctx, r.dbc)
}

func (r Retrieval) Get(ctx context.Context, hostID string, retrievalID int) (*models.Retrieval, error) {
	return models.Retrievals(
		qm.InnerJoin(models.TableNames.Peers+" ON "+models.TableNames.Peers+"."+models.PeerColumns.ID+" = "+models.RetrievalColumns.RetrieverID),
		models.PeerWhere.MultiHash.EQ(hostID),
		models.RetrievalWhere.ID.EQ(retrievalID),
	).One(ctx, r.dbc)
}
