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

package home

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/photowey/liquigen/pkg/filez"
)

const (
	LiquigenConfigFile = "liquigen.json"
)

var (
	Home   = ".liquigen"
	Usr, _ = user.Current()
	Dir    = filepath.Join(Usr.HomeDir, string(os.PathSeparator), Home)
)

func AppHome() {
	liquigenHome := Dir
	if ok := filez.DirExists(liquigenHome); !ok {
		if err := os.MkdirAll(liquigenHome, os.ModePerm); err != nil {
			panic(fmt.Sprintf("mkdir liquigen home dir:%s error:%v", liquigenHome, err))
		}
	}

	if filez.FileNotExists(liquigenHome, LiquigenConfigFile) {
		buf := bytes.NewBufferString(liquigenConfigJsonContent)
		tomatoConfigFile := filepath.Join(liquigenHome, strings.ToLower(LiquigenConfigFile))
		if err := os.WriteFile(tomatoConfigFile, buf.Bytes(), 0o644); err != nil {
			panic(fmt.Sprintf("writing file %s: %v", tomatoConfigFile, err))
		}
	}
}
