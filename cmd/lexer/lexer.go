package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	src      string
	pos      int
}

func (lex *lexer) advanceNum(num int) {
	lex.pos += num
}

func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

/*
// This is currently unused
func (lex *lexer) at() byte {
	return lex.src[lex.pos]
}
*/

func (lex *lexer) rem() string {
	return lex.src[lex.pos:]
}

func (lex *lexer) atEOF() bool {
	return lex.pos >= len(lex.src)
}

func Tokenize(src string) []Token {
	lex := CreateLexer(src)
	for !lex.atEOF() {
		matched := false
		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.rem())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("Lexing Error: unreachable pair detected near: %s\nline: %v", lex.rem(), lex.pos))
		}
	}

	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
}

func defHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advanceNum(len(value))
		lex.push(NewToken(kind, value))
	}
}

func CreateLexer(src string) *lexer {
	return &lexer{
		pos:    0,
		src:    src,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z_0-9]*`), smblHandler},
			{regexp.MustCompile(`<[^<]*>`), charHandler},
			{regexp.MustCompile(`"[^"]*"|'[^']*'`), strHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`<<`), defHandler(SHIFT_LEFT, "<<")},
			{regexp.MustCompile(`>>`), defHandler(SHIFT_RIGHT, ">>")},
			{regexp.MustCompile(`~\|`), defHandler(XOR, "~|")},
			{regexp.MustCompile(`\[`), defHandler(OPEN_BRACK, "[")},
			{regexp.MustCompile(`\]`), defHandler(CLOSE_BRACK, "]")},
			{regexp.MustCompile(`\{`), defHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`~=`), defHandler(NOT_EQL, "~=")},
			{regexp.MustCompile(`==`), defHandler(EQL, "==")},
			{regexp.MustCompile(`=>`), defHandler(ARROW, "=>")},
			{regexp.MustCompile(`=`), defHandler(ASSIGN, "=")},
			{regexp.MustCompile(`~`), defHandler(NOT, "~")},
			{regexp.MustCompile(`<=`), defHandler(LESS_EQL, "<=")},
			{regexp.MustCompile(`<`), defHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defHandler(MORE_EQL, ">=")},
			{regexp.MustCompile(`>`), defHandler(MORE, ">")},
			{regexp.MustCompile(`@`), defHandler(SQRT, "@")},
			{regexp.MustCompile(`\|\|`), defHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defHandler(AND, "&&")},
			{regexp.MustCompile(`\.\.`), defHandler(DBL_DOT, "..")},
			{regexp.MustCompile(`\.`), defHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`:`), defHandler(COLON, ":")},
			{regexp.MustCompile(`\?`), defHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defHandler(COMMA, ",")},
			{regexp.MustCompile(`\+\+`), defHandler(DBL_PLUS, "++")},
			{regexp.MustCompile(`--`), defHandler(DBL_MINS, "--")},
			{regexp.MustCompile(`\+=`), defHandler(PLUS_EQL, "+=")},
			{regexp.MustCompile(`-=`), defHandler(MINS_EQL, "-=")},
			{regexp.MustCompile(`\+`), defHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defHandler(MINS, "-")},
			{regexp.MustCompile(`//`), defHandler(DBL_SLSH, "/")},
			{regexp.MustCompile(`/`), defHandler(SLSH, "/")},
			{regexp.MustCompile(`\*`), defHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defHandler(PERC, "%")},
			{regexp.MustCompile(`\$`), defHandler(MODL, "$")},
			{regexp.MustCompile(`\^`), defHandler(POWR, "^")},
		},
	}
}

func numHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.rem())
	lex.push(NewToken(NUMBER, match))
	lex.advanceNum(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.rem())

	if match != "" {
		lex.advanceNum(len(match))
	}
}

func strHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.rem())
	stringLiteral := lex.rem()[match[0]:match[1]]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceNum(len(stringLiteral))
}

func charHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.rem())
	if len(match) < 3 { // Minimum valid: <x>
		panic(fmt.Sprintf("Lexing Error: invalid character literal: %s\nline: %v", match, lex.pos))
	}

	inner := match[1 : len(match)-1] // Remove < and >
	if runes := []rune(inner); len(runes) != 1 {
		panic(fmt.Sprintf("Lexing Error: character literal must contain exactly one rune, got %d: %s\nline: %v", len(runes), inner, lex.pos))
	}

	lex.push(NewToken(RUNE, inner))
	lex.advanceNum(len(match))
}

func smblHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.rem())
	if kind, exists := relu[match]; exists {
		lex.push(NewToken(kind, match))
	} else {
		lex.push(NewToken(IDENTIFIER, match))
	}

	lex.advanceNum(len(match))
}
