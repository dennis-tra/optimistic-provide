package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dennis-tra/optimistic-provide/pkg/db"

	"github.com/dennis-tra/optimistic-provide/pkg/server"
	logging "github.com/ipfs/go-log"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var (
	// RawVersion and build tag of the
	// PCP command line tool. This is
	// replaced on build via e.g.:
	// -ldflags "-X main.RawVersion=${VERSION}"
	RawVersion  = "dev"
	ShortCommit = "5f3759df" // quake

	// IDLength is here as a variable so that it can be decreased for tests with mocknet where IDs are way shorter.
	// The call to FmtPeerID would panic if this value stayed at 16.
	IDLength = 16
	log      = logging.Logger("optprov")
)

func main() {
	// ShortCommit version tag
	verTag := fmt.Sprintf("v%s+%s", RawVersion, ShortCommit)

	app := &cli.App{
		Name:      "optprov",
		Usage:     "A libp2p DHT performance measurement tool.",
		UsageText: "optprov [global options] command [command options] [arguments...]",
		Authors: []*cli.Author{
			{
				Name:  "Dennis Trautwein",
				Email: "optimistic-provide@dtrautwein.eu",
			},
		},
		Version: verTag,
		Before:  Before,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				Usage:       "Set this flag to a value from 0 (least verbose) to 6 (most verbose). Overrides the --debug flag",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_LOG_LEVEL"},
				Value:       "info",
				DefaultText: "info",
			},
			&cli.StringFlag{
				Name:        "host",
				Usage:       "To which network address should the API server bind",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_HOST"},
				Value:       "0.0.0.0",
				DefaultText: "0.0.0.0",
			},
			&cli.StringFlag{
				Name:        "port",
				Usage:       "On which port should the API server listen",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_PORT"},
				Value:       "7000",
				DefaultText: "7000",
			},
			&cli.StringFlag{
				Name:        "pprof-host",
				Usage:       "To which network address should the pprof server bind",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_PPROF_HOST"},
				Value:       "localhost",
				DefaultText: "localhost",
			},
			&cli.StringFlag{
				Name:        "pprof-port",
				Usage:       "On which port should the pprof server listen",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_PPROF_PORT"},
				Value:       "6666",
				DefaultText: "6666",
			},

			&cli.StringFlag{
				Name:        "db-host",
				Usage:       "On which host address can the database be reached",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_DATABASE_HOST"},
				DefaultText: "localhost",
				Value:       "localhost",
			},
			&cli.StringFlag{
				Name:        "db-port",
				Usage:       "On which port can the database be reached",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_DATABASE_PORT"},
				DefaultText: "5432",
				Value:       "5432",
			},
			&cli.StringFlag{
				Name:        "db-name",
				Usage:       "The name of the database to use",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_DATABASE_NAME"},
				DefaultText: "optprov",
				Value:       "optprov",
			},
			&cli.StringFlag{
				Name:        "db-password",
				Usage:       "The password for the database to use",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_DATABASE_PASSWORD"},
				DefaultText: "password",
				Value:       "password",
			},
			&cli.StringFlag{
				Name:        "db-user",
				Usage:       "The user with which to access the database to use",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_DATABASE_USER"},
				DefaultText: "optprov",
				Value:       "optprov",
			},
			&cli.StringFlag{
				Name:        "db-sslmode",
				Usage:       "The sslmode to use when connecting the the database",
				EnvVars:     []string{"OPTIMISTIC_PROVIDE_DATABASE_SSL_MODE"},
				DefaultText: "disable",
				Value:       "disable",
			},
		},
		Action:               RootAction,
		EnableBashCompletion: true,
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Errorf("error: %v\n", err)
		os.Exit(1)
	}
}

// Before is executed before any subcommands are run, but after the context is ready
// If a non-nil error is returned, no subcommands are run.
func Before(c *cli.Context) error {
	logLevel := "info"
	if c.IsSet("log-level") {
		logLevel = c.String("log-level")
	}

	if err := logging.SetLogLevel("optprov", logLevel); err != nil {
		return errors.Wrap(err, "set optprov log level")
	}
	if err := logging.SetLogLevel("dht", logLevel); err != nil {
		return errors.Wrap(err, "set DHT log level")
	}

	go func() {
		pprofAddr := c.String("pprof-host") + ":" + c.String("pprof-port")
		log.Debugw("Starting profiling endpoint at", pprofAddr)
		if err := http.ListenAndServe(pprofAddr, nil); err != nil {
			log.Errorw("Error serving pprof", err)
		}
	}()
	return nil
}

// RootAction is the function that is called when running `optprov provide`.
func RootAction(c *cli.Context) error {
	log.Info("Starting DHT measurement server")

	dbc, err := db.NewClient(c.String("db-host"), c.String("db-port"), c.String("db-name"), c.String("db-user"), c.String("db-password"), c.String("db-sslmode"))
	if err != nil {
		return errors.Wrap(err, "new db client")
	}

	// Start API server
	srv := server.Start(c.Context, c.String("host"), c.String("port"), dbc)

	// Listen for the interrupt signal.
	<-c.Context.Done()

	log.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Info("Server exiting")
	return nil
}
