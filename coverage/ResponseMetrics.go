package coverage

import (
	"github.com/go-openapi/spec"
	"net/http"
)

type ResponseMetrics struct {
	CodeMetrics map[int]*ResponseCodeMetric
	poll MetricsPoll
}

func (m *ResponseMetrics) Coverage() float32 {
	return m.poll.Coverage()
}

func (m *ResponseMetrics) Reset() {
	m.poll.Reset()
}

func (m *ResponseMetrics) ProcessResponse(req *http.Request, res *http.Response) {
	if metric, exists := m.CodeMetrics[res.StatusCode]; exists {
		metric.ProcessResponse(req, res)
	}
	//TODO undocumented
}

func CreateResponseMetric(r *spec.Responses) ResponseMetrics {
	codes := make(map[int]*ResponseCodeMetric)
	poll := MetricsPoll{
		make([]Metric, 0),
	}
	for code, response := range r.ResponsesProps.StatusCodeResponses {
		metric := CreateResponseCodeMetric(&response)
		codes[code] = &metric
		poll.Add(&metric)
	}
	return ResponseMetrics{
		CodeMetrics: codes,
		poll: poll,
	}
}
