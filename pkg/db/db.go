package db

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"

	"github.com/dennis-tra/optimistic-provide/pkg/config"
)

type Client struct {
	*sql.DB
}

func NewClient(cfg *config.Config) (*Client, error) {
	// Open database handle
	srcName := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.SSLMode,
	)
	dbh, err := sql.Open("postgres", srcName)
	if err != nil {
		return nil, errors.Wrap(err, "opening database")
	}

	// Ping database to verify connection.
	if err = dbh.Ping(); err != nil {
		return nil, errors.Wrap(err, "pinging database")
	}

	return &Client{DB: dbh}, nil
}
