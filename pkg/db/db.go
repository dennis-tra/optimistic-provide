package db

import (
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

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

	driver, err := postgres.WithInstance(dbh, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "optprov", driver)
	if err != nil {
		return nil, err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	boil.SetDB(dbh)

	return &Client{DB: dbh}, nil
}
