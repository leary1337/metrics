package main

import (
	"github.com/leary1337/metrics/internal/agent"
	"github.com/leary1337/metrics/internal/agent/config"
)

func main() {
	a := agent.NewAgent(config.NewConfig())
	a.Run()
}
