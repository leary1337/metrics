package agent

import (
	"math/rand"
	"runtime"
)

type Metrics struct {
	Gauge   GaugeMetrics
	Counter CounterMetrics
}

type GaugeMetrics struct {
	Alloc         float64
	BuckHashSys   float64
	Frees         float64
	GCCPUFraction float64
	GCSys         float64
	HeapAlloc     float64
	HeapIdle      float64
	HeapInuse     float64
	HeapObjects   float64
	HeapReleased  float64
	HeapSys       float64
	LastGC        float64
	Lookups       float64
	MCacheInuse   float64
	MCacheSys     float64
	MSpanInuse    float64
	MSpanSys      float64
	Mallocs       float64
	NextGC        float64
	NumForcedGC   float64
	NumGC         float64
	OtherSys      float64
	PauseTotalNs  float64
	StackInuse    float64
	StackSys      float64
	Sys           float64
	TotalAlloc    float64
	RandomValue   float64
}

type CounterMetrics struct {
	PollCount int64
}

func (m *Metrics) UpdateMetrics() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	m.Gauge.Alloc = float64(memStats.Alloc)
	m.Gauge.BuckHashSys = float64(memStats.BuckHashSys)
	m.Gauge.Frees = float64(memStats.Frees)
	m.Gauge.GCCPUFraction = memStats.GCCPUFraction
	m.Gauge.GCSys = float64(memStats.GCSys)
	m.Gauge.HeapAlloc = float64(memStats.HeapAlloc)
	m.Gauge.HeapIdle = float64(memStats.HeapIdle)
	m.Gauge.HeapInuse = float64(memStats.HeapInuse)
	m.Gauge.HeapObjects = float64(memStats.HeapObjects)
	m.Gauge.HeapReleased = float64(memStats.HeapReleased)
	m.Gauge.HeapSys = float64(memStats.HeapSys)
	m.Gauge.LastGC = float64(memStats.LastGC)
	m.Gauge.Lookups = float64(memStats.Lookups)
	m.Gauge.MCacheInuse = float64(memStats.MCacheInuse)
	m.Gauge.MCacheSys = float64(memStats.MCacheSys)
	m.Gauge.MSpanInuse = float64(memStats.MSpanInuse)
	m.Gauge.MSpanSys = float64(memStats.MSpanSys)
	m.Gauge.Mallocs = float64(memStats.Mallocs)
	m.Gauge.NextGC = float64(memStats.NextGC)
	m.Gauge.NumForcedGC = float64(memStats.NumForcedGC)
	m.Gauge.NumGC = float64(memStats.NumGC)
	m.Gauge.OtherSys = float64(memStats.OtherSys)
	m.Gauge.PauseTotalNs = float64(memStats.PauseTotalNs)
	m.Gauge.StackInuse = float64(memStats.StackInuse)
	m.Gauge.StackSys = float64(memStats.StackSys)
	m.Gauge.Sys = float64(memStats.Sys)
	m.Gauge.TotalAlloc = float64(memStats.TotalAlloc)
	m.Gauge.RandomValue = rand.Float64()

	m.Counter.PollCount++
}

func (m *Metrics) AsGaugeMap() map[string]any {
	return map[string]any{
		"Alloc":         m.Gauge.Alloc,
		"BuckHashSys":   m.Gauge.BuckHashSys,
		"Frees":         m.Gauge.Frees,
		"GCCPUFraction": m.Gauge.GCCPUFraction,
		"GCSys":         m.Gauge.GCSys,
		"HeapAlloc":     m.Gauge.HeapAlloc,
		"HeapIdle":      m.Gauge.HeapIdle,
		"HeapInuse":     m.Gauge.HeapInuse,
		"HeapObjects":   m.Gauge.HeapObjects,
		"HeapReleased":  m.Gauge.HeapReleased,
		"HeapSys":       m.Gauge.HeapSys,
		"LastGC":        m.Gauge.LastGC,
		"Lookups":       m.Gauge.Lookups,
		"MCacheInuse":   m.Gauge.MCacheInuse,
		"MCacheSys":     m.Gauge.MCacheSys,
		"MSpanInuse":    m.Gauge.MSpanInuse,
		"MSpanSys":      m.Gauge.MSpanSys,
		"Mallocs":       m.Gauge.Mallocs,
		"NextGC":        m.Gauge.NextGC,
		"NumForcedGC":   m.Gauge.NumForcedGC,
		"NumGC":         m.Gauge.NumGC,
		"OtherSys":      m.Gauge.OtherSys,
		"PauseTotalNs":  m.Gauge.PauseTotalNs,
		"StackInuse":    m.Gauge.StackInuse,
		"StackSys":      m.Gauge.StackSys,
		"Sys":           m.Gauge.Sys,
		"TotalAlloc":    m.Gauge.TotalAlloc,
		"RandomValue":   m.Gauge.RandomValue,
	}
}

func (m *Metrics) AsCounterMap() map[string]any {
	return map[string]any{
		"PollCount": m.Counter.PollCount,
	}
}
