package config

import (
	"flag"
	"time"
)

type Config struct {
	ServerAddr     string
	ReportInterval time.Duration
	PollInterval   time.Duration
}

const (
	DefaultServerAddr        = `localhost:8080`
	DefaultReportIntervalSec = 10
	DefaultPollIntervalSec   = 2
)

func NewConfig() *Config {
	config := &Config{}
	flag.StringVar(&config.ServerAddr, "a", DefaultServerAddr, "server address (default localhost:8080)")
	reportSec := flag.Int(
		"r",
		DefaultReportIntervalSec,
		"frequency of sending metrics to the server in sec (default 10 seconds)",
	)
	pollSec := flag.Int(
		"p",
		DefaultPollIntervalSec,
		"frequency of polling metrics from the runtime package in sec (default 2 seconds).",
	)
	flag.Parse()

	config.ReportInterval = time.Duration(*reportSec) * time.Second
	config.PollInterval = time.Duration(*pollSec) * time.Second
	return config
}
