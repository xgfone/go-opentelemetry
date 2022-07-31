// Copyright 2022 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jaegerexporter

import (
	"log"
	"net"
	"net/http"

	"github.com/xgfone/go-opentelemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewAgentEndpoint returns a new agent endpoint option with the address
// of the jaeger agent.
func NewAgentEndpoint(addr string) jaeger.EndpointOption {
	opts := []jaeger.AgentEndpointOption{
		jaeger.WithLogger(log.Default()),
	}

	if addr != "" {
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			host, port = addr, "6831"
		}

		opts = append(opts, jaeger.WithAgentHost(host), jaeger.WithAgentPort(port))
	}

	return jaeger.WithAgentEndpoint(opts...)
}

// NewCollectorEndpoint returns a new collector endpoint option.
func NewCollectorEndpoint(c *http.Client, url, username, password string) jaeger.EndpointOption {
	opts := make([]jaeger.CollectorEndpointOption, 4)
	if c != nil {
		opts = append(opts, jaeger.WithHTTPClient(c))
	}
	if url != "" {
		opts = append(opts, jaeger.WithEndpoint(url))
	}
	if username != "" && password != "" {
		opts = append(opts, jaeger.WithUsername(username), jaeger.WithPassword(password))
	}

	return jaeger.WithCollectorEndpoint(opts...)
}

// SpanProcessorFunc is a function to new a tracer span processor.
type SpanProcessorFunc func(trace.SpanExporter) trace.SpanProcessor

// BatchSpanProcessor returns a new batch span processor.
func BatchSpanProcessor(opts ...trace.BatchSpanProcessorOption) SpanProcessorFunc {
	return func(exporter trace.SpanExporter) trace.SpanProcessor {
		return trace.NewBatchSpanProcessor(exporter, opts...)
	}
}

// SimpleSpanProcessor returns a new simple span processor.
func SimpleSpanProcessor() SpanProcessorFunc {
	return func(exporter trace.SpanExporter) trace.SpanProcessor {
		return trace.NewSimpleSpanProcessor(exporter)
	}
}

// Install installs the jaeger as the tracer exporter.
func Install(jaegerEndpoint jaeger.EndpointOption, spanProcessor SpanProcessorFunc,
	opts ...trace.TracerProviderOption) (err error) {
	if jaegerEndpoint == nil {
		jaegerEndpoint = jaeger.WithAgentEndpoint(jaeger.WithLogger(log.Default()))
	}

	exporter, err := jaeger.New(jaegerEndpoint)
	if err == nil {
		if spanProcessor == nil {
			spanProcessor = BatchSpanProcessor()
		}

		opts = append([]trace.TracerProviderOption{
			trace.WithSpanProcessor(spanProcessor(exporter)),
			trace.WithResource(opentelemetry.DefaultResource),
		}, opts...)

		otel.SetTracerProvider(trace.NewTracerProvider(opts...))
	}

	return
}
