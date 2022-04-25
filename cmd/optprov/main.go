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

	log "github.com/sirupsen/logrus"

	logging "github.com/ipfs/go-log"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/dennis-tra/optimistic-provide/pkg/api"
	"github.com/dennis-tra/optimistic-provide/pkg/config"
)

var (
	// RawVersion and build tag of the
	// PCP command line tool. This is
	// replaced on build via e.g.:
	// -ldflags "-X main.RawVersion=${VERSION}"
	RawVersion  = "dev"
	ShortCommit = "5f3759df" // quake
)

func main() {
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
		Version: fmt.Sprintf("v%s+%s", RawVersion, ShortCommit),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"OPTIMISTIC_PROVIDE_CONFIG_FILE"},
			},
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

// RootAction is the function that is called when running `optprov provide`.
func RootAction(c *cli.Context) error {
	log.Info("Starting DHT measurement server")

	cfg, err := config.NewConfig(c)
	if err != nil {
		return errors.Wrap(err, "new config")
	}

	if lvl, err := log.ParseLevel(cfg.App.LogLevel); err != nil {
		return errors.Wrap(err, "set log level")
	} else {
		log.SetLevel(lvl)
	}

	if err := logging.SetLogLevel("dht", cfg.App.LogLevel); err != nil {
		return errors.Wrap(err, "set DHT log level")
	}

	go func() {
		pprofAddr := cfg.PProf.Host + ":" + cfg.PProf.Port
		log.Debugf("Starting profiling endpoint at %s", pprofAddr)
		if err := http.ListenAndServe(pprofAddr, nil); err != nil {
			log.WithError(err).Error("Error serving pprof")
		}
	}()

	// Run API server
	srv, err := api.Run(c.Context, cfg)
	if err != nil {
		return errors.Wrap(err, "start api")
	}

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
