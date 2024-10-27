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
	"os"

	"github.com/photowey/liquigen/internal/cmd/changelog"
	"github.com/photowey/liquigen/pkg/stringz"
	"github.com/spf13/cobra"
)

var (
	author           string
	email            string
	changeSetVersion string

	host     string
	port     int
	username string
	password string
	dialect  string
	database string
	format   string

	sqlFile string

	changelogCmd = &cobra.Command{
		Use:     "changelog",
		Aliases: []string{"chg", "chlog"},
		Short:   "Generate Database changelogs",
		Run: func(cmd *cobra.Command, args []string) {
			argz, err := populateArgs()
			if err != nil {
				panic(err)
			}

			if stringz.IsNotBlankString(sqlFile) {
				changelog.OnSQLMode(argz)

				return
			}

			changelog.OnDatabaseMode(argz)
		},
	}
)

func populateArgs() (*changelog.Args, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get current working directory failed: %v", err)
	}

	return &changelog.Args{
		Author:   author,
		Email:    email,
		Version:  changeSetVersion,
		Cwd:      cwd,
		Path:     cwd,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Dialect:  dialect,
		Database: database,
		Format:   format,
		SQLFile:  sqlFile,
	}, nil
}

func init() {
	changelogCmd.PersistentFlags().StringVarP(&author, "author", "a", "", "Author")
	changelogCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "Email")
	changelogCmd.PersistentFlags().StringVarP(&changeSetVersion, "version", "V", "", "Change set version")

	// Database mode
	changelogCmd.PersistentFlags().StringVarP(&host, "host", "H", "", "Target database host")
	changelogCmd.PersistentFlags().IntVarP(&port, "port", "P", 0, "Target database port")
	changelogCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Target database authentication username")
	changelogCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Target database authentication password")
	changelogCmd.PersistentFlags().StringVarP(&dialect, "dialect", "D", "", "Target database dialect")
	changelogCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "Target database name")
	changelogCmd.PersistentFlags().StringVarP(&dialect, "format", "f", "", "Target database changelog format")

	// SQL file mode
	changelogCmd.PersistentFlags().StringVarP(&sqlFile, "sql", "s", "", "SQL file")
}
