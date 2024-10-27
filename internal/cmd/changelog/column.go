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

// ----------------------------------------------------------------

const (
	PrimaryColumn = `<column name="{{ .Name }}" type="${type.{{ .Type }}}" remarks="{{ .Comment }}" autoIncrement="{{ .AutoIncrement }}">
                <constraints primaryKey="true" nullable="false"/>
            </column>`
	BigintColumn = `<column name="{{ .Name }}" type="${type.bigint}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	TinyIntColumn = `<column name="{{ .Name }}" type="${type.tinyint}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	SmallIntColumn = `<column name="{{ .Name }}" type="${type.smallint}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	MediumIntColumn = `<column name="{{ .Name }}" type="${type.mediumint}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	IntColumn = `<column name="{{ .Name }}" type="${type.int}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	FloatColumn = `<column name="{{ .Name }}" type="${type.float}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	DoubleColumn = `<column name="{{ .Name }}" type="${type.double}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	DecimalTemplate = `<column name="{{ .Name }}" type="${type.decimal}({{ .Precision }}, {{ .Scale }})" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	CharTemplate = `<column name="{{ .Name }}" type="${type.char}({{ .TypeLength }})" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	VarcharTemplate = `<column name="{{ .Name }}" type="${type.varchar}({{ .TypeLength }})" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	TextTemplate = `<column name="{{ .Name }}" type="${type.text}" remarks="{{ .Comment }}">
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	DateColumn = `<column name="{{ .Name }}" type="${type.date}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	TimeColumn = `<column name="{{ .Name }}" type="${type.time}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	DatetimeTemplate = `<column name="{{ .Name }}" type="${type.datetime}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	TimestampTemplate = `<column name="{{ .Name }}" type="${type.timestamp}" remarks="{{ .Comment }}" {{ .DefaultValue }}>
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
	UpdateTimestampTemplate = `<column name="{{ .Name }}" type="${type.timestamp}" remarks="{{ .Comment }}"
                    defaultValueComputed="CURRENT_TIMESTAMP">
                <constraints nullable="{{ .Nullable }}"/>
            </column>`
)
