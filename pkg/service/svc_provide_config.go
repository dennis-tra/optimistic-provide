package service

import "github.com/dennis-tra/optimistic-provide/pkg/api/types"

type ProvideConfig struct {
	Sync bool
	Type types.ProvideType
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
