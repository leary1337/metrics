package server

import (
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() {
	// Создаем хранилище метрик в памяти
	storage := NewMemStorage()

	serverHandler := NewServerHandler(storage)

	mux := http.NewServeMux()
	mux.HandleFunc(`/update/`, serverHandler.Update)

	err := http.ListenAndServe(s.addr, mux)
	if err != nil {
		log.Println(err)
		return
	}
}
