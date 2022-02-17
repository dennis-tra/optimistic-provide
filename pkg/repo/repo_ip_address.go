package repo

import (
	"context"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/maxmind"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IPAddressRepo interface {
	UpsertIPAddress(ctx context.Context, address string, info *maxmind.AddrInfo, isPublic bool) (*models.IPAddress, error)
}

var _ IPAddressRepo = &IPAddress{}

type IPAddress struct {
	dbc *db.Client
}

func NewIPAddressRepo(dbc *db.Client) IPAddressRepo {
	return &IPAddress{
		dbc: dbc,
	}
}

func (ia *IPAddress) UpsertIPAddress(ctx context.Context, address string, info *maxmind.AddrInfo, isPublic bool) (*models.IPAddress, error) {
	dbAddress := &models.IPAddress{
		Address:   address,
		Country:   null.NewString(info.Country, info.Country != ""),
		Continent: null.NewString(info.Continent, info.Continent != ""),
		Asn:       null.NewInt(int(info.ASN), info.ASN != 0),
		IsPublic:  isPublic,
	}
	return dbAddress, dbAddress.Upsert(ctx, ia.dbc, true, []string{models.IPAddressColumns.Address}, boil.Whitelist(models.IPAddressColumns.UpdatedAt), boil.Infer())
}
