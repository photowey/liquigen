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

	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/photowey/liquigen/configs"
	"github.com/photowey/liquigen/internal/cmd/database/ast/parser/mysql"
	"github.com/photowey/liquigen/pkg/stringz"
)

// ----------------------------------------------------------------

const (
	VersionTemplate = "1.0.0"
)

// ----------------------------------------------------------------

var (
	green  = color.FgGreen.Render
	blue   = color.FgBlue.Render
	yellow = color.FgYellow.Render
	cyan   = color.FgCyan.Render
	red    = color.FgRed.Render
)

var promptTemplates = &promptui.PromptTemplates{
	Prompt:          fmt.Sprintf(`{{ "%s" | bold }} {{ . | bold }}{{ ": " | bold}}`, promptui.IconInitial),
	Valid:           fmt.Sprintf(`{{ "%s" | cyan }} {{ . | cyan }}{{ ": " | cyan}}`, promptui.IconGood),
	Invalid:         fmt.Sprintf(`{{ "%s" | yellow }} {{ . | yellow }}{{ ": " | yellow}}`, promptui.IconBad),
	ValidationError: fmt.Sprintf(`{{ ">>" | red }} {{ . | red }}{{ ": " | bold}}`),
	Success:         fmt.Sprintf(`{{ "%s" | bold }} {{ . | faint }}{{ ": " | bold}}`, promptui.IconGood),
}

// ----------------------------------------------------------------

func validatePath(args *Args) {
	if stringz.IsBlankString(args.Path) {
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty output dir")
			}
			return nil
		}

		defaultPath := args.Cwd

		prompt := promptui.Prompt{
			Label:     "Output dir",
			Validate:  validate,
			Templates: promptTemplates,
			Default:   defaultPath,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct project dir: %v\n", err)

			// Loop
			validatePath(args)
		}

		args.Path = result
	}
}

func validateAuthor(args *Args) {
	if stringz.IsBlankString(args.Author) {
		project := configs.ConfigProject()
		if stringz.IsNotBlankString(project.Author) {
			args.Author = project.Author
			return
		}
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty author's name")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:     "Author",
			Validate:  validate,
			Templates: promptTemplates,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct author of project: %v\n", err)

			// Loop
			validateAuthor(args)
		}

		args.Author = result
	}
}

func validateVersion(args *Args) {
	if stringz.IsBlankString(args.Version) {
		project := configs.ConfigProject()

		if stringz.IsNotBlankString(project.Version) {
			args.Version = project.Version
			return
		}
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty project's version")
			}
			return nil
		}

		defaultVersion := VersionTemplate

		prompt := promptui.Prompt{
			Label:     "Version",
			Validate:  validate,
			Templates: promptTemplates,
			Default:   defaultVersion,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct version of project: %v\n", err)

			// Loop
			validateVersion(args)
		}

		args.Version = result
	}
}

func validateDialect(args *Args) {
	if stringz.IsBlankString(args.Dialect) {
		project := configs.ConfigProject()

		if stringz.IsNotBlankString(project.Dialect) {
			args.Dialect = project.Dialect
			return
		}
		validate := func(input string) error {
			if stringz.IsBlankString(input) {
				return fmt.Errorf("empty project's dialect")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:     "Dialect",
			Validate:  validate,
			Templates: promptTemplates,
			Default:   mysql.Dialect,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Please enter the correct dialect of project: %v\n", err)

			// Loop
			validateDialect(args)
		}

		args.Dialect = result
	}
}
