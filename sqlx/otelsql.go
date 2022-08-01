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

package sqlx

import (
	"database/sql"

	"github.com/XSAM/otelsql"
	"github.com/xgfone/go-opentelemetry"
	"github.com/xgfone/sqlx"
)

// Opener returns the sql opener to open a *sql.DB.
func Opener(opts ...otelsql.Option) sqlx.Opener {
	return func(driverName, dataSourceName string) (*sql.DB, error) {
		return otelsql.Open(driverName, dataSourceName, opts...)
	}
}

// Install replaces the default opener of sqlx.DefaultOpener with the new opener
// wrapped by OpenTelemetry.
func Install(opts ...otelsql.Option) {
	attrs := opentelemetry.DefaultResource.Attributes()
	opts = append(opts, otelsql.WithAttributes(attrs...))
	sqlx.DefaultOpener = Opener(opts...)
}

func init() { Install() }
