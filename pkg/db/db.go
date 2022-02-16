package db

import (
	"context"
	"database/sql"
	"fmt"
	"sort"

	"github.com/dennis-tra/optimistic-provide/pkg/maxmind"

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
	infos, err := c.mmclient.MaddrInfo(ctx, maddr)
	if err != nil {
		return nil, errors.Wrap(err, "resolve maddr infos")
	}

	ipAddressIDs := []int64{}
	for address, info := range infos {
		dbIPAddress, err := c.UpsertIPAddress(exec, address, info)
		if err != nil {
			return nil, errors.Wrap(err, "upsert ip address")
		}
		ipAddressIDs = append(ipAddressIDs, int64(dbIPAddress.ID))
	}

	dbMaddr := &models.MultiAddress{
		Maddr: maddr.String(),
	}

	rows, err := queries.Raw("SELECT upsert_multi_address($1, $2)", maddr.String(), types.Int64Array(ipAddressIDs)).Query(exec)
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

func (c *Client) UpsertIPAddress(exec boil.ContextExecutor, address string, info *maxmind.AddrInfo) (*models.IPAddress, error) {
	dbAddress := &models.IPAddress{
		Address:   address,
		Country:   null.NewString(info.Country, info.Country != ""),
		Continent: null.NewString(info.Continent, info.Continent != ""),
		Asn:       null.NewInt(int(info.ASN), info.ASN != 0),
	}

	rows, err := queries.Raw("SELECT upsert_ip_address($1, $2, $3, $4)", dbAddress.Address, dbAddress.Country, dbAddress.Continent, dbAddress.Asn).Query(exec)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, rows.Err()
	}
	if err = rows.Scan(&dbAddress.ID); err != nil {
		return nil, err
	}
	return dbAddress, rows.Close()
}
