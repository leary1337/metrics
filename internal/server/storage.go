package server

type MemStorage struct {
	metrics map[string]Metric
}

func NewMemStorage() *MemStorage {
	return &MemStorage{metrics: make(map[string]Metric)}
}

func (m *MemStorage) AddMetric(metric Metric) {
	// Gauge - замещаем значение
	if metric.Type == GaugeType {
		m.metrics[metric.ID] = metric
		return
	}
	// Counter - прибавляем к прошлому
	if metric.Type == CounterType {
		oldMetric, ok := m.metrics[metric.ID]
		if !ok {
			m.metrics[metric.ID] = metric
		} else {
			oldMetric.SetValue(oldMetric.Value + metric.Value)
			m.metrics[metric.ID] = oldMetric
		}
		return
	}
}

func (m *MemStorage) GetByID(id string) (Metric, bool) {
	metric, ok := m.metrics[id]
	return metric, ok
}

func (m *MemStorage) AsList() []Metric {
	metrics := make([]Metric, 0, len(m.metrics))
	for _, v := range m.metrics {
		metrics = append(metrics, v)
	}
	return metrics
}
