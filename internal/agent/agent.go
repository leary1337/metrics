package agent

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/leary1337/metrics/internal/agent/config"
)

type Agent struct {
	cfg    *config.Config
	m      *Metrics
	client *resty.Client
}

func NewAgent(cfg *config.Config) *Agent {
	return &Agent{
		cfg:    cfg,
		m:      &Metrics{},
		client: resty.New().SetTimeout(30 * time.Second),
	}
}

func (a *Agent) Run() {
	pollTicker := time.NewTicker(a.cfg.PollInterval)
	reportTicker := time.NewTicker(a.cfg.ReportInterval)

	for {
		select {
		case <-pollTicker.C:
			a.m.UpdateMetrics()
		case <-reportTicker.C:
			err := a.sendMetrics()
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (a *Agent) sendMetrics() error {
	for name, value := range a.m.AsGaugeMap() {
		if err := a.sendMetric(name, GaugeMetricType, value); err != nil {
			return err
		}
	}
	for name, value := range a.m.AsCounterMap() {
		if err := a.sendMetric(name, CounterMetricType, value); err != nil {
			return err
		}
	}
	return nil
}

func (a *Agent) sendMetric(name, metricType string, value any) error {
	_, err := a.client.R().
		SetHeader("Content-Type", "text/plain").
		Post(fmt.Sprintf("http://%s/update/%s/%s/%v", a.cfg.ServerAddr, metricType, name, value))
	if err != nil {
		return err
	}
	return nil
}
