package config

import (
	"flag"

	"github.com/caarlos0/env"
)

type Config struct {
	ServerAddr        string `env:"ADDRESS"`
	ReportIntervalSec int    `env:"REPORT_INTERVAL"`
	PollIntervalSec   int    `env:"POLL_INTERVAL"`
}

const (
	DefaultServerAddr        = `localhost:8080`
	DefaultReportIntervalSec = 10
	DefaultPollIntervalSec   = 2
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	// Сначала определите все флаги.
	flag.StringVar(&cfg.ServerAddr, "a", "", "server address (default localhost:8080)")
	flag.IntVar(&cfg.ReportIntervalSec, "r", 0, "frequency of sending metrics to the server in sec (default 10 seconds)")
	flag.IntVar(&cfg.PollIntervalSec, "p", 0, "frequency of polling metrics from the runtime package in sec (default 2 seconds).")
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
	if cfg.ReportIntervalSec == 0 {
		cfg.ReportIntervalSec = DefaultReportIntervalSec
	}
	if cfg.PollIntervalSec == 0 {
		cfg.PollIntervalSec = DefaultPollIntervalSec
	}
	return cfg, nil
}
