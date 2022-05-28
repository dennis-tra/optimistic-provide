package service

import (
	"context"
	"encoding/json"
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
	StartProvide(ctx context.Context, h *dht.Host, config types.ProvideMeasurementConfiguration) (*models.Measurement, error)
	Stop(measurementID int) error
}

var _ MeasurementService = &Measurement{}

type Measurement struct {
	ps ProvideService

	measurementsLk sync.RWMutex
	measurements   map[int]context.CancelFunc
}

func NewProvideMeasurementService(ps ProvideService) *Measurement {
	return &Measurement{
		ps:           ps,
		measurements: map[int]context.CancelFunc{},
	}
}

func (m *Measurement) StartProvide(ctx context.Context, h *dht.Host, config types.ProvideMeasurementConfiguration) (*models.Measurement, error) {
	configData, err := json.Marshal(config)
	if err != nil {
		return nil, errors.Wrap(err, "marshal provide measurement config")
	}

	dbMeasure := &models.Measurement{
		HostID:        h.DBHost.ID,
		StartedAt:     time.Now(),
		Configuration: configData,
	}
	if err = dbMeasure.Insert(ctx, boil.GetContextDB(), boil.Infer()); err != nil {
		return nil, errors.Wrap(err, "insert provide measurement")
	}

	multiQueryConcurrency := 2
	if config.Concurrency != nil {
		multiQueryConcurrency = *config.Concurrency
	}

	provideCtx, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < config.Iterations; i++ {
			log.Infow("Providing Content", "iteration", i, "hostID", util.FmtPeerID(h.ID()), "total", config.Iterations)
			_, err := m.ps.Provide(provideCtx, h, ProvideSync(), ProvideType(config.ProvideType), ProvideMeasurementID(dbMeasure.ID), ProvideMultiQueryConcurrency(multiQueryConcurrency))
			if errors.Is(err, context.Canceled) {
				break
			} else if err != nil {
				log.Warnw("Error during provide operation", "err", err)
			}
		}

		m.measurementsLk.Lock()
		delete(m.measurements, dbMeasure.ID)
		m.measurementsLk.Unlock()

		dbMeasure.EndedAt = null.TimeFrom(time.Now())
		if _, err := dbMeasure.Update(context.Background(), boil.GetContextDB(), boil.Infer()); err != nil {
			log.Warnw("Could not update measurement", "err", err.Error())
		}

		log.Infow("Measurement finished")
	}()

	m.measurementsLk.Lock()
	m.measurements[dbMeasure.ID] = cancel
	m.measurementsLk.Unlock()

	return dbMeasure, nil
}

func (m *Measurement) Stop(measurementID int) error {
	m.measurementsLk.Lock()
	defer m.measurementsLk.Unlock()

	measurementCtxCancel, found := m.measurements[measurementID]
	if !found {
		return fmt.Errorf("provide measurement is not running")
	}
	measurementCtxCancel()
	return nil
}
