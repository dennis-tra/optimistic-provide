package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"os"
	"path"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

// ProvideCommand contains the provide sub-command configuration.
var ProvideCommand = &cli.Command{
	Name:   "provide",
	Usage:  "Starts a DHT measurement experiment by providing and requesting random content.",
	Action: ProvideAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:        "requesters",
			Usage:       "How many requesting libp2p hosts should be spawned",
			EnvVars:     []string{"TENMA_PROVIDE_REQUESTER_COUNT"},
			DefaultText: "1",
			Value:       1,
		},
		&cli.StringFlag{
			Name:        "out",
			Aliases:     []string{"o"},
			Usage:       "Write measurement data to this directory",
			EnvVars:     []string{"TENMA_PROVIDE_OUT"},
			DefaultText: "out",
			Value:       "out",
		},
		&cli.BoolFlag{
			Name:    "init-rt",
			Usage:   "Whether to initialize the routing table of the provider and requesters.",
			EnvVars: []string{"TENMA_PROVIDE_INIT_ROUTING_TABLE"},
		},
		&cli.IntFlag{
			Name:        "runs",
			Usage:       "How many measurement runs should be performed",
			EnvVars:     []string{"TENMA_PROVIDE_RUN_COUNT"},
			DefaultText: "1",
			Value:       1,
		},
	},
}

// ProvideAction is the function that is called when running `tenma provide`.
func ProvideAction(c *cli.Context) error {
	log.Infoln("Starting Tenma DHT measurement...")

	// Attempt to create the results directory
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

	runStart := time.Now()
	for i := 0; i < c.Int("runs"); i++ {
		log.WithField("total", c.Int("runs")).Infof("Starting measurement run %d", i+1)

		measurement, err := StartMeasurement(c, provider, requesters)
		if errors.Is(err, context.Canceled) {
			return c.Context.Err()
		} else if err != nil {
			log.WithError(err).Warnln("error in measurement run")
			continue
		}

		filename := path.Join(c.String("out"), fmt.Sprintf("%s_measurement_%03d.json", runStart.Format("2006-01-02T15:04"), i+1))
		if err = measurement.Save(filename); err != nil {
			log.WithError(err).Warnln("error saving measurement data")
			continue
		}
		log.WithField("filename", filename).Infoln("Saved measurement data")
	}

	return nil
}

// StartMeasurement generates new random content and instructs the provider to advertise the existence to the DHT.
// After this call has resolved the requesters try to find the provider record. All steps (dialing, DHT RPC calls)
// are recorded for later analysis.
func StartMeasurement(c *cli.Context, provider *Provider, requesters map[peer.ID]*Requester) (*Measurement, error) {
	// Generate random content that we'll provide in the DHT.
	content, err := NewRandomContent()
	if err != nil {
		return nil, errors.Wrap(err, "new random content")
	}
	log.WithField("contentID", content.cid.String()[:IDLength]).Infof("Generated random content")

	// Find peers to store the provider record
	startTime := time.Now()
	pr, err := provider.Run(c.Context, content, provider.RunAction)
	if err != nil {
		return nil, errors.Wrap(err, "provide run")
	}

	// Start requesters to find the provider record
	runsChan := make(chan *Run)
	for _, r := range requesters {
		r2 := r // copy pointer
		go func() {
			run, err := r2.Run(c.Context, content, r2.RunAction)
			if err != nil {
				log.WithError(err).Warnln("error requesting content")
			}
			runsChan <- run
		}()
	}

	// Wait for Requesters to finish
	var runs []*Run
	for i := 0; i < len(requesters); i++ {
		if run := <-runsChan; run != nil {
			runs = append(runs, run)
		}
	}
	close(runsChan)

	requesterRunData := map[string]RunData{}
	for _, run := range runs {
		requesterRunData[run.LocalID.Pretty()] = run.Data(content)
	}

	return &Measurement{
		StartedAt:  startTime,
		EndedAt:    time.Now(),
		ContentID:  content.cid.String(),
		Provider:   pr.Data(content),
		Requesters: requesterRunData,
		InitRT:     false,
	}, nil
}
