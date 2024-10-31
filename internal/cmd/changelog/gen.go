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
	"fmt"
	"strings"
	"time"

	"github.com/photowey/liquigen/configs"
	"github.com/photowey/liquigen/internal/cmd/database/ast"
	"github.com/photowey/liquigen/internal/cmd/database/ast/parser/mysql"
	"github.com/photowey/liquigen/internal/cmd/database/ast/parser/postgres"
	"github.com/photowey/liquigen/pkg/stringz"
)

const (
	EmptyString = ""
)

func gen(args *Args) {
	doGenerate(args)

	cwd := args.Cwd
	fmt.Println(yellow("File: generated ->"), cyan("$pwd: "+cwd))
}

func doGenerate(args *Args) {
	db := configs.ConfigDatabase()
	excludes := db.Excludes
	includes := db.Includes

	astz := args.Ast
	databasePtr := astz.Database

	err := writeNormal(args)
	if err != nil {
		panic(err)
	}

	for _, tablePtr := range databasePtr.Tables {
		if len(excludes) > 0 && stringz.ArrayContains(excludes, tablePtr.Name) {
			continue
		}
		if len(includes) > 0 && stringz.ArrayNotContains(includes, tablePtr.Name) {
			continue
		}

		ctx := initCtx(args, tablePtr)
		populateCtx(args, ctx)

		changelog(args, ctx)

		if err = write(ctx); err != nil {
			fmt.Printf("liquigen: write tmpl failed, err:%v\n", err)
			return
		}
	}
}

func initCtx(args *Args, astTable *ast.Table) *Context {
	now := time.Now()
	layout := DatetimeLayout

	var columns []*Column

	for _, column := range astTable.Columns {
		tmp := populateInitColumn(column)
		c := populateTemplateColumn(args, tmp)

		columns = append(columns, c)
	}

	table := &Table{
		Name:    astTable.Name,
		Comment: astTable.Comment,
		Columns: columns,
		Indexes: []*Index{},
	}

	return &Context{
		Author:  args.Author,
		Date:    now.Format(layout),
		Version: args.Version,

		Dialect:  args.Dialect,
		MySQL:    mysql.Dialect,
		Postgres: postgres.Dialect,

		Cwd:  args.Cwd,
		Path: args.Path,

		Table: table,
	}
}

func populateTemplateColumn(args *Args, tmp *Column) *Column {
	dv := stringz.RemoveQuotes(tmp.DefaultValue)
	if stringz.IsNotBlankString(dv) {
		if tmp.testIsTimestampColumn() || tmp.testIsUpdateTimestamp() {
			dv = fmt.Sprintf("defaultValueComputed=\"%s\"", dv)
		} else if tmp.testIsDatetimeColumn() {
			dv = fmt.Sprintf("defaultValueDate=\"%s\"", dv)
		} else {
			dv = fmt.Sprintf("defaultValue=\"%s\"", dv)
		}
	}

	tmp.DefaultValue = dv

	c := &Column{
		Name:         tmp.Name,
		Type:         tmp.Type,
		Comment:      tmp.Comment,
		DefaultValue: dv,

		AutoIncrement: tmp.AutoIncrement,
		Nullable:      tmp.Nullable,

		TypeLength: tmp.TypeLength,
		Precision:  tmp.Precision,
		Scale:      tmp.Scale,

		PrimaryColumn: primaryColumn(tmp),

		BigintColumn:    bigintColumn(tmp),
		TinyIntColumn:   tinyIntColumn(tmp),
		SmallIntColumn:  smallIntColumn(tmp),
		MediumIntColumn: mediumIntColumn(tmp),
		IntColumn:       intColumn(tmp),

		FloatColumn:   floatColumn(tmp),
		DoubleColumn:  doubleColumn(tmp),
		DecimalColumn: decimalColumn(tmp),

		CharColumn:    charColumn(tmp),
		VarcharColumn: varcharColumn(tmp),
		TextColumn:    textColumn(tmp),

		DateColumn:            dateColumn(tmp),
		TimeColumn:            timeColumn(tmp),
		DatetimeColumn:        datetimeColumn(tmp),
		TimestampColumn:       timestampColumn(tmp),
		UpdateTimestampColumn: updateTimestampTemplate(tmp),

		Dialect:  args.Dialect,
		MySQL:    mysql.Dialect,
		Postgres: postgres.Dialect,
	}
	return c
}

func populateInitColumn(column *ast.Column) *Column {
	tmp := &Column{
		Name:         column.Name,
		Type:         strings.ToLower(column.DataType),
		Comment:      column.Comment,
		DefaultValue: column.Default,

		AutoIncrement:   column.AutoIncrement,
		Nullable:        !column.NotNull,
		UpdateTimestamp: column.UpdateTimestamp,

		TypeLength: toInt(column.Length, 0),
		Precision:  toInt(column.Precision, 10),
		Scale:      toInt(column.Scale, 0),
	}
	return tmp
}

// --------------------------------------------------------------------------------

func toInt(x *int, dv int) int {
	if x != nil {
		return *x
	} else {
		return dv
	}
}

// --------------------------------------------------------------------------------

func primaryColumn(column *Column) string {
	if column.testIsNotPrimaryColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, PrimaryTemplate)

	return string(bytes)
}

func bigintColumn(column *Column) string {
	if column.testIsNotBigintColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, BigintTemplate)

	return string(bytes)
}

func tinyIntColumn(column *Column) string {
	if column.testIsNotTinyIntColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, TinyIntTemplate)

	return string(bytes)
}

func smallIntColumn(column *Column) string {
	if column.testIsNotSmallIntColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, SmallIntTemplate)

	return string(bytes)
}

func mediumIntColumn(column *Column) string {
	if column.testIsNotMediumIntColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, MediumIntTemplate)

	return string(bytes)
}

func intColumn(column *Column) string {
	if column.testIsNotIntColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, IntTemplate)

	return string(bytes)
}

func floatColumn(column *Column) string {
	if column.testIsNotFloatColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, FloatTemplate)

	return string(bytes)
}

func doubleColumn(column *Column) string {
	if column.testIsNotDatetimeColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, DoubleColumn)

	return string(bytes)
}

func decimalColumn(column *Column) string {
	if column.testIsNotDecimalColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, DecimalTemplate)

	return string(bytes)
}

func charColumn(column *Column) string {
	if column.testIsNotCharColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, CharTemplate)

	return string(bytes)
}

func varcharColumn(column *Column) string {
	if column.testIsNotVarcharColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, VarcharTemplate)

	return string(bytes)
}

func textColumn(column *Column) string {
	if column.testIsNotTextColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, TextTemplate)

	return string(bytes)
}

func dateColumn(column *Column) string {
	if column.testIsNotDateColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, DateTemplate)

	return string(bytes)
}

func timeColumn(column *Column) string {
	if column.testIsNotTimeColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, TimeTemplate)

	return string(bytes)
}

func datetimeColumn(column *Column) string {
	if column.testIsNotDatetimeColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, DatetimeTemplate)

	return string(bytes)
}

func timestampColumn(column *Column) string {
	if column.testIsNotTimestampColumn() {
		return EmptyString
	}
	bytes, _ := parseField(column, TimestampTemplate)

	return string(bytes)
}

func updateTimestampTemplate(column *Column) string {
	if column.testIsNotUpdateTimestamp() {
		return EmptyString
	}
	bytes, _ := parseField(column, UpdateTimestampTemplate)

	return string(bytes)
}

// --------------------------------------------------------------------------------

func populateCtx(args *Args, ctx *Context) {}

func changelog(args *Args, ctx *Context) {}
