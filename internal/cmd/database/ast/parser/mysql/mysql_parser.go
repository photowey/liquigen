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

package mysql

import (
	"errors"
	"strings"

	"github.com/photowey/liquigen/internal/cmd/database"
	"github.com/photowey/liquigen/internal/cmd/database/ast"
	"github.com/photowey/liquigen/internal/cmd/database/ast/lexer"
	"github.com/photowey/liquigen/internal/cmd/database/ast/parser"
	"github.com/photowey/liquigen/pkg/stringz"
)

// ----------------------------------------------------------------

const (
	Dialect = "mysql"

	CreateTableStatement = "CREATE TABLE"
	DropTableStatement   = "DROP TABLE"
	AlterTableStatement  = "ALTER TABLE"
)

// ----------------------------------------------------------------

var _ parser.Parser = (*Parser)(nil)

// ----------------------------------------------------------------

func init() {
	parser.Register(NewParser())
}

// ----------------------------------------------------------------

type Parser struct {
	dialect string
}

// ----------------------------------------------------------------

func NewParser() parser.Parser {
	return &Parser{
		dialect: Dialect,
	}
}

// ----------------------------------------------------------------

func (p Parser) Dialect() string {
	return p.dialect
}

func (p Parser) Parse(sql string) (*ast.Ast, error) {
	return parse(sql)
}

func parse(sql string) (*ast.Ast, error) {
	db, statements, err := parseSQL(sql)
	if err != nil {
		return nil, err
	}

	return &ast.Ast{
		SQL:        sql,
		Statements: statements,
		Database:   db,
	}, nil
}

func predicateIsDropTableStatement(statement string) bool {
	return strings.HasPrefix(strings.ToUpper(statement), DropTableStatement)
}

func predicateIsCreateTableStatement(statement string) bool {
	return strings.HasPrefix(strings.ToUpper(statement), CreateTableStatement)
}

func predicateIsAlterTableStatement(statement string) bool {
	return strings.HasPrefix(strings.ToUpper(statement), AlterTableStatement)
}

func parseSQL(sql string) (*ast.Database, []string, error) {
	sql = database.RemoveComments(sql)
	statements := database.SplitSQLStatements(sql)

	var tables []*ast.Table

	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if stringz.IsBlankString(statement) {
			continue
		}

		// DROP TABLE ...
		if predicateIsDropTableStatement(statement) {
			continue
		}

		// ALTER TABLE ...
		if !predicateIsAlterTableStatement(statement) &&
			// CREATE TABLE ...
			!predicateIsCreateTableStatement(statement) {
			return nil, nil, errors.New("bad create table SQL statements")
		}

		tokenizer, err := tryTokenize(statement)
		if err != nil {
			return nil, statements, err
		}

		table, err := tryParse(tokenizer)
		if err != nil {
			return nil, statements, err
		}

		if table.AlterStatement {
			for _, tb := range tables {
				if tb.Name == table.Name {
					tb.Comment = stringz.RemoveQuotes(table.Comment)
				}
			}

			continue
		}

		tables = append(tables, table)
	}

	databaseName := "Unknown"
	if len(tables) > 0 {
		if stringz.IsNotBlankString(tables[0].Database) {
			databaseName = stringz.RemoveQuotes(tables[0].Database)
		}
	}

	return &ast.Database{
		Name:   databaseName,
		Tables: tables,
	}, statements, nil
}

func tryTokenize(statement string) (*ast.Tokenizer, error) {
	tokens, err := tokenize(statement)
	if err != nil {
		return nil, err
	}

	return &ast.Tokenizer{Tokens: tokens}, nil
}

