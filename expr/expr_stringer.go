package expr

import (
	"fmt"
	"strings"
)

func (e Binary) String() string {
	return parenthesize(e.Operator.Lexeme, e.Left, e.Right)
}

func (e Grouping) String() string {
	return parenthesize("group", e.Expression)
}

func (e Literal) String() string {
	if e.Value == nil {
		return "nil"
	}

	switch e.Value.(type) {
	case string:
		return fmt.Sprintf("%q", e.Value)

	default:
		return fmt.Sprintf("%v", e.Value)
	}
}

func (e Unary) String() string {
	return parenthesize(e.Operator.Lexeme, e.Right)
}

func parenthesize(name string, exprs ...Expr) string {
	var builder strings.Builder

	builder.WriteRune('(')
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteRune(' ')
		builder.WriteString(expr.String())
	}
	builder.WriteRune(')')
	return builder.String()
}
