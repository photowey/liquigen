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
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/photowey/liquigen/pkg/stringz"
)

const (
	FixedNameTemplate = "template_employee_1.0.0"
	DatetimeLayout    = "20060102"
	TmplSuffix        = ".tmpl"
	templatePackr2    = "changelogs"
	templatePackr2Dir = "./templates"
)

// ----------------------------------------------------------------

//go:generate packr2
func write(ctx *Context) (err error) {
	path := ctx.Path

	box := packr.New(templatePackr2, templatePackr2Dir)
	if err = os.MkdirAll(path, 0o755); err != nil {
		return
	}

	err = writeTemplate(ctx, box)
	if err != nil {
		return err
	}

	return
}

func writeNormal(ctx *Args) (err error) {
	path := ctx.Path

	box := packr.New(templatePackr2, templatePackr2Dir)
	if err = os.MkdirAll(path, 0o755); err != nil {
		return
	}

	for _, item := range box.List() {
		tmpl, _ := box.FindString(item)

		tmpItem := item

		if testIsTargetTmplFile(tmpItem) {
			continue
		}

		i := strings.LastIndex(tmpItem, string(os.PathSeparator))
		if i > 0 {
			dir := tmpItem[:i]
			if err = os.MkdirAll(filepath.Join(path, dir), 0o755); err != nil {
				return
			}
		}

		tmpItem = strings.TrimSuffix(tmpItem, TmplSuffix)
		if err = doWriteOriginalFile(item, ctx.Cwd, filepath.Join(path, tmpItem), tmpl); err != nil {
			return
		}
	}

	return
}

func writeTemplate(ctx *Context, box *packr.Box) (err error) {
	path := ctx.Path

	for _, item := range box.List() {
		tmpl, _ := box.FindString(item)

		tmpItem := item
		if testIsNotTargetTmplFile(item) {
			continue
		}

		// template_employee_1.0.0.xml.tmpl
		i := strings.LastIndex(tmpItem, string(os.PathSeparator))
		if i > 0 {
			dir := tmpItem[:i]

			tmpItem = strings.ReplaceAll(
				tmpItem,
				stringz.ToProjectPath(FixedNameTemplate),
				stringz.ToProjectPath(fmt.Sprintf("%s_%s", ctx.Table.Name, ctx.Version)))

			if err = os.MkdirAll(filepath.Join(path, dir), 0o755); err != nil {
				return
			}
		}

		tmpItem = strings.ReplaceAll(
			tmpItem,
			stringz.ToProjectPath(FixedNameTemplate),
			stringz.ToProjectPath(fmt.Sprintf("%s_%s", ctx.Table.Name, ctx.Version)))

		tmpItem = strings.TrimSuffix(tmpItem, TmplSuffix)
		if err = doWriteFile(ctx, item, filepath.Join(path, tmpItem), tmpl); err != nil {
			return
		}
	}

	return
}

func testIsTargetTmplFile(target string) bool {
	return strings.Contains(target, FixedNameTemplate)
}

func testIsNotTargetTmplFile(target string) bool {
	return !testIsTargetTmplFile(target)
}

func doWriteFile(ctx *Context, item, path, tmpl string) (err error) {
	content, err := tryParseIfNecessary(ctx, item, tmpl)
	if err != nil {
		return
	}

	// Twice?
	// fmt.Println(yellow("File: generated ->"), cyan("$pwd"+path[len(ctx.Cwd):]))

	content = replaceSpace(content)

	return os.WriteFile(path, []byte(content), 0o755)
}

func doWriteOriginalFile(item, cwd, path, tmpl string) (err error) {
	content, err := tryParseIfNecessaryOriginal(item, tmpl)
	if err != nil {
		return
	}

	// Twice?
	// fmt.Println(yellow("Original File: generated ->"), cyan("$pwd"+path[len(cwd):]))

	content = replaceSpace(content)

	return os.WriteFile(path, []byte(content), 0o755)
}

func tryParseIfNecessary(ctx *Context, item, tmpl string) (string, error) {
	return doParseIfNecessary(ctx, item, tmpl)
}

func tryParseIfNecessaryOriginal(item, tmpl string) (string, error) {
	return doParseIfNecessary(NewContext(), item, tmpl)
}

func doParseIfNecessary(ctx *Context, item, tmpl string) (string, error) {
	if testIsTargetTmplFile(item) {
		bytes, err := parseTmpl(ctx, tmpl)
		if err != nil {
			return EmptyString, err
		}
		return string(bytes), nil
	}
	return tmpl, nil
}

func replaceSpace(content string) string {
	content = strings.TrimSpace(content)
	re := regexp.MustCompile(`\n{2,}`)
	x := re.ReplaceAllString(content, "\n")

	return strings.ReplaceAll(x, " >", ">")
}