func tokenize(sql string) ([]ast.Token, error) {
	var tokens []ast.Token
	sql = strings.TrimSpace(sql)
	words := stringz.Fields(sql)

	for _, word := range words {
		switch strings.ToUpper(word) {
		case lexer.TokenKeywordCreate:
			tokens = append(tokens, ast.Token{Type: lexer.TokenCreate, Literal: word})
		case lexer.TokenKeywordTable:
			tokens = append(tokens, ast.Token{Type: lexer.TokenTable, Literal: word})

		case lexer.TokenKeywordIf:
			tokens = append(tokens, ast.Token{Type: lexer.TokenIf, Literal: word})
		case lexer.TokenKeywordNot:
			tokens = append(tokens, ast.Token{Type: lexer.TokenNot, Literal: word})
		case lexer.TokenKeywordExists:
			tokens = append(tokens, ast.Token{Type: lexer.TokenExists, Literal: word})

		case lexer.TokenKeywordPrimary:
			tokens = append(tokens, ast.Token{Type: lexer.TokenPrimary, Literal: word})
		case lexer.TokenKeywordUnique:
			tokens = append(tokens, ast.Token{Type: lexer.TokenUnique, Literal: word})
		case lexer.TokenKeywordForeign:
			tokens = append(tokens, ast.Token{Type: lexer.TokenForeign, Literal: word})
		case lexer.TokenKeywordReferences:
			tokens = append(tokens, ast.Token{Type: lexer.TokenReferences, Literal: word})
		case lexer.TokenKeywordKey:
			tokens = append(tokens, ast.Token{Type: lexer.TokenKey, Literal: word})
		case lexer.TokenKeywordUnsigned:
			tokens = append(tokens, ast.Token{Type: lexer.TokenUnsigned, Literal: word})
		case lexer.TokenKeywordZerofill:
			tokens = append(tokens, ast.Token{Type: lexer.TokenZerofill, Literal: word})

		case lexer.TokenKeywordNull:
			tokens = append(tokens, ast.Token{Type: lexer.TokenNull, Literal: word})
		case lexer.TokenKeywordAutoIncrement:
			tokens = append(tokens, ast.Token{Type: lexer.TokenAutoIncrement, Literal: word})
		case lexer.TokenKeywordDefault:
			tokens = append(tokens, ast.Token{Type: lexer.TokenDefault, Literal: word})

		case lexer.TokenKeywordOn:
			tokens = append(tokens, ast.Token{Type: lexer.TokenOn, Literal: word})
		case lexer.TokenKeywordUpdate:
			tokens = append(tokens, ast.Token{Type: lexer.TokenUpdate, Literal: word})
		case lexer.TokenKeywordCurrentTimestamp:
			tokens = append(tokens, ast.Token{Type: lexer.TokenCurrentTimestamp, Literal: word})

		case lexer.TokenKeywordCharset:
			tokens = append(tokens, ast.Token{Type: lexer.TokenCharset, Literal: word})
		case lexer.TokenKeywordCharacter:
			tokens = append(tokens, ast.Token{Type: lexer.TokenCharacter, Literal: word})

		case lexer.TokenKeywordCollate:
			tokens = append(tokens, ast.Token{Type: lexer.TokenCollate, Literal: word})
		case lexer.TokenKeywordComment:
			tokens = append(tokens, ast.Token{Type: lexer.TokenComment, Literal: word})

		case lexer.TokenKeywordPartition:
			tokens = append(tokens, ast.Token{Type: lexer.TokenPartition, Literal: word})
		case lexer.TokenKeywordBy:
			tokens = append(tokens, ast.Token{Type: lexer.TokenBy, Literal: word})

		case lexer.TokenKeywordIndex:
			tokens = append(tokens, ast.Token{Type: lexer.TokenIndex, Literal: word})

		case lexer.TokenKeywordConstraint:
			tokens = append(tokens, ast.Token{Type: lexer.TokenConstraint, Literal: word})
		case lexer.TokenKeywordSet:
			tokens = append(tokens, ast.Token{Type: lexer.TokenSet, Literal: word})

		case lexer.TokenKeywordCheck:
			tokens = append(tokens, ast.Token{Type: lexer.TokenCheck, Literal: word})

		case lexer.TokenKeywordEngine:
			tokens = append(tokens, ast.Token{Type: lexer.TokenEngine, Literal: word})

		case lexer.TokenKeywordComma:
			tokens = append(tokens, ast.Token{Type: lexer.TokenComma, Literal: word})
		case lexer.TokenKeywordLeftParen:
			tokens = append(tokens, ast.Token{Type: lexer.TokenLeftParen, Literal: word})
		case lexer.TokenKeywordRightParen:
			tokens = append(tokens, ast.Token{Type: lexer.TokenRightParen, Literal: word})
		case lexer.TokenKeyAlter:
			tokens = append(tokens, ast.Token{Type: lexer.TokenAlter, Literal: word})
		default:
			if TestIsMySQLDataType(word) {
				tokens = append(tokens, ast.Token{Type: lexer.TokenDataType, Literal: word})
			} else {
				tokens = append(tokens, ast.Token{Type: lexer.TokenIdentifier, Literal: word})
			}
		}
	}

	tokens = append(tokens, ast.Token{Type: lexer.TokenEOF})

	return tokens, nil
}

