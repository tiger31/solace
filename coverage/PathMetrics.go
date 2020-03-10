package coverage

import (
	"ally/coverage/metrics"
)

type Metric interface {
	Coverage() float32
	Reset()
}

// -- MetricsPoll -- //

type MetricsPoll struct {
	Poll []Metric
}

func (m *MetricsPoll) Coverage() float32 {
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

// -- RequestMetrics -- //

type RequestMetrics struct {
	PresentMetric *metrics.PresentMetric
	poll MetricsPoll
}

func (m *RequestMetrics) Coverage() float32 {
	return m.poll.Coverage()
}

func (m *RequestMetrics) Reset() {
	m.poll.Reset()
}

func CreateRequestMetric() RequestMetrics {
	PresentMetric :=& metrics.PresentMetric{}
	return RequestMetrics{
		PresentMetric: PresentMetric,
		poll:   MetricsPoll{
			[]Metric{PresentMetric},
		},
	}
}

// -- ResponseMetrics -- //

type ResponseMetrics struct {
	PresentMetric *metrics.PresentMetric
	poll MetricsPoll
}

func (m *ResponseMetrics) Coverage() float32 {
	return m.poll.Coverage()
}

func (m *ResponseMetrics) Reset() {
	m.poll.Reset()
}

func CreateResponseMetric() ResponseMetrics {
	PresentMetric := &metrics.PresentMetric{}
	return ResponseMetrics{
		PresentMetric: PresentMetric,
		poll:   MetricsPoll{
			[]Metric{PresentMetric},
		},
	}
}

// -- PathMetrics -- //

type PathMetrics struct {
	Req RequestMetrics
	Res ResponseMetrics
}

func (m *PathMetrics) Coverage() float32 {
	return (m.Req.Coverage() + m.Res.Coverage()) / 2
}

func (m *PathMetrics) Reset() {
	m.Req.Reset()
	m.Res.Reset()
}

// -- SpecMetrics -- //

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

