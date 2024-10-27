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

package stringz

import (
	"fmt"
	"os"
	"strings"
)

func String(source any) string {
	return fmt.Sprintf("%v", source)
}

func ReplaceTemplate(template string, args ...any) string {
	return fmt.Sprintf(template, args...)
}

func IsBlankString(str string) bool {
	return "" == str
}

func IsNotBlankString(str string) bool {
	return !IsBlankString(str)
}

func IsEmptyStringSlice(target []string) bool {
	return len(target) == 0
}

func IsNotEmptyStringSlice(target []string) bool {
	return !IsEmptyStringSlice(target)
}

// ----------------------------------------------------------------

func Tail(content, separator string) string {
	if IsBlankString(content) {
		return ""
	}
	items := strings.Split(content, separator)

	return items[len(items)-1]
}

// ----------------------------------------------------------------

func Pascal(content string) string {
	if IsBlankString(content) {
		return ""
	}
	items := strings.ToUpper(content[0:1]) + content[1:]

	return items
}

// ----------------------------------------------------------------

func ToPath(content, separator string) string {
	if IsBlankString(content) {
		return ""
	}

	pathSeparator := string(os.PathSeparator)
	// if isWindows() {
	// 	pathSeparator += pathSeparator
	// }
	items := strings.ReplaceAll(content, separator, pathSeparator)

	return items
}

// ----------------------------------------------------------------

func ArrayContains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}

	return false
}

func ArrayNotContains(haystack []string, needle string) bool {
	return !ArrayContains(haystack, needle)
}

// ----------------------------------------------------------------

func ToProjectPath(content string) string {
	if IsBlankString(content) {
		return ""
	}

	return ToPath(content, "-")
}

// ----------------------------------------------------------------

func Fields(source string) []string {
	var tokens []string
	var current []rune

	insideQuotes := false
	quoteChar := rune(0)

	for _, ch := range source {
		if insideQuotes {
			if ch == quoteChar {
				insideQuotes = false
				current = append(current, ch)
				tokens = append(tokens, string(current))
				current = []rune{}
			} else {
				current = append(current, ch)
			}
		} else {
			switch ch {
			case ' ', '\t', '\n', '\r':
				if len(current) > 0 {
					tokens = append(tokens, string(current))
					current = []rune{}
				}
			case '(', ')', ',', ';', '=':
				if len(current) > 0 {
					tokens = append(tokens, string(current))
					current = []rune{}
				}
				tokens = append(tokens, string(ch))
			case '\'', '"':
				if len(current) > 0 {
					tokens = append(tokens, string(current))
					current = []rune{}
				}
				insideQuotes = true
				quoteChar = ch
				current = append(current, ch)
			default:
				current = append(current, ch)
			}
		}
	}

	if len(current) > 0 {
		tokens = append(tokens, string(current))
	}

	return tokens
}

// RemoveQuotes removes quotes from a string.
func RemoveQuotes(source string) string {
	if (strings.HasPrefix(source, "'") && strings.HasSuffix(source, "'")) ||
		(strings.HasPrefix(source, "\"") && strings.HasSuffix(source, "\"")) ||
		(strings.HasPrefix(source, "`") && strings.HasSuffix(source, "`")) {
		return source[1 : len(source)-1]
	}

	return source
}
