package coverage

import "net/http"

type Metric interface {
	Coverage() float32
	Reset()
}

type RequestProcessor interface {
	ProcessRequest(r *http.Request)
}

type ResponseProcessor interface {
	ProcessResponse(req *http.Request, res *http.Response)
}

// -- MetricsPoll -- //

type MetricsPoll struct {
	Poll []Metric
	_coverage float32
}

func (m *MetricsPoll) Coverage() float32 {
	var sum float32
	for _, metric := range m.Poll {
		sum += metric.Coverage()
	}
	m._coverage = sum / float32(len(m.Poll))
	return m._coverage
}

func (m *MetricsPoll) Add(metric Metric) {
	m.Poll = append(m.Poll, metric)
}

func (m *MetricsPoll) Reset() {
	for _, metric := range m.Poll {
		metric.Reset()
	}
}
