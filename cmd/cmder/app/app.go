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

	"github.com/photowey/liquigen/internal/version"
	"github.com/spf13/cobra"
)

const (
	LongText = "A lightweight Liquibase file generator cmd tool implemented in Golang."
)

var root = &cobra.Command{
	Use:     "liquigen",
	Short:   "A lightweight Liquibase changelog generator.",
	Long:    LongText,
	Version: version.Now(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome to liquigen cmder %s~", version.Now())
	},
}

func init() {
	cobra.OnInitialize(onInit)
	root.AddCommand(changelogCmd)
	root.AddCommand(configCmd)
	root.AddCommand(usageCmd)
	root.AddCommand(versionCmd)
}

func Run() {
	if err := root.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
