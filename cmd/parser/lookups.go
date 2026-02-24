package parser

import (
	"github.com/omar/TeaScript/cmd/ast"
	"github.com/omar/TeaScript/cmd/lexer"
)

type binding_power int

const (
	def_bp binding_power = iota
	comma
	assign
	logical
	relational
	additive
	multiplicative
	exponantiative
	unary
	call
	member
	primary
)

type stmtHandler func(p *parser) ast.Stmt
type nudHandler func(p *parser) ast.Expr
type ledHandler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmtHandler
type nud_lookup map[lexer.TokenKind]nudHandler
type led_lookup map[lexer.TokenKind]ledHandler
type bp_lookup map[lexer.TokenKind]binding_power

var stmt_lu = stmt_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var bp_lu = bp_lookup{}

func led(kind lexer.TokenKind, bp binding_power, led_fn ledHandler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, bp binding_power, nud_fn nudHandler) {
	bp_lu[kind] = bp
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmtHandler) {
	bp_lu[kind] = def_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {
	// Logical operators
	led(lexer.AND, logical, parseBinExpr)
	led(lexer.OR, logical, parseBinExpr)
	led(lexer.XOR, logical, parseBinExpr)
	led(lexer.DBL_DOT, logical, parseBinExpr)
	led(lexer.SHIFT_LEFT, logical, parseBinExpr)
	led(lexer.SHIFT_RIGHT, logical, parseBinExpr)

	// Relational operators
	led(lexer.LESS, relational, parseBinExpr)
	led(lexer.MORE, relational, parseBinExpr)
	led(lexer.LESS_EQL, relational, parseBinExpr)
	led(lexer.MORE_EQL, relational, parseBinExpr)
	led(lexer.EQL, relational, parseBinExpr)
	led(lexer.NOT_EQL, relational, parseBinExpr)

	// Additive / Multiplicative / Exponential
	led(lexer.PLUS, additive, parseBinExpr)
	led(lexer.MINS, additive, parseBinExpr)
	led(lexer.STAR, multiplicative, parseBinExpr)
	led(lexer.SLSH, multiplicative, parseBinExpr)
	led(lexer.DBL_SLSH, multiplicative, parseBinExpr)
	led(lexer.MODL, multiplicative, parseBinExpr)
	led(lexer.POWR, exponantiative, parseBinExpr)

	// Literals & symbols
	nud(lexer.NUMBER, primary, parsePrimaryExpr)
	nud(lexer.STRING, primary, parsePrimaryExpr)
	nud(lexer.RUNE, primary, parsePrimaryExpr)
	nud(lexer.IDENTIFIER, primary, parsePrimaryExpr)

	// Unary operators
	nud(lexer.PLUS, unary, parseUnaryExpr)
	nud(lexer.MINS, unary, parseUnaryExpr)
	nud(lexer.NOT, unary, parseUnaryExpr)
	nud(lexer.SQRT, unary, parseUnaryExpr)

	// TODO: Add statement registrations like var, if, while here if needed
}
