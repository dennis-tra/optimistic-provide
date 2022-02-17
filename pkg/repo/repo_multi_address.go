package repo

import (
	"context"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/types"

	"github.com/dennis-tra/optimistic-provide/pkg/db"
	"github.com/dennis-tra/optimistic-provide/pkg/lib"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
)

type MultiAddressRepo interface {
	UpsertMultiAddress(ctx context.Context, maddr *models.MultiAddress, addrs []*models.IPAddress, isPublic bool) (*models.MultiAddress, error)
}

var _ MultiAddressRepo = &MultiAddress{}

type MultiAddress struct {
	dbc *db.Client
}

func NewMultiAddressRepo(dbc *db.Client) MultiAddressRepo {
	return &MultiAddress{
		dbc: dbc,
	}
}

func (c *MultiAddress) UpsertMultiAddress(ctx context.Context, maddr *models.MultiAddress, ipAddresses []*models.IPAddress, isPublic bool) (*models.MultiAddress, error) {
	var countries []string
	var continents []string
	var asns []int

	ipAddressIDs := []int64{}
	for _, ipAddress := range ipAddresses {
		if ipAddress.Country.String != "" {
			countries = append(countries, ipAddress.Country.String)
		}

		if ipAddress.Continent.String != "" {
			continents = append(continents, ipAddress.Continent.String)
		}

		if ipAddress.Asn.Int != 0 {
			asns = append(asns, ipAddress.Asn.Int)
		}
	}

	query := queries.Raw("SELECT upsert_multi_address($1, $2, $3, $4, $5, $6)",
		maddr.Maddr,
		null.StringFromPtr(lib.UniqueStr(countries)),
		null.StringFromPtr(lib.UniqueStr(continents)),
		null.IntFromPtr(lib.UniqueInt(asns)),
		isPublic,
		types.Int64Array(ipAddressIDs),
	)
	rows, err := query.Query(c.dbc)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, rows.Err()
	}

	if err = rows.Scan(&maddr.ID); err != nil {
		return nil, err
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return maddr, err
}
