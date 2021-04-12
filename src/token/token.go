package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERRISk = "*"
	SLASH     = "/"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keyWords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"if" : IF,
	"else" : ELSE,
	"return" : RETURN,
	"true" : TRUE,
	"false" : FALSE,
}

func LookUpIdent(ident string) TokenType {

	if tok, ok := keyWords[ident]; ok {
		return tok
	}

	return IDENT
}
