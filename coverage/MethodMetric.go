package coverage

import "net/http"

type MethodMetric struct {
 Request *RequestMetrics   `json:"request"`
 Response *ResponseMetrics `json:"response"`
 MIME *MIMEMetric          `json:"content_types"`
 Memoized float32          `json:"coverage"`
}

func (m *MethodMetric) Coverage() float32 {
 m.Memoized = (m.Request.Coverage() + m.Response.Coverage() + m.MIME.Coverage()) / 3
 return m.Memoized
}

func (m *MethodMetric) Reset() {
 m.Request.Reset()
 m.Response.Reset()
 m.MIME.Reset()
}

func (m *MethodMetric) ProcessRequest(r *http.Request) {
 m.Request.ProcessRequest(r)
}

func (m *MethodMetric) ProcessResponse(req *http.Request, res *http.Response) {
 m.Response.ProcessResponse(req, res)
}
