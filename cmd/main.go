package main

import (
	"fmt"
	"github.com/omar/TeaScript/cmd/lexer"
	"github.com/omar/TeaScript/cmd/parser"
)

func main() {
	var src string = `43.56 + 6 ^ 2`

	tokens := lexer.Tokenize(string(src))

	ast := String(parser.Parse(tokens))

	fmt.Println(ast)
}
