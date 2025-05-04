package expr

import "github.com/moroz/go-lox/token"

type Expr interface {
	String() string
}

type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

type Grouping struct {
	Expression Expr
}

type Literal struct {
	Value any
}

type Unary struct {
	Operator *token.Token
	Right    Expr
}
