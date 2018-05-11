package config

import (
	"github.com/apex/log"
	"github.com/caarlos0/env"
)

type Config struct {
	Port           string   `env:"PORT" envDefault:"8081"`
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envSeparator:"," envDefault:"*"`
}

func MustGet() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.WithError(err).Fatal("failed to load config")
	}
	return cfg
}
