package scanner

import "github.com/moroz/go-lox/token"

var keywords map[string]token.TokenType = map[string]token.TokenType{
	"and":    token.TokenType_And,
	"class":  token.TokenType_Class,
	"else":   token.TokenType_Else,
	"false":  token.TokenType_False,
	"for":    token.TokenType_For,
	"fun":    token.TokenType_Fun,
	"if":     token.TokenType_If,
	"nil":    token.TokenType_Nil,
	"or":     token.TokenType_Or,
	"print":  token.TokenType_Print,
	"return": token.TokenType_Return,
	"super":  token.TokenType_Super,
	"this":   token.TokenType_This,
	"true":   token.TokenType_True,
	"var":    token.TokenType_Var,
	"while":  token.TokenType_While,
}
