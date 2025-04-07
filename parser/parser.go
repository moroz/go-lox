package parser

import (
	"github.com/moroz/go-lox/expr"
	"github.com/moroz/go-lox/token"
)

type Parser struct {
	vm      vm
	tokens  []token.Token
	current int
}

type vm interface {
	ReportParseError(err ParseError)
}

type ParseError struct {
	Token   token.Token
	Message string
}

func (e ParseError) Error() string {
	return e.Message
}

func NewParser(vm vm, tokens []token.Token) *Parser {
	return &Parser{vm: vm, tokens: tokens}
}

func (p *Parser) expression() (expr.Expr[any], error) {
	return p.equality()
}

func (p *Parser) equality() (expr.Expr[any], error) {
	ex, err := p.comparison()
	if err != nil {
		return ex, err
	}

	for p.match(token.TokenType_BangEqual, token.TokenType_EqualEqual) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return right, err
		}
		ex = expr.Binary[any]{
			Left:     ex,
			Operator: operator,
			Right:    right,
		}
	}

	return ex, nil
}

func (p *Parser) comparison() (expr.Expr[any], error) {
	ex, err := p.term()
	if err != nil {
		return ex, err
	}

	for p.match(token.TokenType_Greater, token.TokenType_GreaterEqual, token.TokenType_Less, token.TokenType_LessEqual) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return right, err
		}
		ex = expr.Binary[any]{
			Left:     ex,
			Operator: operator,
			Right:    right,
		}
	}

	return ex, nil
}

func (p *Parser) term() (expr.Expr[any], error) {
	ex, err := p.factor()
	if err != nil {
		return ex, err
	}

	for p.match(token.TokenType_Minus, token.TokenType_Plus) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return right, err
		}
		ex = expr.Binary[any]{
			Left:     ex,
			Operator: operator,
			Right:    right,
		}
	}

	return ex, nil
}

func (p *Parser) factor() (expr.Expr[any], error) {
	ex, err := p.unary()
	if err != nil {
		return ex, err
	}

	for p.match(token.TokenType_Slash, token.TokenType_Star) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return right, err
		}
		ex = expr.Binary[any]{
			Left:     ex,
			Operator: operator,
			Right:    right,
		}
	}

	return ex, nil
}

func (p *Parser) unary() (expr.Expr[any], error) {
	if p.match(token.TokenType_Bang, token.TokenType_Minus) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return right, err
		}
		return expr.Unary[any]{
			Operator: operator,
			Right:    right,
		}, nil
	}

	return p.primary()
}

func (p *Parser) primary() (expr.Expr[any], error) {
	if p.match(token.TokenType_False) {
		return expr.Literal[any]{
			Value: false,
		}, nil
	}

	if p.match(token.TokenType_True) {
		return expr.Literal[any]{
			Value: true,
		}, nil
	}

	if p.match(token.TokenType_Nil) {
		return expr.Literal[any]{
			Value: nil,
		}, nil
	}

	if p.match(token.TokenType_Number, token.TokenType_String) {
		return expr.Literal[any]{
			Value: p.previous().Literal,
		}, nil
	}

	if p.match(token.TokenType_LeftParen) {
		ex, err := p.expression()
		if err != nil {
			return ex, err
		}
		_, err = p.consume(token.TokenType_RightParen, "Expect ')' after expression.")
		if err != nil {
			return nil, err
		}
		return expr.Grouping[any]{Expression: ex}, nil
	}

	panic("unexpected branch")
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) consume(t token.TokenType, message string) (*token.Token, error) {
	if p.check(t) {
		next := p.advance()
		return &next, nil
	}

	return nil, ParseError{
		Token:   p.peek(),
		Message: message,
	}
}

func (p *Parser) check(t token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().TokenType == t
}

func (p *Parser) advance() token.Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == token.TokenType_EOF
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}
