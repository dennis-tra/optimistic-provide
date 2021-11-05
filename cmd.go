package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p-core/peer"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	_ "net/http/pprof"
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
)

func main() {
	// ShortCommit version tag
	verTag := fmt.Sprintf("v%s+%s", RawVersion, ShortCommit)

	app := &cli.App{
		Name:      "tenma",
		Usage:     "A libp2p DHT performance measurement tool.",
		UsageText: "tenma [global options] command [command options] [arguments...]",
		Authors: []*cli.Author{
			{
				Name:  "Dennis Trautwein",
				Email: "tenma@dtrautwein.eu",
			},
		},
		Version: verTag,
		Before:  Before,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Set this flag to enable debug logging",
				EnvVars: []string{"TENMA_DEBUG"},
			},
			&cli.IntFlag{
				Name:        "log-level",
				Usage:       "Set this flag to a value from 0 (least verbose) to 6 (most verbose). Overrides the --debug flag",
				EnvVars:     []string{"TENMA_LOG_LEVEL"},
				Value:       4,
				DefaultText: "4",
			},
			&cli.StringSliceFlag{
				Name:    "protocols",
				Usage:   "Comma separated list of protocols that this crawler should look for",
				EnvVars: []string{"TENMA_PROTOCOLS"},
			},
		},
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			ProvideCommand,
		},
	}

	sigs := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	go func() {
		sig := <-sigs
		log.Printf("Received %s signal - Stopping...\n", sig.String())
		signal.Stop(sigs)
		cancel()
	}()

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Errorf("error: %v\n", err)
		os.Exit(1)
	}
}

// Before is executed before any subcommands are run, but after the context is ready
// If a non-nil error is returned, no subcommands are run.
func Before(c *cli.Context) error {
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	if c.IsSet("log-level") {
		ll := c.Int("log-level")
		log.SetLevel(log.Level(ll))
	}

	if c.IsSet("pprof-port") {
		go func() {
			pprof := fmt.Sprintf("localhost:%d", c.Int("pprof-port"))
			log.Debugln("Starting profiling endpoint at", pprof)
			if err := http.ListenAndServe(pprof, nil); err != nil {
				log.WithError(err).Warnln("Error serving pprof")
			}
		}()
	}

	return nil
}

func FmtPeerID(id peer.ID) string {
	return id.Pretty()[:IDLength]
}
