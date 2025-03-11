//go:generate stringer -type=TokenType -trimprefix=TokenType_

package main

type TokenType int

const (
	// Single-character tokens.
	TokenType_LeftParen TokenType = iota
	TokenType_RightParen
	TokenType_LeftBrace
	TokenType_RightBrace
	TokenType_Comma
	TokenType_Dot
	TokenType_Minus
	TokenType_Plus
	TokenType_Semicolon
	TokenType_Slash
	TokenType_Star

	// One or two character tokens.
	TokenType_Bang
	TokenType_BangEqual
	TokenType_Equal
	TokenType_EqualEqual
	TokenType_Greater
	TokenType_GreaterEqual
	TokenType_Less
	TokenType_LessEqual

	// Literals.
	TokenType_Identifier
	TokenType_String
	TokenType_Number

	// Keywords.
	TokenType_And
	TokenType_Class
	TokenType_Else
	TokenType_False
	TokenType_Fun
	TokenType_For
	TokenType_If
	TokenType_Nil
	TokenType_Or
	TokenType_Print
	TokenType_Return
	TokenType_Super
	TokenType_This
	TokenType_True
	TokenType_Var
	TokenType_While

	TokenType_EOF
)

var keywords map[string]TokenType = map[string]TokenType{
	"and":    TokenType_And,
	"class":  TokenType_Class,
	"else":   TokenType_Else,
	"false":  TokenType_False,
	"for":    TokenType_For,
	"fun":    TokenType_Fun,
	"if":     TokenType_If,
	"nil":    TokenType_Nil,
	"or":     TokenType_Or,
	"print":  TokenType_Print,
	"return": TokenType_Return,
	"super":  TokenType_Super,
	"this":   TokenType_This,
	"true":   TokenType_True,
	"var":    TokenType_Var,
	"while":  TokenType_While,
}
