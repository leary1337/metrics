package main

import "github.com/leary1337/metrics/internal/agent"

const ServerAddress = `http://localhost:8080`

func main() {
	a := agent.NewAgent(ServerAddress)
	a.Run()
}
