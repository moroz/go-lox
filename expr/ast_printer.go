package expr

import (
	"fmt"
	"strconv"
	"strings"
)

type AstPrinter struct{}

func (p *AstPrinter) Print(expr Expr[string]) string {
	return expr.Accept(p)
}

func (p *AstPrinter) VisitBinaryExpr(expr Binary[string]) string {
	return p.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (p *AstPrinter) VisitGroupingExpr(expr Grouping[string]) string {
	return p.parenthesize("group", expr.Expression)
}

func (p *AstPrinter) VisitLiteralExpr(expr Literal[string]) string {
	if expr.Value == nil {
		return "nil"
	}

	switch expr.Value.(type) {
	case float64:
		return strconv.FormatFloat(expr.Value.(float64), 'g', -1, 64)

	case string:
		return "\"" + expr.Value.(string) + "\""

	default:
		return fmt.Sprintf("%v", expr.Value)
	}
}

func (p *AstPrinter) VisitUnaryExpr(expr Unary[string]) string {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (p *AstPrinter) parenthesize(name string, exprs ...Expr[string]) string {
	var builder strings.Builder

	builder.WriteRune('(')
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteRune(' ')
		builder.WriteString(expr.Accept(p))
	}
	builder.WriteRune(')')
	return builder.String()
}
