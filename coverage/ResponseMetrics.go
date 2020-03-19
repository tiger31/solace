package coverage

import (
	"github.com/go-openapi/spec"
	"net/http"
)

type ResponseMetrics struct {
	CodesMetrics map[int]*ResponseCodeMetric	`json:"codes"`
	poll         MetricsPoll
	Memoized     float32											`json:"coverage"`
}

func (m *ResponseMetrics) Coverage() float32 {
	m.Memoized = m.poll.Coverage()
	return m.Memoized
}

func (m *ResponseMetrics) Reset() {
	m.poll.Reset()
}

func (m *ResponseMetrics) ProcessResponse(req *http.Request, res *http.Response) {
	if metric, exists := m.CodesMetrics[res.StatusCode]; exists {
		metric.ProcessResponse(req, res)
	}
	//TODO undocumented
}

func CreateResponseMetric(r *spec.Responses) ResponseMetrics {
	codes := make(map[int]*ResponseCodeMetric)
	poll := MetricsPoll{
		Poll: make([]Metric, 0),
	}
	for code, response := range r.ResponsesProps.StatusCodeResponses {
		metric := CreateResponseCodeMetric(&response)
		codes[code] = &metric
		poll.Add(&metric)
	}
	return ResponseMetrics{
		CodesMetrics: codes,
		poll:         poll,
	}
}
