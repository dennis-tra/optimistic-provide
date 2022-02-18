package config

import (
	"io/ioutil"
	"os"

	"github.com/google/uuid"

	"gopkg.in/yaml.v3"

	logging "github.com/ipfs/go-log"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("optprov")

type (
	// Config -.
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		PProf `yaml:"pprof"`
		DB    `yaml:"db"`
	}

	// App -.
	App struct {
		Instance string `yaml:"-"`
		Version  string `yaml:"version"`
		LogLevel string `yaml:"log_level"`
	}

	// HTTP -.
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	// PProf -.
	PProf struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	// DB -.
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"ssl_mode"`
	}
)

// NewConfig returns app config.
func NewConfig(c *cli.Context) (*Config, error) {
	cfg := &Config{}

	if c.IsSet("config") {
		log.Debug("Loading config file from " + c.String("config"))

		f, err := os.Open(c.String("config"))
		if err != nil {
			return nil, errors.Wrap(err, "open config file")
		}

		dat, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, errors.Wrap(err, "read config file")
		}

		if err = yaml.Unmarshal(dat, cfg); err != nil {
			if err != nil {
				return nil, errors.Wrap(err, "unmarshal")
			}
		}
	}

	cfg.Instance = uuid.New().String()
	cfg.App.Version = c.App.Version

	if c.IsSet("log-level") || cfg.App.LogLevel == "" {
		cfg.App.LogLevel = c.String("log-level")
	}

	if c.IsSet("host") || cfg.HTTP.Host == "" {
		cfg.HTTP.Host = c.String("host")
	}

	if c.IsSet("port") || cfg.HTTP.Port == "" {
		cfg.HTTP.Port = c.String("port")
	}

	if c.IsSet("pprof-host") || cfg.PProf.Host == "" {
		cfg.PProf.Host = c.String("pprof-host")
	}

	if c.IsSet("pprof-port") || cfg.PProf.Port == "" {
		cfg.PProf.Port = c.String("pprof-port")
	}

	if c.IsSet("db-host") || cfg.DB.Host == "" {
		cfg.DB.Host = c.String("db-host")
	}

	if c.IsSet("db-port") || cfg.DB.Port == "" {
		cfg.DB.Port = c.String("db-port")
	}

	if c.IsSet("db-name") || cfg.DB.Name == "" {
		cfg.DB.Name = c.String("db-name")
	}

	if c.IsSet("db-user") || cfg.DB.User == "" {
		cfg.DB.User = c.String("db-user")
	}

	if c.IsSet("db-password") || cfg.DB.Password == "" {
		cfg.DB.Password = c.String("db-password")
	}

	if c.IsSet("db-sslmode") || cfg.DB.SSLMode == "" {
		cfg.DB.SSLMode = c.String("db-sslmode")
	}

	return cfg, nil
}
