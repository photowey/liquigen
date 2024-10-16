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
	"path/filepath"
	"strings"

	"github.com/photowey/liquigen/configs"
	"github.com/photowey/liquigen/internal/home"
	"github.com/photowey/liquigen/pkg/filez"
)

func onInit() {
	home.AppHome()
	tryLoadConfig()
}

func tryLoadConfig() {
	liquigenHome := home.Dir
	configFile := filepath.Join(liquigenHome, strings.ToLower(home.LiquigenConfigFile))

	if filez.FileExists(liquigenHome, home.LiquigenConfigFile) {
		configs.Init(configFile)

		return
	}

	fmt.Printf("the liquigen config file not exists")
}
