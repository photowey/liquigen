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

package configs

import (
	"fmt"
	"os"

	"github.com/photowey/liquigen/pkg/jsonz"
)

var _config Config

type Config struct {
	Project  Project  `toml:"project" json:"project" yaml:"project"`
	Database Database `toml:"database" json:"database" yaml:"database"`
}

type Project struct {
	Author  string `toml:"author" json:"author" yaml:"author"`
	Email   string `toml:"email" json:"email"`
	Version string `toml:"version" json:"version" yaml:"version"`
	Dialect string `toml:"dialect" json:"dialect" yaml:"dialect"`
	SQL     string `toml:"sql" json:"sql" yaml:"sql"`
}

type Database struct {
	Host     string   `toml:"host" json:"host" yaml:"host"`
	Port     int      `toml:"port" json:"port" yaml:"port"`
	Dialect  string   `toml:"dialect" json:"dialect" yaml:"dialect"`
	Driver   string   `toml:"driver" json:"driver" yaml:"driver"`
	Database string   `toml:"database" json:"database" yaml:"database"`
	Username string   `toml:"username" json:"username" yaml:"username"`
	Password string   `toml:"password" json:"password" yaml:"password"`
	Includes []string `toml:"includes" json:"includes" yaml:"includes"`
	Excludes []string `toml:"excludes" json:"excludes" yaml:"excludes"`
	Prefixes []string `toml:"prefixes" json:"prefixes" yaml:"prefixes"`
}

func Init(configFile string) {
	conf, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("parse the liquigen config error:%s", err.Error())
		return
	}

	jsonz.UnmarshalStruct(conf, &_config)
}

func Configs() Config {
	return _config
}

func ConfigProject() Project {
	return _config.Project
}

func ConfigDatabase() Database {
	return _config.Database
}
