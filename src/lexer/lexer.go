package lexer

import (
	"github.com/nivekithan/monkey/src/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhiteSpace()

	ch := lexer.ch

	switch ch {
	case '=':
		tok = newToken(token.ASSIGN, ch)
	case ';':
		tok = newToken(token.SEMICOLON, ch)
	case ',':
		tok = newToken(token.COMMA, ch)
	case '+':
		tok = newToken(token.PLUS, ch)
	case '-':
		tok = newToken(token.MINUS, ch)
	case '!':
		tok = newToken(token.BANG, ch)
	case '*':
		tok = newToken(token.ASTERRISk, ch)
	case '/':
		tok = newToken(token.SLASH, ch)
	case '(':
		tok = newToken(token.LPAREN, ch)
	case ')':
		tok = newToken(token.RPAREN, ch)
	case '{':
		tok = newToken(token.LBRACE, ch)
	case '}':
		tok = newToken(token.RBRACE, ch)
	case '<':
		tok = newToken(token.LT, ch)
	case '>':
		tok = newToken(token.GT, ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifer()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readIdentifer() string {
	position := lexer.position

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
