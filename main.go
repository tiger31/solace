package main

import (
	"ally/coverage"
	"ally/proxy"
	"fmt"
	"github.com/go-openapi/spec"
	"io/ioutil"
)

func main () {
	data, err := ioutil.ReadFile("swagger.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	swagger := spec.Swagger{
		VendorExtensible: spec.VendorExtensible{},
		SwaggerProps:     spec.SwaggerProps{},
	}
	err = swagger.UnmarshalJSON(data)
	tree := coverage.FromSwagger(&swagger)
	proxy.CreateProxy(tree)
}

