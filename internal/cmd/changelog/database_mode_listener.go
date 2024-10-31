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
)

// ----------------------------------------------------------------

func OnDatabaseMode(args *Args) {
	fmt.Printf("database mode: the Author: [%s]\n", args.Author)
	fmt.Printf("database mode: the Email: [%s]\n", args.Email)
	fmt.Printf("database mode: the Version: [%s]\n", args.Version)

	fmt.Printf("database mode: the Host: [%s]\n", args.Host)
	fmt.Printf("database mode: the Post: [%d]\n", args.Port)
	fmt.Printf("database mode: the Username: [%s]\n", args.Username)
	fmt.Printf("database mode: the Password: [%s]\n", args.Password)
	fmt.Printf("database mode: the Dialect: [%s]\n", args.Dialect)
	fmt.Printf("database mode: the Database name: [%s]\n", args.Database)

	fmt.Printf("database mode: the Format: [%s]\n", args.Format)

	fmt.Printf("database mode: the sqlFile: [%s]\n", args.SQLFile)

	fmt.Printf(red("database mode: unsupported now"))
}
