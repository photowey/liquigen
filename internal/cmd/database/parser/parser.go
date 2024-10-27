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

package parser

import (
	"github.com/photowey/liquigen/internal/cmd/database/ast"
)

// ----------------------------------------------------------------

var (
	_registry          = NewRegistry()
	_         Registry = (*registry)(nil)
)

// ----------------------------------------------------------------

type Parser interface {
	Dialect() string
	Parse(string) (ast.Database, error)
}

type Registry interface {
	Register(parser Parser)
	Acquire(dialect string) (Parser, bool)
	Contains(dialect string) bool
}

// ----------------------------------------------------------------

type registry struct {
	parsers map[string]Parser
}

func NewRegistry() Registry {
	return &registry{
		parsers: make(map[string]Parser),
	}
}

// ----------------------------------------------------------------

func (r *registry) Register(parser Parser) {
	r.parsers[parser.Dialect()] = parser
}

func (r *registry) Acquire(dialect string) (Parser, bool) {
	parser, ok := r.parsers[dialect]

	return parser, ok
}

func (r *registry) Contains(dialect string) bool {
	_, ok := r.parsers[dialect]

	return ok
}

// ----------------------------------------------------------------

func Register(parser Parser) {
	_registry.Register(parser)
}

func Acquire(dialect string) (Parser, bool) {
	return _registry.Acquire(dialect)
}

func Contains(dialect string) bool {
	return _registry.Contains(dialect)
}
