package coverage

import "net/http"

type RequestMetrics struct {
	PresentMetric *PresentMetric	`json:"present_metric"`
	poll          MetricsPoll
	Memoized      float32					`json:"coverage"`
}

func (m *RequestMetrics) Coverage() float32 {
	m.Memoized = m.poll.Coverage()
	return m.Memoized
}

func (m *RequestMetrics) Reset() {
	m.poll.Reset()
}

func (m *RequestMetrics) ProcessRequest(r *http.Request) {
	m.PresentMetric.Present = true
}

func CreateRequestMetric() RequestMetrics {
	PresentMetric :=& PresentMetric{}
	return RequestMetrics{
		PresentMetric: PresentMetric,
		poll:   MetricsPoll{
			Poll: []Metric{PresentMetric},
		},
	}
}
