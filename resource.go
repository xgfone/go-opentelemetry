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

package opentelemetry

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

// DefaultResource is the default resource used by metric and tracer.
var DefaultResource *resource.Resource

func init() {
	hostResource, err := resource.New(context.Background(), resource.WithHost())
	if err != nil {
		panic(err)
	}

	DefaultResource, err = resource.Merge(resource.Default(), hostResource)
	if err != nil {
		panic(err)
	}

	http.DefaultClient = HTTPClient(http.DefaultTransport)
}

// SetServiceName sets the service name in the resource.
func SetServiceName(serviceName string) {
	opt := resource.WithAttributes(semconv.ServiceNameKey.String(serviceName))
	rsc, err := resource.New(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	DefaultResource, err = resource.Merge(DefaultResource, rsc)
	if err != nil {
		panic(err)
	}
}
