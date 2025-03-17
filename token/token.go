package token

import "fmt"

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   any
	Line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{tokenType, lexeme, literal, line}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.TokenType, t.Lexeme, t.Literal)
}
