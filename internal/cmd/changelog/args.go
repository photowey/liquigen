/*
 * Copyright © 2024 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package changelog

import (
	"github.com/photowey/liquigen/internal/cmd/database/ast"
)

// ----------------------------------------------------------------

type Args struct {
	Author  string
	Email   string
	Version string
	Cwd     string
	Path    string

	Host     string
	Port     int
	Username string
	Password string
	Dialect  string
	Database string

	Format string

	SQLFile string
	SQL     string
	Ast     *ast.Ast
}
