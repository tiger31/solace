package coverage

import (
	"github.com/go-openapi/spec"
	"net/url"
	"regexp"
	"strings"
)

var PathVariable, _ = regexp.Compile(`\{.*\}`)

type SwaggerRef struct {
	Item *spec.PathItem
	Metrics *PathMetric
}

type Vertex struct {
	Value *SwaggerRef
	IsPathVariable bool
	Vertexes map[string]*Vertex
}

type CovSpec struct {
	Root *Vertex
	Metrics *SpecMetrics
}

/**
	Adds path po tree
	Path splits by / and each part used as tree vertex
 */
func (v *Vertex) AddPath(path string, item spec.PathItem, metrics *SpecMetrics) {
	parts := strings.Split(path, "/")
	vertex := v
	for _, part := range parts { // For each part in path
		if _, exists := vertex.Vertexes[part]; !exists { //If this vertex not in tree
			pathVariable := PathVariable.MatchString(part) //Check if part of path looks like "{foo}" -> it's path param
			vertex.Vertexes[part] = &Vertex{
				Value: 					nil,
				Vertexes:				make(map[string]*Vertex),
				IsPathVariable:	pathVariable,
			}
		}
		vertex = vertex.Vertexes[part] //Set children vertex as next
	}
	pathMetrics := CreatePathMetric(&item)
	vertex.Value = &SwaggerRef{
		Item:    &item,
		Metrics: &pathMetrics,
	} //Set to target vertex it's PathItem
	metrics.Add(path, &pathMetrics)
}

func (v* Vertex) getPathChildren() []*Vertex {
	children := make([]*Vertex, 0, len(v.Vertexes))
	for _, value := range v.Vertexes {
		if value.IsPathVariable {
			children = append(children, value)
		}
	}
	return children
}

func (v *Vertex) ResolvePath(path string) *SwaggerRef {
	if path == "" {
		return v.Value
	}
	parts := strings.Split(path, "/")
	vertex := v
	for i, part := range parts {
		if vert, exists := vertex.Vertexes[part]; exists {
			vertex = vert
		} else {
			probabilities := vertex.getPathChildren()
			if len(probabilities) > 0 {
				for _, prob := range probabilities {
					if item := prob.ResolvePath(strings.Join(parts[i+1:], "/")); item != nil {
						return item
					}
				}
				return nil
			}
			return nil
		}
	}
	return vertex.Value
}

func (v *Vertex) ResolveURL(url *url.URL) *SwaggerRef {
	return v.ResolvePath(url.Path)
}

func FromSwagger(swagger *spec.Swagger) *CovSpec {
	vertex := Vertex{
		Value: nil,
		Vertexes: make(map[string]*Vertex),
	}
	spec := SpecMetrics{
		Paths: make(map[string]*PathMetric),
		poll: MetricsPoll{
			Poll: make([]Metric, 0, len(swagger.Paths.Paths)),
		},
	}
	metrics := CovSpec{
		Root:    &vertex,
		Metrics: &spec,
	}
	for path, item := range swagger.Paths.Paths {
		vertex.AddPath(path, item, &spec)
	}
	return &metrics
}

