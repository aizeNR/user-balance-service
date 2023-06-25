package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App            App
		Postgres       Postgres
		HTTP           HTTP
	}

	Postgres struct {
		URL         string `env:"POSTGRES_URL" env-required:"true"`
	}

	HTTP struct {
		Port int `env:"HTTP_PORT" env-required:"true"`
	}

	App struct {
		Env string `env:"APP_ENV" env-required:"true"`
	}
)

// New returns app config.
func New() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

func (c *Config) IsProduction() bool {
	return c.App.Env == "production"
}

func (h HTTP) GetPort() string {
	return preparePort(h.Port)
}

func preparePort(port int) string {
	return fmt.Sprintf(":%d", port)
}
