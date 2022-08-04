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

package otelhttpx

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// InstallClient wraps the default http client with OpenTelemetry.
func InstallClient(opts ...otelhttp.Option) {
	http.DefaultClient.Transport = RoundTripper(http.DefaultClient.Transport, opts...)
}

// RoundTripper returns a new http.RoundTripper wrapped by OpenTelemetry.
func RoundTripper(rt http.RoundTripper, opts ...otelhttp.Option) http.RoundTripper {
	switch rt.(type) {
	case *otelhttp.Transport:
		return rt
	case nil:
		return otelhttp.NewTransport(http.DefaultTransport, opts...)
	default:
		return otelhttp.NewTransport(rt, opts...)
	}
}

// Handler returns a new http.Handler wrapped by OpenTelemetry.
func Handler(handler http.Handler, operation string, opts ...otelhttp.Option) http.Handler {
	return otelhttp.NewHandler(handler, operation, opts...)
}
