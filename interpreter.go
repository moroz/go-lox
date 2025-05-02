package golox

import (
	"github.com/moroz/go-lox/expr"
	"github.com/moroz/go-lox/token"
)

type Interpreter struct{}

func (i *Interpreter) Evaluate(e expr.Expr) any {
	switch e.(type) {
	case (*expr.Literal):
		return e.(*expr.Literal).Value

	case (*expr.Grouping):
		return i.Evaluate(e.(*expr.Grouping).Expression)

	case (*expr.Unary):
		e := e.(*expr.Unary)
		right := i.Evaluate(e.Right)

		switch e.Operator.TokenType {
		case token.TokenType_Bang:
			return !isTruthy(right)
		case token.TokenType_Minus:
			return -(right.(float64))
		}
		return nil
	}

	return nil
}

func isTruthy(val any) bool {
	return val == nil || val == false
}
