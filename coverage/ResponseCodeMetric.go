package coverage

import (
	"github.com/go-openapi/spec"
	"net/http"
)

type ResponseCodeMetric struct {
	Present *PresentMetric
	poll MetricsPoll
}

func (m *ResponseCodeMetric) Coverage() float32 {
	return m.poll.Coverage()
}

func (m *ResponseCodeMetric) Reset() {
	m.poll.Reset()
}

func (m *ResponseCodeMetric) ProcessResponse(req *http.Request, res *http.Response) {
	m.Present.Present = true
}

func CreateResponseCodeMetric(r *spec.Response) ResponseCodeMetric {
	PresentMetric := &PresentMetric{}
	return ResponseCodeMetric{
		Present: PresentMetric,
		poll:    MetricsPoll{
			Poll: []Metric{PresentMetric},
		},
	}
}