func tryParse(tokenizer *ast.Tokenizer) (*ast.Table, error) {
	table := &ast.Table{
		CreateStatement: true,
		AlterStatement:  false,
	}

	// ALTER TABLE ...
	// ALTER TABLE table_name COMMENT = 'comment';
	if tokenizer.Peek().Type == lexer.TokenAlter {
		tokenizer.Next() // ALTER
		if tokenizer.HasNext() {
			if tokenizer.Peek().Type == lexer.TokenTable {
				tokenizer.Next() // TABLE

				tab := tokenizer.Next() // table_name
				table.Name = tab.Literal

				tokenizer.Next() // COMMENT

				cmt := tokenizer.Next()
				table.Comment = stringz.RemoveQuotes(cmt.Literal)

				table.AlterStatement = true
				table.CreateStatement = false

				return table, nil
			}
		}
	}

	// Maybe comments ...
	// CREATE TABLE ...
	for tokenizer.Peek().Type != lexer.TokenCreate {
		tokenizer.Next()
		if tokenizer.HasNotNext() {
			break
		}
	}

	tokenizer.Next() // CREATE

OUTER:
	for tokenizer.Peek().Type != lexer.TokenTable {
		for tokenizer.Peek().Type != lexer.TokenCreate {
			tokenizer.Next()

			if tokenizer.HasNotNext() {
				break OUTER
			}
		}

		tokenizer.Next() // CREATE
	}

	if tokenizer.HasNotNext() {
		return nil, errors.New("bad SQL statements")
	}

	tokenizer.Next() //  TABLE

	// IF NOT EXISTS
	if tokenizer.Peek().Type == lexer.TokenIf {
		tokenizer.Next() // IF
		tokenizer.Next() // NOT
		tokenizer.Next() // EXISTS
	}

	// Table name
	tokenTableName := tokenizer.Next()
	table.Name = tokenTableName.Literal
	if strings.Contains(table.Name, ".") {
		pair := strings.Split(table.Name, ".")
		table.Database = pair[0]
		table.Name = pair[1]
	}

	// (
	tokenizer.Next() // '('

	for tokenizer.HasNext() {
		if tokenizer.Peek().Type == lexer.TokenRightParen {
			break
		}

		column := &ast.Column{
			Name:     tokenizer.Next().Literal, // Name
			DataType: tokenizer.Next().Literal, // Data type
		}

		// Length | Precision | Scale
		if tokenizer.Peek().Type == lexer.TokenLeftParen {
			tokenizer.Next()                      // '('
			length := ast.ToInt(tokenizer.Next()) // Length
			column.Length = length

			if tokenizer.Peek().Type == lexer.TokenComma {
				tokenizer.Next()                         // ','
				precision := ast.ToInt(tokenizer.Next()) // Precision
				column.Precision = precision

				if tokenizer.Peek().Type == lexer.TokenComma {
					tokenizer.Next()                     // 跳过 ','
					scale := ast.ToInt(tokenizer.Next()) // Scale
					column.Scale = scale
				}
			}
			tokenizer.Next() // ')'
		}

		// Other
		for tokenizer.HasNext() && tokenizer.Peek().Type != lexer.TokenComma && tokenizer.Peek().Type != lexer.TokenRightParen {
			switch tokenizer.Peek().Type {
			case lexer.TokenNot:
				tokenizer.Next() // 'NOT'
				if tokenizer.Peek().Type == lexer.TokenNull {
					tokenizer.Next()      // 'NULL'
					column.NotNull = true // NOT NULL
				}
			case lexer.TokenPrimary:
				tokenizer.Next() // 'PRIMARY'
				if tokenizer.Peek().Type == lexer.TokenKey {
					tokenizer.Next() // 'KEY'
					column.PrimaryKey = true
					column.NotNull = true
				}
			case lexer.TokenUnique:
				tokenizer.Next() // 'UNIQUE'
				if tokenizer.Peek().Type == lexer.TokenKey {
					tokenizer.Next()
					column.Unique = true
				}
			case lexer.TokenAutoIncrement:
				tokenizer.Next() // 'AUTO_INCREMENT'
				column.AutoIncrement = true
			case lexer.TokenDefault:
				tokenizer.Next() // 'DEFAULT'
				defaultValueToken := tokenizer.Next()
				column.Default = stringz.RemoveQuotes(defaultValueToken.Literal) // ' | "
			case lexer.TokenOn:
				tokenizer.Next() // 'ON'
				if tokenizer.Peek().Type == lexer.TokenUpdate {
					tokenizer.Next() // 'UPDATE'
					if tokenizer.Peek().Type == lexer.TokenCurrentTimestamp {
						DefaultToken := tokenizer.Next()
						column.Default = DefaultToken.Literal
						column.UpdateTimestamp = true
					}
				}
			case lexer.TokenComment:
				tokenizer.Next()                                                // 'COMMENT'
				column.Comment = stringz.RemoveQuotes(tokenizer.Next().Literal) // Comment
			default:
				tokenizer.Next()
			}
		}

		table.Columns = append(table.Columns, column)

		if tokenizer.Peek().Type == lexer.TokenComma {
			tokenizer.Next() // ','
		}
	}

	// )
	tokenizer.Next() // ')'

	for tokenizer.HasNext() {
		// Table comment
		if tokenizer.Peek().Type == lexer.TokenComment {
			tokenizer.Next()                                               // 'COMMENT'
			tokenizer.Next()                                               // '='
			table.Comment = stringz.RemoveQuotes(tokenizer.Next().Literal) // Comment
		}

		// Other k/v.

		if tokenizer.Peek().Type == lexer.TokenSemicolon {
			tokenizer.Next() // ';'
			break
		}

		tokenizer.Next()
	}

	return table, nil
}
