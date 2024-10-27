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

package ast

import (
	"github.com/photowey/liquigen/internal/cmd/database/ast/lexer"
)

type Token struct {
	Type    lexer.TokenType
	Literal string
}

type Tokenizer struct {
	Tokens   []Token
	position int
}

func NewTokenizer(tokens []Token) *Tokenizer {
	return &Tokenizer{
		Tokens:   tokens,
		position: 0,
	}
}

func (t *Tokenizer) Next() Token {
	if t.position >= len(t.Tokens) {
		return Token{Type: lexer.TokenEOF} // 处理结束标记
	}
	token := t.Tokens[t.position]
	t.position++

	return token
}

func (t *Tokenizer) Peek() Token {
	if t.position >= len(t.Tokens) {
		return Token{Type: lexer.TokenEOF} // 处理结束标记
	}
	return t.Tokens[t.position]
}

func (t *Tokenizer) HasNext() bool {
	return t.position < len(t.Tokens)
}

func (t *Tokenizer) HasNotNext() bool {
	return !t.HasNext()
}
