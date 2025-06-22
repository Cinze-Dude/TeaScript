package parser

import (
	"go_tut/TeaScript/cmd/ast"
	"go_tut/TeaScript/cmd/lexer"
)

func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	return &parser{
		tokens: tokens,
		pos:    0,
	}
}

type parser struct {
	tokens []lexer.Token
	pos    int
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)
	p := createParser(tokens)

	for p.hasTokens() {
		Body = append(Body, parseStmt(p))
	}

	return ast.BlockStmt{
		Body: Body,
	}
}

// Helper Methods
func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.tokens[p.pos].Kind
}

func (p *parser) adv() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}
