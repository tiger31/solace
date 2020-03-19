package coverage

import (
	"github.com/go-openapi/spec"
	"net/http"
	"reflect"
	"strings"
)

var methods = []string{"Get", "Put", "Post", "Delete", "Options", "Head", "Patch"}

type PathMetric struct {
	Methods map[string]*MethodMetric
	poll MetricsPoll
}

func (m *PathMetric) Coverage() float32 {
	return m.poll.Coverage()
}

func (m *PathMetric) Reset() {
	m.poll.Reset()
}

func (m *PathMetric) ProcessRequest(r *http.Request) {
	if metric, exists := m.Methods[r.Method]; exists {
		metric.ProcessRequest(r)
	}
}

func (m *PathMetric) ProcessResponse(req *http.Request, res *http.Response) {
	if metric, exists := m.Methods[req.Method]; exists {
		metric.ProcessResponse(req, res)
	}
	//TODO undocumented
}

func CreatePathMetric(item *spec.PathItem) PathMetric {
	path := PathMetric{
		Methods: make(map[string]*MethodMetric),
		poll:    MetricsPoll{},
	}
	ref := reflect.ValueOf(item.PathItemProps)
	for _, method := range methods {
		operation := ref.FieldByName(method).Interface().(*spec.Operation)
		if operation != nil {
			req := CreateRequestMetric()
			res := CreateResponseMetric(operation.Responses)
			metric := MethodMetric{
				Request:  &req,
				Response: &res,
			}
			path.Methods[strings.ToUpper(method)] = &metric
			path.poll.Add(&metric)
		}
	}
	return path
}


