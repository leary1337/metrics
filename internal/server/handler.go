package server

import (
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	storage *MemStorage
}

func NewServerHandler(storage *MemStorage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	path := strings.TrimPrefix(r.URL.Path, "/update/")
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	metricValue, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	metric := NewMetric(parts[0], parts[1], metricValue)
	if !metric.IsValidType() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.storage.AddMetric(*metric)
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
