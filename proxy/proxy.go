package proxy

import (
	"ally/coverage"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-httpproxy/httpproxy"
)

const (
	coveragePath = "/coverage"
	coverageHeaderSetup = "x-coverage-setup"
	coverageHeaderTeardown = "x-coverage-teardown"
	coverageHeaderValue = "1"
)

var requestsMap = make(map[*http.Request]*coverage.SwaggerRef)
var spec *coverage.CovSpec
var running = false

func coverageSetup(req *http.Request) bool {
	return req.URL.Path == coveragePath && req.Header.Get(coverageHeaderSetup) == coverageHeaderValue
}

func coverageTeardown(req *http.Request) bool {
	return req.URL.Path == coveragePath && req.Header.Get(coverageHeaderTeardown) == coverageHeaderValue
}

func OnError(ctx *httpproxy.Context, where string, err *httpproxy.Error, opErr error) {
	log.Printf("ERR: %s: %s [%s]", where, err, opErr)
}

func OnAccept(ctx *httpproxy.Context, w http.ResponseWriter, r *http.Request) bool {
	if r.Method == "GET" && !r.URL.IsAbs() && r.URL.Path == "/info" {
		w.Write([]byte("This is go-httpproxy."))
		return true
	}
	return false
}

func OnAuth(ctx *httpproxy.Context, authType string, user string, pass string) bool {
	if user == "test" && pass == "1234" {
		return true
	}
	return false
}

func OnConnect(ctx *httpproxy.Context, host string) (ConnectAction httpproxy.ConnectAction, newHost string) {
	// Apply "Man in the Middle" to all ssl connections. Never change host.
	return httpproxy.ConnectMitm, host
}

func OnRequest(ctx *httpproxy.Context, req *http.Request) (resp *http.Response) {
	if coverageSetup(req) {
		if !running {
			resp := httpproxy.InMemoryResponse(200, nil, nil)
			resp.Header.Set("Content-Type", "application/json")
			running = true
			return resp
		} else {
			body, _ := coverage.SetupError{Err: "must teardown current coverage before setup new one"}.MarshalJSON()
			resp := httpproxy.InMemoryResponse(403, nil, body)
			resp.Header.Set("Content-Type", "application/json")
			return resp
		}
	} else if coverageTeardown(req) {
		body, _ := json.Marshal(spec.Metrics.Coverage())
		spec.Metrics.Reset()
		running = false
		resp := httpproxy.InMemoryResponse(200, nil, body)
		return resp
	} else {
		if running {
			ref := spec.Root.ResolveURL(req.URL)
			ref.Metrics.ProcessRequest(req)
			requestsMap[req] = ref
		}
		return
	}
}

func OnResponse(ctx *httpproxy.Context, req *http.Request, resp *http.Response) {
	if running {
		ref := requestsMap[req]
		ref.Metrics.ProcessResponse(req, resp)
		resp.Header.Add("Via", "solace-coverage-tool")
	}
}

func CreateProxy(covSpec *coverage.CovSpec) {
	spec = covSpec
	prx, _ := httpproxy.NewProxy()

	// Set handlers.
	prx.OnError = OnError
	prx.OnAccept = OnAccept
	prx.OnAuth = OnAuth
	prx.OnConnect = OnConnect
	prx.OnRequest = OnRequest
	prx.OnResponse = OnResponse

	// Listen...
	if err := http.ListenAndServe(":8080", prx); err != nil {
		log.Fatal(err)
	}
}
