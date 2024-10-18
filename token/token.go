/*
Package token defines the token types and structures used in the Leopard programming language lexer.

This package includes token representation for various language constructs,
including identifiers, literals, operators, delimiters, and keywords.
*/
package token

// TokenType represents the type of a token.
type TokenType string

// Token represents a lexical token in the language, consisting of a type
// and its literal string value.
type Token struct {
	Type    TokenType
	Literal string
}

// Token type constants.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ..
	INT    = "INT"   // 12345
	STRING = "STRING"

	// Operators.
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords.
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// keywords maps string representations of keywords to their corresponding
// token types.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns the TokenType associated with the given identifier.
// If the identifier is not a keyword, it returns the IDENT token type.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
