package coverage

import (
	"net/http"
	"strings"
)

var any = ContentType{
	Type:      "*",
	Subtype:   "*",
}

type ContentType struct {
	Type string
	Subtype string
	Parameter string
}

func ParseContentType(value string) ContentType {
	ct := ContentType{}
	split := strings.Split(value, "/")
	if len(split) < 2 {
		return any
	}
	ct.Type = split[0]
	split = strings.Split(split[1], "; ")
	ct.Subtype = split[0]
	if len(split) > 1 {
		ct.Parameter = split[1]
	}
	return ct
}

type MIMEMetric struct {
	Consumes					map[string]*ContentTypeMetric	`json:"consumes,omitempty"` //TODO If consumes not covers any, and server response code != 415 its undocumented behaviour
	ConsumesAny				bool													`json:"consumes_any"`
	AnyOther					*PresentMetric								`json:"any_other"`
	Produces					map[string]*ContentTypeMetric	`json:"produces,omitempty"` //TODO undocumented produces if not covers any
	Memoized					float32												`json:"coverage"`
	poll							MetricsPoll
}

func (m *MIMEMetric) Coverage() float32 {
	m.Memoized = m.poll.Coverage()
	return m.Memoized
}

func (m *MIMEMetric) Reset() {
	m.poll.Reset()
}

func (m *MIMEMetric) ProcessRequest(req *http.Request) {
	content := req.Header.Get("Content-Type")
	ct := ParseContentType(content)
	if ctm, exists := m.Consumes[ct.Type]; exists {
		ctm.ProcessRequestContentType(ct)
	} else {
		if any, exists := m.Consumes[any.Type]; exists {
			any.ProcessRequestContentType(ct)
		} else {
			m.AnyOther.Present = true
		}
	}
}

func (m *MIMEMetric) ProcessResponse(req *http.Request, res *http.Response) {
	content := res.Header.Get("Content-Type")
	ct := ParseContentType(content)
	if ctm, exists := m.Produces[ct.Type]; exists {
		ctm.ProcessRequestContentType(ct)
	} else {
		if any, exists := m.Produces[any.Type]; exists {
			any.ProcessRequestContentType(ct)
		} else {
			//TODO undocumented
		}
	}
}

func CreateMIMEMetric(consumes, produces []string) MIMEMetric {
	metric := MIMEMetric{
		Consumes: make(map[string]*ContentTypeMetric),
		Produces: make(map[string]*ContentTypeMetric),
		poll:     MetricsPoll{
			Poll: make([]Metric, 0),
		},
	}
	for _, mime := range consumes {
		ct := ParseContentType(mime)
		if ct.Type == any.Type && ct.Subtype == any.Subtype {
			metric.ConsumesAny = true
		}
		m, exists := metric.Consumes[ct.Type]
		if !exists {
			ctm := CreateContentTypeMetric()
			m = &ctm
			metric.Consumes[ct.Type] = m
			metric.poll.Add(m)
		}
		m.Add(ct.Subtype)
	}
	for _, mime := range produces {
		ct := ParseContentType(mime)
		m, exists := metric.Produces[ct.Type]
		if !exists {
			ctm := CreateContentTypeMetric()
			m = &ctm
			metric.Produces[ct.Type] = m
			metric.poll.Add(m)
		}
		m.Add(ct.Subtype)
	}
	if !metric.ConsumesAny {
		metric.AnyOther = &PresentMetric{}
		metric.poll.Add(metric.AnyOther)
	}
	for _, consumer := range metric.Consumes {
		consumer.CheckOther()
	}
	for _, producer := range metric.Produces {
		producer.CheckOther()
	}
	return metric
}

type ContentTypeMetric struct {
	Subtypes map[string]*PresentMetric	`json:"subtypes,omitempty"`
	SubtypesCoversAny bool							`json:"subtypes_covers_any"`
	Memoized float32										`json:"coverage"`
	poll MetricsPoll
}

func (m *ContentTypeMetric) Coverage() float32 {
	m.Memoized = m.poll.Coverage()
	return m.Memoized
}

func (m *ContentTypeMetric) Reset() {
	m.poll.Reset()
}

func (m *ContentTypeMetric) Add(subtype string) {
	present := &PresentMetric{}
	m.Subtypes[subtype] = present
	m.poll.Add(present)
}

func (m *ContentTypeMetric) CheckOther() {
	if _, exists := m.Subtypes["*"]; exists {
		m.SubtypesCoversAny = true
	}
}

func (m *ContentTypeMetric) ProcessRequestContentType(ct ContentType) {
	if ctm, exists := m.Subtypes[ct.Subtype]; exists { //If endpoint can process request content subtype
		ctm.Present = true
	} else if m.SubtypesCoversAny { //Else if endpoint can handle _/*
		m.Subtypes[any.Subtype].Present = true
	} else { //Else we got unimplemented content type and we can count it as a part of equivalence class

		//TODO undocumented
	}
}

func (m *ContentTypeMetric) ProcessResponseContentType(ct ContentType) {
	if ctm, exists := m.Subtypes[ct.Subtype]; exists { //If endpoint returned documented content subtype
		ctm.Present = true
	} else if m.SubtypesCoversAny { //Else if endpoint produce any subtype of type
		m.Subtypes[any.Subtype].Present = true
	} else { //Else we got undocumented content-type
		//TODO undocumented
	}
}

func CreateContentTypeMetric() ContentTypeMetric {
	return ContentTypeMetric{
		Subtypes:          make(map[string]*PresentMetric),
		poll:              MetricsPoll{
			Poll: make([]Metric, 0),
		},
	}
}