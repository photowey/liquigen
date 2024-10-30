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
	"strings"

	"github.com/photowey/liquigen/internal/cmd/database/types"
)

// ----------------------------------------------------------------

type Context struct {
	Author  string
	Version string
	Date    string

	Dialect  string
	MySQL    string
	Postgres string
	SQLite   string

	Cwd  string
	Path string

	*Table
}

func NewContext() *Context {
	return &Context{}
}

type Table struct {
	Name    string
	Comment string

	Columns []*Column
	Indexes []*Index
}

type Column struct {
	Name         string
	Type         string
	Comment      string
	DefaultValue string

	AutoIncrement   bool
	Nullable        bool
	UpdateTimestamp bool

	TypeLength int
	Precision  int
	Scale      int

	PrimaryColumn string
	BigintColumn  string

	TinyIntColumn   string
	SmallIntColumn  string
	MediumIntColumn string
	IntColumn       string

	FloatColumn   string
	DoubleColumn  string
	DecimalColumn string

	CharColumn    string
	VarcharColumn string
	TextColumn    string

	DateColumn     string
	TimeColumn     string
	DatetimeColumn string

	TimestampColumn       string
	UpdateTimestampColumn string

	Dialect  string
	MySQL    string
	Postgres string
	SQLite   string
}

type Index struct {
	Name    string
	Columns []string
}

// ----------------------------------------------------------------

func (c *Column) testIsPrimaryColumn() bool {
	return c.AutoIncrement
}

func (c *Column) testIsNotPrimaryColumn() bool {
	return !c.testIsPrimaryColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsBigintColumn() bool {
	return strings.ToLower(c.Type) == types.BIGINT
}

func (c *Column) testIsNotBigintColumn() bool {
	return !c.testIsBigintColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsTinyIntColumn() bool {
	return strings.ToLower(c.Type) == types.TINYINT
}

func (c *Column) testIsNotTinyIntColumn() bool {
	return !c.testIsTinyIntColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsSmallIntColumn() bool {
	return strings.ToLower(c.Type) == types.SMALLINT
}

func (c *Column) testIsNotSmallIntColumn() bool {
	return !c.testIsSmallIntColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsMediumIntColumn() bool {
	return strings.ToLower(c.Type) == types.MEDIUMINT
}

func (c *Column) testIsNotMediumIntColumn() bool {
	return !c.testIsMediumIntColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsIntColumn() bool {
	return strings.ToLower(c.Type) == types.INT
}

func (c *Column) testIsNotIntColumn() bool {
	return !c.testIsIntColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsFloatColumn() bool {
	return strings.ToLower(c.Type) == types.FLOAT
}

func (c *Column) testIsNotFloatColumn() bool {
	return !c.testIsFloatColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsDoubleColumn() bool {
	return strings.ToLower(c.Type) == types.DOUBLE
}

func (c *Column) testIsNotDoubleColumn() bool {
	return !c.testIsFloatColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsDecimalColumn() bool {
	return strings.ToLower(c.Type) == types.DECIMAL
}

func (c *Column) testIsNotDecimalColumn() bool {
	return !c.testIsDecimalColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsCharColumn() bool {
	return strings.ToLower(c.Type) == types.CHAR
}

func (c *Column) testIsNotCharColumn() bool {
	return !c.testIsCharColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsVarcharColumn() bool {
	return strings.ToLower(c.Type) == types.VARCHAR
}

func (c *Column) testIsNotVarcharColumn() bool {
	return !c.testIsCharColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsTextColumn() bool {
	return strings.ToLower(c.Type) == types.TEXT
}

func (c *Column) testIsNotTextColumn() bool {
	return !c.testIsTextColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsDateColumn() bool {
	return strings.ToLower(c.Type) == types.DATE
}

func (c *Column) testIsNotDateColumn() bool {
	return !c.testIsDateColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsTimeColumn() bool {
	return strings.ToLower(c.Type) == types.TIME
}

func (c *Column) testIsNotTimeColumn() bool {
	return !c.testIsTimeColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsDatetimeColumn() bool {
	return strings.ToLower(c.Type) == types.DATETIME
}

func (c *Column) testIsNotDatetimeColumn() bool {
	return !c.testIsDatetimeColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsTimestampColumn() bool {
	return strings.ToLower(c.Type) == types.TIMESTAMP
}

func (c *Column) testIsNotTimestampColumn() bool {
	return !c.testIsDatetimeColumn()
}

// ----------------------------------------------------------------

func (c *Column) testIsUpdateTimestamp() bool {
	return c.UpdateTimestamp
}

func (c *Column) testIsNotUpdateTimestamp() bool {
	return !c.testIsUpdateTimestamp()
}
