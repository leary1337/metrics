package agent

import "time"

const (
	PollInterval   = 2 * time.Second
	ReportInterval = 10 * time.Second
)

const (
	GaugeMetricType   = "gauge"
	CounterMetricType = "counter"
)
