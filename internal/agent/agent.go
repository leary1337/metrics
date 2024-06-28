package agent

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Agent struct {
	serverAddr string
	m          *Metrics
	client     *http.Client
}

func NewAgent(serverAddr string) *Agent {
	return &Agent{
		serverAddr: serverAddr,
		m:          &Metrics{},
		client:     &http.Client{Timeout: 30 * time.Second},
	}
}

func (a *Agent) Run() {
	pollTicker := time.NewTicker(PollInterval)
	reportTicker := time.NewTicker(ReportInterval)

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
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/update/%s/%s/%v", a.serverAddr, metricType, name, value),
		nil,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "text/plain")
	_, err = a.client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
