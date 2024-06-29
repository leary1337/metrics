package server

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	storage   *MemStorage
	templates *template.Template
}

func NewHandler(storage *MemStorage) (*Handler, error) {
	templates, err := template.ParseFiles(filepath.Join("templates", "all_metrics.html"))
	if err != nil {
		return nil, err
	}
	return &Handler{
		storage:   storage,
		templates: templates,
	}, nil
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	mType := chi.URLParam(r, "mType")
	mName := chi.URLParam(r, "mName")
	mValue := chi.URLParam(r, "mValue")
	if mName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	metricValue, err := strconv.ParseFloat(mValue, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	metric := NewMetric(mType, mName, metricValue)
	if !metric.IsValidType() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.storage.AddMetric(*metric)
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMetricValue(w http.ResponseWriter, r *http.Request) {
	mType := chi.URLParam(r, "mType")
	mName := chi.URLParam(r, "mName")
	if mType == "" || mName == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	m, ok := h.storage.GetByID(mType + "_" + mName)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	_, err := w.Write([]byte(strconv.FormatFloat(m.Value, 'f', -1, 64)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetAllMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := h.templates.ExecuteTemplate(w, "all_metrics.html", h.storage.AsList())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
