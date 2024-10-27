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

package database

import (
	"testing"
)

func TestRemoveComments(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test /*...*/",
			args: args{sql: `
/* xxx */
/* 
  1.111 
  2.222 
  3.333 
*/
SELECT * FROM hello;
`},
			want: "SELECT * FROM hello;",
		},
		{
			name: "test --",
			args: args{sql: `
-- comment
-- comment
SELECT * FROM hello;
`},
			want: "SELECT * FROM hello;",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveComments(tt.args.sql); got != tt.want {
				t.Errorf("RemoveComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
