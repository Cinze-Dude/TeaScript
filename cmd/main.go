package main

import (
	"fmt"
	"go_tut/TeaScript/cmd/lexer"
	"go_tut/TeaScript/cmd/parser"
)

func main() {
	var src string = `10 * 2`

	tokens := lexer.Tokenize(string(src))

	ast := parser.Parse(tokens)

	fmt.Println(ast)
}
