module github.com/xgfone/go-opentelemetry/jaegerexporter

require (
	github.com/xgfone/go-opentelemetry v0.1.0
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/exporters/jaeger v1.8.0
	go.opentelemetry.io/otel/sdk v1.8.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.33.0 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
)

replace github.com/xgfone/go-opentelemetry => ../

go 1.17
