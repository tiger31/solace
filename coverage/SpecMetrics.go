package coverage

type SpecMetrics struct {
	poll MetricsPoll
}

func (m *SpecMetrics) Coverage() float32 {
	return m.poll.Coverage()
}

func (m *SpecMetrics) Reset() {
	m.poll.Reset()
}

func (m *SpecMetrics) Add(metric Metric) {
	m.poll.Add(metric)
}
