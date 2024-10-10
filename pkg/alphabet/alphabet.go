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

package alphabet

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	EmptyString = ""
)

func CamelCase(src string) string {
	if src == EmptyString {
		return EmptyString
	}

	return strings.ToLower(src[:1]) + src[1:]
}

func Snake2Pascal(snakeCase string) string {
	snakeCase = strings.Replace(snakeCase, "_", " ", -1)
	snakeCase = cases.Title(language.Dutch).String(snakeCase)
	return strings.Replace(snakeCase, " ", "", -1)
}

func Snake2Camel(snakeCase string) string {
	return CamelCase(Snake2Pascal(snakeCase))
}

func PascalCase(src string) string {
	if src == EmptyString {
		return EmptyString
	}

	return strings.ToUpper(strings.ToLower(src[:1])) + src[1:]
}

func SnakeCase(src string) string {
	if src == EmptyString {
		return src
	}
	srcLen := len(src)
	result := make([]byte, 0, srcLen*2)
	caseSymbol := false
	for i := 0; i < srcLen; i++ {
		char := src[i]
		if i > 0 && char >= 'A' && char <= 'Z' && caseSymbol { // _xxx || yyy__zzz
			result = append(result, '_')
		}
		caseSymbol = char != '_'

		result = append(result, char)
	}

	snakeCase := strings.ToLower(string(result))

	return snakeCase
}

func CleanTableComment(comment string, tableName string) string {
	return strings.ReplaceAll(comment, "("+tableName+")", "")
}
