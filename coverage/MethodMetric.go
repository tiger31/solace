package coverage

import "net/http"

type MethodMetric struct {
 Request *RequestMetrics
 Response *ResponseMetrics
}

func (m *MethodMetric) Coverage() float32 {
 return (m.Request.Coverage() + m.Response.Coverage()) / 2
}

func (m *MethodMetric) Reset() {
 m.Request.Reset()
 m.Response.Reset()
}

func (m *MethodMetric) ProcessRequest(r *http.Request) {
 m.Request.ProcessRequest(r)
}

func (m *MethodMetric) ProcessResponse(req *http.Request, res *http.Response) {
 m.Response.ProcessResponse(req, res)
}
