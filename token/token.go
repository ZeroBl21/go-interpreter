package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers + Literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 123456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Table of the avaliable keywords
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// Checks if the given indentifier is in a fact a keyword. If it is,
// returns the keyword's TokenType constant.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
