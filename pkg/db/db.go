package db

import (
	"database/sql"
	"fmt"
	"sort"

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

	return &Client{dbh}, nil
}

func (c *Client) UpsertLocalPeer(pid peer.ID) (*models.Peer, error) {
	protocols := []string{}
	for _, prot := range kaddht.DefaultProtocols {
		protocols = append(protocols, string(prot))
	}
	return c.UpsertPeer(c.DB, pid, "optprov", protocols)
}

func (c *Client) UpsertPeer(exec boil.ContextExecutor, pid peer.ID, av string, protocols []string) (*models.Peer, error) {
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
