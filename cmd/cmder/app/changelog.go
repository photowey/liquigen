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

package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	host     string
	port     int32
	username string
	password string
	dialect  string
	database string
	format   string
	author   string

	changelogCmd = &cobra.Command{
		Use:     "changelog",
		Aliases: []string{"ch", "chge"},
		Short:   "Generate Database changelogs",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("the Host: [%s]\n", host)
			fmt.Printf("the Post: [%d]\n", port)
			fmt.Printf("the Username: [%s]\n", username)
			fmt.Printf("the Password: [%s]\n", password)
			fmt.Printf("the Dialect: [%s]\n", dialect)
			fmt.Printf("the Database name: [%s]\n", database)
			fmt.Printf("the Format: [%s]\n", format)
			fmt.Printf("the Author: [%s]\n", author)
		},
	}
)

func init() {
	changelogCmd.PersistentFlags().StringVarP(&host, "host", "H", "127.0.0.1", "Target database host")
	changelogCmd.PersistentFlags().Int32VarP(&port, "port", "P", 5432, "Target database port")
	changelogCmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "Target database authentication username")
	changelogCmd.PersistentFlags().StringVarP(&password, "password", "p", "root", "Target database authentication password")
	changelogCmd.PersistentFlags().StringVarP(&dialect, "dialect", "D", "mysql", "Target database dialect")
	changelogCmd.PersistentFlags().StringVarP(&database, "database", "d", "hello_world", "Target database name")
	changelogCmd.PersistentFlags().StringVarP(&dialect, "format", "f", "xml", "Target database changelog format")
	changelogCmd.PersistentFlags().StringVarP(&author, "author", "a", "admin", "Author")
}
