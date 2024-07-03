package main

import (
	"log"

	"github.com/leary1337/metrics/internal/agent"
	"github.com/leary1337/metrics/internal/agent/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	a := agent.NewAgent(cfg)
	a.Run()
}
