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
	"time"

	"github.com/manifoldco/promptui"
	"github.com/photowey/liquigen/internal/cmd/database"
	"github.com/photowey/liquigen/internal/cmd/database/ast/parser"
	"github.com/photowey/liquigen/pkg/filez"
	"github.com/photowey/liquigen/pkg/stringz"
)

func OnSQLMode(args *Args) {
	// report(args)

	sql, err := readSQL(args.SQLFile)
	if err != nil {
		panic(err)
	}

	args.SQL = sql
	// reportSQL(args)

	if sqlParser, ok := parser.Acquire(args.Dialect); ok {
		parseSQL(sqlParser, args)
		confirm(args)

		gen(args)

		return

	}

	panic(fmt.Errorf("the dialect %s not found", args.Dialect))
}

func parseSQL(sqlParser parser.Parser, args *Args) {
	_ast, err := sqlParser.Parse(args.SQL)
	if err != nil {
		panic(err)
	}

	args.Ast = _ast
}

func readSQL(sqlFile string) (string, error) {
	sqlFile, err := filez.Clean(sqlFile)
	if err != nil {
		return "", err
	}
	sql, err := os.ReadFile(sqlFile)
	if err != nil {
		return "", err
	}

	return string(sql), nil
}

// ----------------------------------------------------------------

func confirm(args *Args) {
	validateInput(args)
	confirmInput(args)
}

func confirmInput(args *Args) {
	now := time.Now()
	layout := "2006/01/02"

	fmt.Println("")
	fmt.Println(green("---------------- $ start liquigen changelog input report ----------------"))
	fmt.Println(blue("Project path:"), args.Path)
	fmt.Println(blue("Project author:"), args.Author)

	if stringz.IsBlankString(args.Email) {
		fmt.Println(yellow("Project author's email:"), "-")
	} else {
		fmt.Println(blue("Project author's email:"), args.Email)
	}

	fmt.Println(blue("Project date:"), now.Format(layout))
	fmt.Println(cyan("Project version:"), args.Version)
	fmt.Println(green("---------------- $ end liquigen changelog input report ----------------"))
	fmt.Println("")

	prompt := promptui.Prompt{
		Label:     "Project info confirm",
		IsConfirm: true,
		Default:   "Y",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Confirm failed: %v\n", err)
		return
	}

	fmt.Printf("You choose: %q\n", result)
}

func validateInput(args *Args) {
	validatePath(args)

	validateAuthor(args)
	validateVersion(args)

	validateDialect(args)
}

// ----------------------------------------------------------------

// Deprecated
func reportSQL(args *Args) {
	fmt.Printf("the SQL file is: \n%s\n", args.SQLFile)
	fmt.Printf("the SQL file content is: \n%s\n", args.SQL)
	fmt.Printf("the SQL file clean content is: \n%s\n", database.RemoveComments(args.SQL))
}

// Deprecated
func report(args *Args) {
	fmt.Printf("sql mode: the Author: [%s]\n", args.Author)
	fmt.Printf("sql mode: the Email: [%s]\n", args.Email)
	fmt.Printf("sql mode: the Version: [%s]\n", args.Version)

	fmt.Printf("sql mode: the Host: [%s]\n", args.Host)
	fmt.Printf("sql mode: the Post: [%d]\n", args.Port)
	fmt.Printf("sql mode: the Username: [%s]\n", args.Username)
	fmt.Printf("sql mode: the Password: [%s]\n", args.Password)
	fmt.Printf("sql mode: the Dialect: [%s]\n", args.Dialect)
	fmt.Printf("sql mode: the Database name: [%s]\n", args.Database)

	fmt.Printf("sql mode: the Format: [%s]\n", args.Format)

	fmt.Printf("sql mode: the sqlFile: [%s]\n", args.SQLFile)
}
