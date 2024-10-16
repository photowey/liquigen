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

package jsonz

//
// json codec 包
//

import (
	"encoding/json"
	"io"
)

// ---------------------------------------------------------------- Encode

// String 将 Go 结构体转换为 json 对象-字符串
func String(body any) string {
	data, _ := StringE(body)

	return data
}

func StringE(body any) (string, error) {
	data, err := json.Marshal(body)

	return string(data), err
}

// Bytes 将 Go 结构体转换为 json 对象-字节数组
func Bytes(body any) []byte {
	data, _ := BytesE(body)

	return data
}

func BytesE(body any) ([]byte, error) {
	data, err := json.Marshal(body)

	return data, err
}

// Pretty 将 Go 结构体转换为 json 对象,
//
// 并采用 \t 缩进格式化
func Pretty(body any) string {
	data, _ := PrettyE(body)

	return data
}

func PrettyE(body any) (string, error) {
	bytes, err := json.MarshalIndent(body, "", "\t")

	return string(bytes), err
}

// ---------------------------------------------------------------- Decode

// DecodeStruct 采用 json.NewDecoder
//
// 从 io.Reader 读取数据并将 json byte 流转换为 Go 结构体
func DecodeStruct(reader io.Reader, target any) error {
	if err := json.NewDecoder(reader).Decode(target); err != nil {
		return err
	}

	return nil
}

// UnmarshalStructE 采用 json.Unmarshal
//
// 将 json byte 数据转换为 Go 结构体
//
// 返回错误
func UnmarshalStructE(data []byte, structy any) error {
	if err := json.Unmarshal(data, structy); err != nil {
		return err
	}

	return nil
}

// UnmarshalStruct 将 json 转换为结构体数据,且忽略错误
//
// 不返回错误
//
// 使用时需要谨慎,除非明确不会抛错
func UnmarshalStruct(data []byte, structy any) {
	_ = json.Unmarshal(data, structy)
}
