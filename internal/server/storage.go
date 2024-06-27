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
		m.metrics[metric.Name] = metric
		return
	}
	// Counter - прибавляем к прошлому
	if metric.Type == CounterType {
		oldMetric, ok := m.metrics[metric.Name]
		if !ok {
			m.metrics[metric.Name] = metric
		} else {
			oldMetric.SetValue(oldMetric.Value + metric.Value)
		}
		return
	}
}
