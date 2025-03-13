package main

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
	Operator Token
	Right    Expr[R]
}

type Grouping[R any] struct {
	Expression Expr[R]
}

type Literal[R any] struct {
	Value any
}

type Unary[R any] struct {
	Operator Token
	Right    Expr[R]
}
