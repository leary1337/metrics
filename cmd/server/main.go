package main

import (
	"log"

	"github.com/leary1337/metrics/internal/server"
	"github.com/leary1337/metrics/internal/server/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	s := server.NewServer(cfg)
	s.Run()
}
