package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/leary1337/metrics/internal/server/config"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run() {
	// Создаем хранилище метрик в памяти
	storage := NewMemStorage()
	h, err := NewHandler(storage)
	if err != nil {
		log.Println(err)
		return
	}

	err = http.ListenAndServe(s.cfg.ServerAddr, MetricRouter(h))
	if err != nil {
		log.Println(err)
		return
	}
}

func MetricRouter(sh *Handler) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get(`/`, sh.GetAllMetrics)
	r.Get(`/value/{mType}/{mName}`, sh.GetMetricValue)
	r.Post(`/update/{mType}/{mName}/{mValue}`, sh.Update)
	return r
}
