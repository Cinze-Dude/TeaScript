package main

import (
	//"os"
	"github.com/sanity-io/litter"
	"github.com/omar/TeaScript/cmd/lexer"
	"github.com/omar/TeaScript/cmd/parser"
)

func main() {
	//bytes, _ := os.ReadFile("./tests/test.tea")
	tokens :=  lexer.Tokenize("10.5 * 6 + 5 ^ 3" /* string(bytes) */)

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}
