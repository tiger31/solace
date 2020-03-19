package coverage

type SpecMetrics struct {
	Paths    map[string]*PathMetric	`json:"paths,omitempty"`
	poll     MetricsPoll
	Memoized float32								`json:"coverage"`
}

func (m *SpecMetrics) Coverage() float32 {
	m.Memoized = m.poll.Coverage()
	return m.Memoized
}

func (m *SpecMetrics) Reset() {
	m.poll.Reset()
}

func (m *SpecMetrics) Add(path string, metric *PathMetric) {
	m.Paths[path] = metric
	m.poll.Add(metric)
}
