package coverage

import (
	"github.com/go-openapi/spec"
	"net/http"
)

type ResponseCodeMetric struct {
	PresentMetric *PresentMetric	`json:"present_metric"`
	poll          MetricsPoll
	Memoized      float32					`json:"coverage"`
}

func (m *ResponseCodeMetric) Coverage() float32 {
	m.Memoized = m.poll.Coverage()
	return m.Memoized
}

func (m *ResponseCodeMetric) Reset() {
	m.poll.Reset()
}

func (m *ResponseCodeMetric) ProcessResponse(req *http.Request, res *http.Response) {
	m.PresentMetric.Present = true
}

func CreateResponseCodeMetric(r *spec.Response) ResponseCodeMetric {
	PresentMetric := &PresentMetric{}
	return ResponseCodeMetric{
		PresentMetric: PresentMetric,
		poll:    MetricsPoll{
			Poll: []Metric{PresentMetric},
		},
	}
}