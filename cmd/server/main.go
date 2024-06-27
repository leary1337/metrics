package main

import (
	"log"
	"net/http"

	"github.com/leary1337/metrics/internal/server"
)

const ServerAddress = `localhost:8080`

func main() {
	// Создаем хранилище метрик в памяти
	storage := server.NewMemStorage()

	serverHandler := server.NewServerHandler(storage)

	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, serverHandler.Update)

	err := http.ListenAndServe(ServerAddress, mux)
	if err != nil {
		log.Fatal(err)
	}
}
