package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddr string `env:"ADDRESS"`
}

const DefaultServerAddr = `localhost:8080`

func NewConfig() (*Config, error) {
	var cfg Config

	// Сначала определите все флаги.
	flag.StringVar(&cfg.ServerAddr, "a", DefaultServerAddr, "server address (default localhost:8080)")
	flag.Parse()

	// Парсинг переменных окружения.
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	// Применение значений по умолчанию, если не заданы ни флаги, ни переменные окружения.
	if cfg.ServerAddr == "" {
		cfg.ServerAddr = DefaultServerAddr
	}
	return &cfg, nil
}
