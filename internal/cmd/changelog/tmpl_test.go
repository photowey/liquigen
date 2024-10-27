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
	"reflect"
	"testing"
)

func Test_parseTmpl(t *testing.T) {
	type args struct {
		tmpl string
		ctx  *Column
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test Primary column",
			args: args{
				tmpl: PrimaryColumn,
				ctx: &Column{
					Name:          "id",
					Type:          "bigint",
					Comment:       "Primary Key",
					AutoIncrement: true,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="id" type="${type.bigint}" remarks="Primary Key" autoIncrement="true">
                <constraints primaryKey="true" nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Bigint column",
			args: args{
				tmpl: BigintColumn,
				ctx: &Column{
					Name:          "user_id",
					Type:          "bigint",
					Comment:       "Users ID",
					DefaultValue:  "defaultValue=\"1730034642683\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="user_id" type="${type.bigint}" remarks="Users ID" defaultValue="1730034642683">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test TinyInt column",
			args: args{
				tmpl: TinyIntColumn,
				ctx: &Column{
					Name:          "order_state",
					Type:          "tinyint",
					Comment:       "Order State",
					DefaultValue:  "defaultValue=\"0\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="order_state" type="${type.tinyint}" remarks="Order State" defaultValue="0">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test SmallInt column",
			args: args{
				tmpl: SmallIntColumn,
				ctx: &Column{
					Name:          "order_state",
					Type:          "smallint",
					Comment:       "Order State",
					DefaultValue:  "defaultValue=\"0\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="order_state" type="${type.smallint}" remarks="Order State" defaultValue="0">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test MediumInt column",
			args: args{
				tmpl: MediumIntColumn,
				ctx: &Column{
					Name:          "order_state",
					Type:          "smallint",
					Comment:       "Order State",
					DefaultValue:  "defaultValue=\"1\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="order_state" type="${type.mediumint}" remarks="Order State" defaultValue="1">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Int column",
			args: args{
				tmpl: IntColumn,
				ctx: &Column{
					Name:          "member_count",
					Type:          "smallint",
					Comment:       "Member Count",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="member_count" type="${type.int}" remarks="Member Count" >
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Float column",
			args: args{
				tmpl: FloatColumn,
				ctx: &Column{
					Name:          "price",
					Type:          "float",
					Comment:       "Price",
					DefaultValue:  "defaultValue=\"88.48\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="price" type="${type.float}" remarks="Price" defaultValue="88.48">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Double column",
			args: args{
				tmpl: DoubleColumn,
				ctx: &Column{
					Name:          "price",
					Type:          "double",
					Comment:       "Price",
					DefaultValue:  "defaultValue=\"88.48\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="price" type="${type.double}" remarks="Price" defaultValue="88.48">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Decimal column",
			args: args{
				tmpl: DecimalTemplate,
				ctx: &Column{
					Name:          "price",
					Type:          "decimal",
					Comment:       "Price",
					DefaultValue:  "defaultValue=\"88.48\"",
					AutoIncrement: false,
					Nullable:      false,
					Precision:     16,
					Scale:         2,
				},
			},
			want: []byte(`<column name="price" type="${type.decimal}(16, 2)" remarks="Price" defaultValue="88.48">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Char column",
			args: args{
				tmpl: CharTemplate,
				ctx: &Column{
					Name:          "member_no",
					Type:          "char",
					Comment:       "Member No.",
					DefaultValue:  "defaultValue=\"00000000000000001\"",
					AutoIncrement: false,
					Nullable:      true,
					TypeLength:    18,
				},
			},
			want: []byte(`<column name="member_no" type="${type.char}(18)" remarks="Member No." defaultValue="00000000000000001">
                <constraints nullable="true"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Varchar column",
			args: args{
				tmpl: VarcharTemplate,
				ctx: &Column{
					Name:          "member_no",
					Type:          "varchar",
					Comment:       "Member No.",
					DefaultValue:  "defaultValue=\"00000000000000001\"",
					AutoIncrement: false,
					Nullable:      true,
					TypeLength:    18,
				},
			},
			want: []byte(`<column name="member_no" type="${type.varchar}(18)" remarks="Member No." defaultValue="00000000000000001">
                <constraints nullable="true"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Text column",
			args: args{
				tmpl: TextTemplate,
				ctx: &Column{
					Name:          "remark",
					Type:          "text",
					Comment:       "Remark",
					AutoIncrement: false,
					Nullable:      true,
				},
			},
			want: []byte(`<column name="remark" type="${type.text}" remarks="Remark">
                <constraints nullable="true"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Date column",
			args: args{
				tmpl: DateColumn,
				ctx: &Column{
					Name:          "create_date",
					Type:          "date",
					Comment:       "Create date",
					DefaultValue:  "defaultValue=\"2024-10-27\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="create_date" type="${type.date}" remarks="Create date" defaultValue="2024-10-27">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test Time column",
			args: args{
				tmpl: TimeColumn,
				ctx: &Column{
					Name:          "create_time",
					Type:          "time",
					Comment:       "Create time",
					DefaultValue:  "defaultValue=\"00:00:00\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="create_time" type="${type.time}" remarks="Create time" defaultValue="00:00:00">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test DateTime column",
			args: args{
				tmpl: DatetimeTemplate,
				ctx: &Column{
					Name:          "create_datetime",
					Type:          "datetime",
					Comment:       "Create Datetime",
					DefaultValue:  "defaultValueDate=\"2024-10-27 00:00:00\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="create_datetime" type="${type.datetime}" remarks="Create Datetime" defaultValueDate="2024-10-27 00:00:00">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test TimeStamp column",
			args: args{
				tmpl: TimestampTemplate,
				ctx: &Column{
					Name:          "create_timestamp",
					Type:          "timestamp",
					Comment:       "Create TimeStamp",
					DefaultValue:  "defaultValueComputed=\"CURRENT_TIMESTAMP\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="create_timestamp" type="${type.timestamp}" remarks="Create TimeStamp" defaultValueComputed="CURRENT_TIMESTAMP">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
		{
			name: "test update TimeStamp column",
			args: args{
				tmpl: UpdateTimestampTemplate,
				ctx: &Column{
					Name:          "update_time",
					Type:          "timestamp",
					Comment:       "Update time",
					DefaultValue:  "defaultValueComputed=\"CURRENT_TIMESTAMP\"",
					AutoIncrement: false,
					Nullable:      false,
				},
			},
			want: []byte(`<column name="update_time" type="${type.timestamp}" remarks="Update time"
                    defaultValueComputed="CURRENT_TIMESTAMP">
                <constraints nullable="false"/>
            </column>`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTmpl(tt.args.tmpl, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTmpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTmpl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
