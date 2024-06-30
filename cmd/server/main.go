package main

import (
	"github.com/leary1337/metrics/internal/server"
	"github.com/leary1337/metrics/internal/server/config"
)

func main() {
	s := server.NewServer(config.NewConfig())
	s.Run()
}
