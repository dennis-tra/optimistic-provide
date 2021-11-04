package main

import (
	_ "net/http/pprof"
	"os"

	"golang.org/x/sync/errgroup"

	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// ProvideCommand contains the provide sub-command configuration.
var ProvideCommand = &cli.Command{
	Name:   "provide",
	Usage:  "Starts a DHT measurement experiment by providing and requesting random content.",
	Action: ProvideAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "requesters",
			Usage:   "How many requesting libp2p hosts should be spawned",
			EnvVars: []string{"TENMA_PROVIDE_REQUESTER_COUNT"},
		},
		&cli.StringFlag{
			Name:    "out",
			Aliases: []string{"o"},
			Usage:   "Write measurement to this directory",
			EnvVars: []string{"TENMA_PROVIDE_OUT"},
		},
		&cli.BoolFlag{
			Name:    "init-rt",
			Usage:   "Whether or not Nebula should wait until the provider's routing table was refreshed",
			EnvVars: []string{"TENMA_PROVIDE_INIT_ROUTING_TABLE"},
		},
		&cli.IntFlag{
			Name:    "runs",
			Usage:   "How many provide runs should be performed",
			EnvVars: []string{"TENMA_PROVIDE_RUN_COUNT"},
		},
	},
}

// ProvideAction is the function that is called when running `nebula provide`.
func ProvideAction(c *cli.Context) error {
	log.Infoln("Starting Tenma DHT measurement...")

	if err := os.Mkdir(c.String("out"), 0o751); !os.IsExist(err) {
		return errors.Wrap(err, "creating out dir "+c.String("out"))
	}

	// Construct the provider libp2p host
	provider, err := NewProvider(c.Context)
	if err != nil {
		return errors.Wrap(err, "new provider")
	}

	// Construct the requester libp2p hosts
	requesters := map[peer.ID]*Requester{}
	for i := 0; i < c.Int("requesters"); i++ {
		r, err := NewRequester(c.Context)
		if err != nil {
			return errors.Wrap(err, "new requester")
		}
		requesters[r.h.ID()] = r
	}

	// Bootstrap both libp2p hosts by connecting to the canonical bootstrap peers.
	group, errCtx := errgroup.WithContext(c.Context)
	group.Go(func() error { return provider.Init(errCtx, c.Bool("init-rt")) })
	for _, r := range requesters {
		r2 := r // copy pointer
		group.Go(func() error { return r2.Init(errCtx, c.Bool("init-rt")) })
	}
	if err = group.Wait(); err != nil {
		return errors.Wrap(err, "bootstrap err group")
	}

	for i := 0; i < c.Int("runs"); i++ {
		log.WithField("total", c.Int("runs")).Infof("Starting measurement run %d", i+1)

		// Generate random content that we'll provide in the DHT.
		content, err := NewRandomContent()
		if err != nil {
			return errors.Wrap(err, "new random content")
		}
		log.Infof("Generated random content %s", content.cid.String())

		pr, err := provider.Run(c.Context, content)
		if err != nil {
			return errors.Wrap(err, "provide run")
		}

		runsChan := make(chan *Run)
		for _, r := range requesters {
			r2 := r // copy pointer
			go func() {
				run, err := r2.Run(c.Context, content)
				if err != nil {
					log.WithError(err).Warnln("error requesting content")
				}
				runsChan <- run
			}()
		}

		var runs []*Run
		for i := 0; i < len(requesters); i++ {
			if run := <-runsChan; run != nil {
				runs = append(runs, run)
			}
		}
		close(runsChan)

		_ = pr

	}
	return nil
}
