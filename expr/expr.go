package expr

import "github.com/moroz/go-lox/token"

type Expr[R any] interface {
	Accept(Visitor[R]) R
}

type Visitor[R any] interface {
	VisitBinaryExpr(Binary[R]) R
	VisitGroupingExpr(Grouping[R]) R
	VisitLiteralExpr(Literal[R]) R
	VisitUnaryExpr(Unary[R]) R
}

type Binary[R any] struct {
	Left     Expr[R]
	Operator token.Token
	Right    Expr[R]
}

type Grouping[R any] struct {
	Expression Expr[R]
}

type Literal[R any] struct {
	Value any
}

type Unary[R any] struct {
	Operator token.Token
	Right    Expr[R]
}

func (e Unary[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitUnaryExpr(e)
}

func (e Literal[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitLiteralExpr(e)
}

func (e Grouping[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitGroupingExpr(e)
}

func (e Binary[R]) Accept(visitor Visitor[R]) R {
	return visitor.VisitBinaryExpr(e)
}
