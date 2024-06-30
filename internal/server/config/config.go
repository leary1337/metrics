package config

import (
	"flag"
)

type Config struct {
	ServerAddr string
}

const DefaultServerAddr = `localhost:8080`

func NewConfig() *Config {
	config := &Config{}
	flag.StringVar(&config.ServerAddr, "a", DefaultServerAddr, "server address (default localhost:8080)")
	flag.Parse()
	return config
}
