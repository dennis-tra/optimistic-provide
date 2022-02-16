package db

import (
	"context"
	"database/sql"
	"fmt"
	"sort"

	"github.com/dennis-tra/optimistic-provide/pkg/maxmind"
	manet "github.com/multiformats/go-multiaddr/net"

	ma "github.com/multiformats/go-multiaddr"

	"github.com/libp2p/go-libp2p-core/host"

	"github.com/dennis-tra/optimistic-provide/pkg/db/models"
	_ "github.com/lib/pq"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type Client struct {
	*sql.DB
	mmclient *maxmind.Client
}

func NewClient(host, port, dbname, user, password, sslMode string) (*Client, error) {
	// Open database handle
	srcName := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host,
		port,
		dbname,
		user,
		password,
		sslMode,
	)
	dbh, err := sql.Open("postgres", srcName)
	if err != nil {
		return nil, errors.Wrap(err, "opening database")
	}

	// Ping database to verify connection.
	if err = dbh.Ping(); err != nil {
		return nil, errors.Wrap(err, "pinging database")
	}

	mmclient, err := maxmind.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "creating maxmind client")
	}

	return &Client{DB: dbh, mmclient: mmclient}, nil
}

func (c *Client) UpsertLocalPeer(pid peer.ID) (*models.Peer, error) {
	protocols := []string{}
	for _, prot := range kaddht.DefaultProtocols {
		protocols = append(protocols, string(prot))
	}
	return c.upsertPeer(c.DB, pid, "optprov", protocols)
}

func (c *Client) UpsertPeer(exec boil.ContextExecutor, h host.Host, p peer.ID) (*models.Peer, error) {
	av := ""
	if agent, err := h.Peerstore().Get(p, "AgentVersion"); err == nil {
		av = agent.(string)
	}

	protocols := []string{}
	if prots, err := h.Peerstore().GetProtocols(p); err == nil {
		protocols = prots
	}

	return c.upsertPeer(exec, p, av, protocols)
}

func (c *Client) upsertPeer(exec boil.ContextExecutor, pid peer.ID, av string, protocols []string) (*models.Peer, error) {
	sort.Strings(protocols)

	dbPeer := &models.Peer{
		MultiHash:    pid.String(),
		AgentVersion: null.NewString(av, av != ""),
		Protocols:    types.StringArray(protocols),
	}

	rows, err := queries.Raw("SELECT upsert_peer($1, $2, $3)", dbPeer.MultiHash, dbPeer.AgentVersion.Ptr(), dbPeer.Protocols).Query(exec)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, rows.Err()
	}
	if err = rows.Scan(&dbPeer.ID); err != nil {
		return nil, err
	}
	return dbPeer, rows.Close()
}

func (c *Client) UpsertMultiAddress(ctx context.Context, exec boil.ContextExecutor, maddr ma.Multiaddr) (*models.MultiAddress, error) {
	isPublic := manet.IsPublicAddr(maddr)

	infos, err := c.mmclient.MaddrInfo(ctx, maddr)
	if err != nil {
		return nil, errors.Wrap(err, "resolve maddr infos")
	}

	countries := []string{}
	continents := []string{}
	asns := []int{}

	ipAddressIDs := []int64{}
	for address, info := range infos {
		dbIPAddress, err := c.UpsertIPAddress(ctx, exec, address, info, isPublic)
		if err != nil {
			return nil, errors.Wrap(err, "upsert ip address")
		}
		ipAddressIDs = append(ipAddressIDs, int64(dbIPAddress.ID))

		if info.Country != "" {
			countries = append(countries, info.Country)
		}

		if info.Continent != "" {
			continents = append(continents, info.Continent)
		}

		if info.ASN != 0 {
			asns = append(asns, int(info.ASN))
		}
	}

	dbMaddr := &models.MultiAddress{
		Maddr: maddr.String(),
	}

	query := queries.Raw("SELECT upsert_multi_address($1, $2, $3, $4, $5, $6)",
		maddr.String(),
		null.StringFromPtr(uniqueStr(countries)),
		null.StringFromPtr(uniqueStr(continents)),
		null.IntFromPtr(uniqueInt(asns)),
		isPublic,
		types.Int64Array(ipAddressIDs),
	)
	rows, err := query.Query(exec)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, rows.Err()
	}

	if err = rows.Scan(&dbMaddr.ID); err != nil {
		return nil, err
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}

	return dbMaddr, err
}

func (c *Client) UpsertIPAddress(ctx context.Context, exec boil.ContextExecutor, address string, info *maxmind.AddrInfo, isPublic bool) (*models.IPAddress, error) {
	dbAddress := &models.IPAddress{
		Address:   address,
		Country:   null.NewString(info.Country, info.Country != ""),
		Continent: null.NewString(info.Continent, info.Continent != ""),
		Asn:       null.NewInt(int(info.ASN), info.ASN != 0),
		IsPublic:  isPublic,
	}

	return dbAddress, dbAddress.Upsert(ctx, exec, true, []string{models.IPAddressColumns.Address}, boil.Whitelist(models.IPAddressColumns.UpdatedAt), boil.Infer())
}

func uniqueInt(input []int) *int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	if len(u) == 1 {
		return &u[0]
	}

	return nil
}

func uniqueStr(input []string) *string {
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	if len(u) == 1 {
		return &u[0]
	}

	return nil
}
