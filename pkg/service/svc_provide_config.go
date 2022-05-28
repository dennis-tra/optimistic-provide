package service

import "github.com/dennis-tra/optimistic-provide/pkg/api/types"

type ProvideConfig struct {
	Sync                  bool
	Type                  types.ProvideType
	MeasurementID         int
	MultiQueryConcurrency int
}

// Apply applies the given options to the config, returning the first error
// encountered (if any).
func (cfg *ProvideConfig) Apply(opts ...ProvideOption) error {
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if err := opt(cfg); err != nil {
			return err
		}
	}
	return nil
}

type ProvideOption func(cfg *ProvideConfig) error

func ProvideSync() ProvideOption {
	return func(cfg *ProvideConfig) error {
		cfg.Sync = true
		return nil
	}
}

func ProvideType(t types.ProvideType) ProvideOption {
	return func(cfg *ProvideConfig) error {
		cfg.Type = t
		return nil
	}
}

func ProvideMeasurementID(id int) ProvideOption {
	return func(cfg *ProvideConfig) error {
		cfg.MeasurementID = id
		return nil
	}
}

func ProvideMultiQueryConcurrency(concurrency int) ProvideOption {
	return func(cfg *ProvideConfig) error {
		cfg.MultiQueryConcurrency = concurrency
		return nil
	}
}
