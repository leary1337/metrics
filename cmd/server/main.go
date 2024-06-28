package main

import "github.com/leary1337/metrics/internal/server"

const ServerAddress = `localhost:8080`

func main() {
	s := server.NewServer(ServerAddress)
	s.Run()
}
