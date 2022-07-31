# go-opentelemetry [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/go-opentelemetry)](https://pkg.go.dev/github.com/xgfone/go-opentelemetry) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/go-opentelemetry/master/LICENSE)

A library to install the opentelemetry exporters rapidly.


## Example
```go.mod
module myapp

require (
	github.com/prometheus/client_golang v1.12.2
	github.com/xgfone/go-opentelemetry v0.1.0
	github.com/xgfone/go-opentelemetry/jaegerexporter v0.1.1
	github.com/xgfone/go-opentelemetry/promexporter v0.1.1
)

go 1.17
```

```go
package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/xgfone/go-opentelemetry"
	"github.com/xgfone/go-opentelemetry/jaegerexporter"
	"github.com/xgfone/go-opentelemetry/promexporter"
)

func init() {
	registry := prometheus.NewRegistry()
	http.Handle("/metrics", promexporter.Handler(registry))

	opentelemetry.SetServiceName("ServiceName")
	jaegerexporter.Install(nil, nil)
	promexporter.Install(registry)
}

func wrapHandler(serverName string, handler http.HandlerFunc) http.Handler {
	return opentelemetry.Handler(handler, serverName)
}

func main() {
	http.Handle("/path", wrapHandler("localhost", func(w http.ResponseWriter, r *http.Request) {
		// TODO
	}))

	http.ListenAndServe(":80", nil)
}
```
