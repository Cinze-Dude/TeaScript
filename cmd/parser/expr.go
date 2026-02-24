package parser

import (
	"fmt"
	"strconv"

	"github.com/omar/TeaScript/cmd/ast"
	"github.com/omar/TeaScript/cmd/lexer"
)

// ----------------------
// Expression parsing
// ----------------------

func parse_expr(p *parser, minBP binding_power) ast.Expr {
	token := p.currentToken()
	nud_fn := nud_lu[token.Kind]
	left := nud_fn(p)

	for {
		op := p.currentToken()
		opBP := bp_lu[op.Kind]
		if opBP <= minBP {
			break
		}

		led_fn := led_lu[op.Kind]

		rightBP := opBP
		// Right-associative operators like exponentiation
		if op.Kind == lexer.POWR {
			rightBP--
			if rightBP < 0 {
				rightBP = 0
			}
		}

		left = led_fn(p, left, rightBP)
	}

	return left
}

// ----------------------
// Primary expressions
// ----------------------

func parsePrimaryExpr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		num, _ := strconv.ParseFloat(p.adv().Value, 64)
		return ast.NumExpr{Value: num}
	case lexer.STRING:
		return ast.StringExpr{Value: p.adv().Value}
	case lexer.RUNE:
		val := []rune(p.adv().Value)
		return ast.RuneExpr{Value: val[0]}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{Value: p.adv().Value}
	default:
		kind, _ := lexer.TokenKindString(p.currentTokenKind())
		panic(fmt.Sprintf("Parsing Error: cannot parse a primary expression from %s", lexer.ColorText(kind, "yellow")))
	}
}

// ----------------------
// Binary expressions
// ----------------------

func parseBinExpr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	opToken := p.adv()
	right := parse_expr(p, bp)
	return ast.BinExpr{
		Operator: opToken,
		Left:     left,
		Right:    right,
	}
}

// ----------------------
// Unary expressions
// ----------------------

func parseUnaryExpr(p *parser) ast.Expr {
	opToken := p.adv()
	operand := parse_expr(p, unary)
	return ast.BinExpr{
		Operator: opToken,
		Left:     nil,     // Unary operators have no left
		Right:    operand, // Store operand in Right for simplicity
	}
}
