package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	Path        = "config/config.yaml"
	DefaultPath = "default.config.yaml"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Log      `yaml:"logger"`
		*Tracer  `yaml:"tracer"`
		Database `yaml:"database"`
	}

	// App -.
	App struct {
		Name string `env-required:"true" yaml:"name"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"`
	}

	// Tracer -.
	Tracer struct {
		Exporter string `yaml:"exporter"`
		Endpoint string `yaml:"endpoint"`
		Disabled bool   `yaml:"disabled"`
	}

	// Database -.
	Database struct {
		Write `env-required:"true" yaml:"write"`
	}

	Write struct {
		Connection string `env-required:"true" yaml:"connection"`
		PoolMax    int    `env-required:"true" yaml:"pool_max"`
		Database   string `env-required:"true" yaml:"database"`
		URI        string `env-required:"true" yaml:"uri"`
	}
)

// NewConfig returns permify config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./"+Path, cfg)
	if err != nil {
		err = cleanenv.ReadConfig("./"+DefaultPath, cfg)
		if err != nil {
			return nil, fmt.Errorf("config error: %w", err)
		}
	}

	return cfg, nil
}
