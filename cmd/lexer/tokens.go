package lexer

import (
	"fmt"
	"slices"
)

type TokenKind int

// Token Kinds
const (
	// Base Operators
	EOF TokenKind = iota
	NUMBER
	STRING
	RUNE
	IDENTIFIER

	// Groupative Operators
	OPEN_BRACK
	OPEN_PAREN
	OPEN_CURLY
	CLOSE_BRACK
	CLOSE_PAREN
	CLOSE_CURLY

	// Assignment
	ASSIGN
	ARROW

	// Comparative Operators
	EQL
	NOT_EQL
	LESS
	MORE
	LESS_EQL
	MORE_EQL

	// Miscallaneous Objects
	COMMENT
	DOT
	DBL_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA

	// Logical Operators
	NOT
	AND
	OR
	XOR
	SHIFT_LEFT
	SHIFT_RIGHT

	// In/Decrementers
	DBL_PLUS
	DBL_MINS
	PLUS_EQL
	MINS_EQL

	// Operators
	PLUS
	MINS
	SLSH
	DBL_SLSH
	STAR
	POWR
	SQRT
	PERC
	MODL

	// RESERVED
	VAR
	CONST
	CLASS
	STRUCT
	TYPE
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	ELIF
	AUTO
	WHILE
	FOR
	KINDOF
	EXPORT
	IN

	// Types
	STR_TYPE
	INT_TYPE
	BOOL_TYPE
	LIST_TYPE
	DICT_TYPE
	RUNE_TYPE
	FUNC_TYPE
)

// Reserved Lookup
var relu map[string]TokenKind = map[string]TokenKind{
	"var":     VAR,
	"const":   CONST,
	"class":   CLASS,
	"struct":  STRUCT,
	"type":    TYPE,
	"new":     NEW,
	"import":  IMPORT,
	"from":    FROM,
	"fn":      FN,
	"if":      IF,
	"else":    ELSE,
	"elif":    ELIF,
	"auto":    AUTO,
	"while":   WHILE,
	"for":     FOR,
	"kindof":  KINDOF,
	"export":  EXPORT,
	"in":      IN,
	"str":     STR_TYPE,
	"int":     INT_TYPE,
	"boolean": BOOL_TYPE,
	"list":    LIST_TYPE,
	"dict":    DICT_TYPE,
	"rune":    RUNE_TYPE,
	"func":    FUNC_TYPE,
}

type Token struct {
	Kind  TokenKind
	Value string
}

// To color text in the terminal
func ColorText(text string, color string) string {
	var colors map[string]int = map[string]int{
		"reset":  0,
		"red":    31,
		"green":  32,
		"yellow": 33,
		"blue":   34,
		"cyan":   36,
		"purple": 35,
		"gray":   37,
		"white":  97,
	}

	realColor := fmt.Sprintf("\033[%vm", colors[color])
	return realColor + text + "\033[0m"
}

// Debug Function
func (token Token) Debug() {
	t, v := TokenKindString(token.Kind)
	if token.ContainedIn(IDENTIFIER, NUMBER, STRING, RUNE) {
		fmt.Printf("{ Kind: %s, Value: %s }\n", ColorText(t, "cyan"), ColorText(token.Value, "red"))
	} else {
		fmt.Printf("{ Kind: %s, Value: %s }\n", ColorText(t, "cyan"), ColorText(v, "red"))
	}
}

// This is tomfoolery
func (token Token) ContainedIn(expectedTokens ...TokenKind) bool {
	return slices.Contains(expectedTokens, token.Kind)
}

// Simplify the Parsing Process
func NewToken(kind TokenKind, value string) Token {
	return Token{kind, value}
}

// IDKWTNT Comment
func TokenKindString(kind TokenKind) (string, string) {
	switch kind {
	case EOF:
		return "eof", "end of file"
	case NUMBER:
		return "number", ""
	case STRING:
		return "string", ""
	case RUNE:
		return "rune", ""
	case IDENTIFIER:
		return "identifier", ""
	case OPEN_BRACK:
		return "open_bracket", "["
	case CLOSE_BRACK:
		return "close_bracket", "]"
	case OPEN_PAREN:
		return "open_paren", "("
	case CLOSE_PAREN:
		return "close_paren", ")"
	case OPEN_CURLY:
		return "open_curly", "{"
	case CLOSE_CURLY:
		return "close_curly", "}"
	case ASSIGN:
		return "assign", "="
	case ARROW:
		return "arrow", "=>"
	case EQL:
		return "equal", "=="
	case NOT_EQL:
		return "not_equal", "~="
	case NOT:
		return "not", "~"
	case LESS:
		return "less", "<"
	case MORE:
		return "greater", ">"
	case LESS_EQL:
		return "less_equal", "<="
	case MORE_EQL:
		return "more_equal", ">="
	case DOT:
		return "dot", "."
	case DBL_DOT:
		return "dot_dot", ".."
	case SEMI_COLON:
		return "semi_colon", ";"
	case COLON:
		return "colon", ":"
	case QUESTION:
		return "question", "?"
	case COMMA:
		return "comma", ","
	case AND:
		return "and", "&&"
	case OR:
		return "or", "||"
	case XOR:
		return "xor", "~|"
	case SHIFT_LEFT:
		return "shift_left", "<<"
	case SHIFT_RIGHT:
		return "shift_right", ">>"
	case DBL_MINS:
		return "minus_minus", "--"
	case DBL_PLUS:
		return "plus_plus", "++"
	case PLUS_EQL:
		return "plus_equal", "+="
	case MINS_EQL:
		return "minus_equal", "-="
	case PLUS:
		return "plus", "+"
	case MINS:
		return "minus", "-"
	case SLSH:
		return "slash", "/"
	case DBL_SLSH:
		return "integer slash", "//"
	case STAR:
		return "star", "*"
	case POWR:
		return "power", "**"
	case SQRT:
		return "square root", "@"
	case PERC:
		return "percent", "%"
	case MODL:
		return "module", "$"
	case VAR:
		return "var", "keyword"
	case CONST:
		return "const", "keyword"
	case CLASS:
		return "class", "keyword"
	case NEW:
		return "new", "keyword"
	case IMPORT:
		return "import", "keyword"
	case FROM:
		return "from", "keyword"
	case FN:
		return "func", "keyword"
	case IF:
		return "if", "keyword"
	case ELSE:
		return "else", "keyword"
	case AUTO:
		return "foreach", "keyword"
	case WHILE:
		return "while", "keyword"
	case FOR:
		return "for", "keyword"
	case KINDOF:
		return "kindof", "keyword"
	case EXPORT:
		return "export", "keyword"
	case IN:
		return "in", "keyword"
	case COMMENT:
		return "comment", "##"
	case STR_TYPE:
		return "type_ann", "str"
	case INT_TYPE:
		return "type_ann", "int"
	case LIST_TYPE:
		return "type_ann", "list"
	case DICT_TYPE:
		return "type_ann", "dictionary"
	case RUNE_TYPE:
		return "type_ann", "rune"
	case FUNC_TYPE:
		return "type_ann", "function"
	default:
		fmt.Println(fmt.Errorf("Lexing error: unknown argument found in lexer: %v", kind))
		return "unknown", "\\"
	}
}
