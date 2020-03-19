package coverage

import "net/http"

type RequestMetrics struct {
	PresentMetric *PresentMetric
	poll MetricsPoll
}

func (m *RequestMetrics) Coverage() float32 {
	return m.poll.Coverage()
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
			[]Metric{PresentMetric},
		},
	}
}
