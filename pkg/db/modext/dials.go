package modext

import (
	"context"
	"database/sql"

	"github.com/dennis-tra/optimistic-provide/pkg/db/models"
)

type dials struct{}

var Dials = dials{}

func (dials) FirstTimeSetup(ctx context.Context, db *sql.DB, p *models.Peer) error {
	return nil
}
