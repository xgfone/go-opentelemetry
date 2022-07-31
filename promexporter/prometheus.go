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

package promexporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xgfone/go-opentelemetry"

	otprom "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/global"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/metric/export/aggregation"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

// Handler returns a http handler to export the prometheus metrics with http.
func Handler(gather prometheus.Gatherer) http.Handler {
	return promhttp.HandlerFor(gather, promhttp.HandlerOpts{})
}

// Install installs the prometheus as the metric exporter.
func Install(register prometheus.Registerer) (err error) {
	if register == nil {
		panic("prometheus.Registerer must not be nil")
	}

	factory := processor.NewFactory(
		selector.NewWithHistogramDistribution(),
		aggregation.CumulativeTemporalitySelector(),
		// processor.WithMemory(true),
	)
	options := []controller.Option{
		controller.WithResource(opentelemetry.DefaultResource),
	}

	ctrl := controller.New(factory, options...)
	exporter, err := otprom.New(otprom.Config{Registerer: register}, ctrl)
	if err == nil {
		global.SetMeterProvider(exporter.MeterProvider())
	}

	return
}
