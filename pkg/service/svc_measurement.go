package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/dennis-tra/optimistic-provide/pkg/api/types"
	"github.com/dennis-tra/optimistic-provide/pkg/dht"
	"github.com/dennis-tra/optimistic-provide/pkg/models"
	"github.com/dennis-tra/optimistic-provide/pkg/util"
)

type MeasurementService interface {
	Start(ctx context.Context, h *dht.Host, t types.ProvideType, iterations int) (*models.ProvideMeasurement, error)
	Stop(measurementID int) error
}

var _ MeasurementService = &Measurement{}

type Measurement struct {
	ps ProvideService

	provideMeasurementsLk sync.RWMutex
	provideMeasurements   map[int]context.CancelFunc
}

func NewProvideMeasurementService(ps ProvideService) *Measurement {
	return &Measurement{
		ps:                  ps,
		provideMeasurements: map[int]context.CancelFunc{},
	}
}

func (m *Measurement) Start(ctx context.Context, h *dht.Host, t types.ProvideType, iterations int) (*models.ProvideMeasurement, error) {
	measurement := &models.ProvideMeasurement{
		HostID:      h.DBHost.ID,
		StartedAt:   time.Now(),
		ProvideType: string(t),
		Iterations:  iterations,
	}
	if err := measurement.Insert(ctx, boil.GetContextDB(), boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "insert provide measurement")
	}

	provideCtx, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < iterations; i++ {
			log.Infow("Providing Content", "iteration", i, "hostID", util.FmtPeerID(h.ID()), "total", iterations)
			if _, err := m.ps.ProvideSync(provideCtx, h, types.ProvideTypeSINGLEQUERY); err != nil {
				log.Warnw("Error during provide operation", "err", err)
			}
		}

		m.provideMeasurementsLk.Lock()
		delete(m.provideMeasurements, measurement.ID)
		m.provideMeasurementsLk.Unlock()

		measurement.EndedAt = null.TimeFrom(time.Now())
		if _, err := measurement.Update(context.Background(), boil.GetContextDB(), boil.Infer()); err != nil {
			log.Warnw("Could not update measurement", "err", err.Error())
		}

		log.Infow("Measurement finished")
	}()

	m.provideMeasurementsLk.Lock()
	defer m.provideMeasurementsLk.Unlock()

	m.provideMeasurements[measurement.ID] = cancel

	return measurement, nil
}

func (m *Measurement) Stop(measurementID int) error {
	m.provideMeasurementsLk.Lock()
	defer m.provideMeasurementsLk.Unlock()

	measurementCtxCancel, found := m.provideMeasurements[measurementID]
	if !found {
		return fmt.Errorf("provide measurement is not running")
	}
	measurementCtxCancel()
	return nil
}
