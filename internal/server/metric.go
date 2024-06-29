package server

type Metric struct {
	ID    string
	Type  MetricType
	Name  string
	Value float64
}

func NewMetric(mType, name string, value float64) *Metric {
	return &Metric{
		ID:    mType + "_" + name,
		Type:  MetricType(mType),
		Name:  name,
		Value: value,
	}
}

func (m *Metric) SetValue(v float64) {
	m.Value = v
}

func (m *Metric) IsValidType() bool {
	switch m.Type {
	case GaugeType, CounterType:
		return true
	default:
		return false
	}
}
