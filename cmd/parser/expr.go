package parser

import (
	"fmt"
	"github.com/omar/TeaScript/cmd/ast"
	"github.com/omar/TeaScript/cmd/lexer"
	"strconv"
)

func parse_expr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	str, _ := lexer.TokenKindString(tokenKind)
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("Parsing Error: NUD handler expected for token: %s\n", str))
	}

	left := nud_fn(p)
	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("Parsing Error: NUD handler expected for token: %s\n", str))
		}

		left = led_fn(p, left, bp)
	}

	return left
}

func parsePrimaryExpr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		num, _ := strconv.ParseFloat(p.adv().Value, 64)
		return ast.NumExpr{
			Value: num,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.adv().Value,
		}
	case lexer.RUNE:
		return ast.RuneExpr{
			Value: rune(p.adv().Value[0]),
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.adv().Value,
		}
	default:
		kind, _ := lexer.TokenKindString(p.currentTokenKind())
		panic(fmt.Sprintf("Parsing Error: cannot parse a primary exression from %s\n", lexer.ColorText(kind, "yellow")))
	}
}

func parseBinExpr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	opToken := p.adv()
	rightBP := bp
	if opToken.Kind == lexer.POWR { // right-associative
		rightBP = bp - 1
	}
	right := parse_expr(p, rightBP)

	return ast.BinExpr{
		Operator: opToken,
		Left:     left,
		Right:    right,
	}
}

