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
}

func (m *MetricsPoll) Coverage() float32 {
	if len(m.Poll) == 0 {
		return  1
	}
	var sum float32
	for _, metric := range m.Poll {
		sum += metric.Coverage()
	}
	return sum / float32(len(m.Poll))
}

func (m *MetricsPoll) Add(metric Metric) {
	m.Poll = append(m.Poll, metric)
}

func (m *MetricsPoll) Reset() {
	for _, metric := range m.Poll {
		metric.Reset()
	}
}
