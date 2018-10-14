package token

type TokenType string

/*

We defined the TokenType to be a string. That allows us to
use many different values as TokenTypes, which in turn
allows us to distinguish between different types of tokens.
*/
type Token struct {
	Type TokenType
	Literal string
}

const (
	// Signifiers a token/character we don't know about
	ILLEGAL 	= "ILLEGAL"
	// end of file, tells parser when to stop
	EOF				= "EOF"

	// Identifiers + literals
	IDENT 		= "IDENT" // add, foobar, x, y
	INT				= "INT"

	// Operators
	ASSIGN 		= "="
	PLUS			=	"+"

	// Delimiters
	COMMA			= ","
	SEMICOLON	= ";"

	LPAREN		=	"("
	RPAREN		= ")"
	LBRACE 		= "{"
	RBRACE		= "}"

	// Keywords
	FUNCTION	= "FUNCTION"
	LET   		= "LET"
)

