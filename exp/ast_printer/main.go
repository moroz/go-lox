package main

import (
	"fmt"

	"github.com/moroz/go-lox/expr"
	"github.com/moroz/go-lox/token"
)

func main() {
	expression := expr.Binary[string]{
		Left: expr.Unary[string]{
			Operator: token.NewToken(token.TokenType_Minus, "-", nil, 1),
			Right:    expr.Literal[string]{Value: 123},
		},
		Operator: token.NewToken(token.TokenType_Star, "*", nil, 1),
		Right: expr.Grouping[string]{
			Expression: expr.Literal[string]{Value: 45.67},
		},
	}

	printer := expr.AstPrinter{}
	fmt.Println(printer.Print(expression))
}
