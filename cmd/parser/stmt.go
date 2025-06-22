package parser

import "go_tut/TeaScript/cmd/ast"

func parseStmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	expr := parse_expr(p, def_bp)

	return ast.ExprStmt{
		Expression: expr,
	}
}
