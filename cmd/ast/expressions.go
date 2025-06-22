package ast

import "go_tut/TeaScript/cmd/lexer"

// ---------------------
// LITERAL EXPRESSIONS
// ---------------------

// Number Expression
type NumExpr struct {
	Value float64
}

func (n NumExpr) expr() {}

// String Expression
type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

// Rune Expression
type RuneExpr struct {
	Value rune
}

func (n RuneExpr) expr() {}

// Symbol Expression
type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// ---------------------
// COMPLEX EXPRESSIONS
// ---------------------

type BinExpr struct {
	Operator lexer.Token
	Left     Expr
	Right    Expr
}

func (n BinExpr) expr() {}
