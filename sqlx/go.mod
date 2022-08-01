module github.com/xgfone/go-opentelemetry/sqlx

require (
	github.com/XSAM/otelsql v0.15.0
	github.com/xgfone/go-opentelemetry v0.1.0
	github.com/xgfone/sqlx v0.21.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/xgfone/cast v0.5.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.33.0 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/metric v0.31.0 // indirect
	go.opentelemetry.io/otel/sdk v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
)

replace github.com/xgfone/go-opentelemetry => ../

go 1.17
