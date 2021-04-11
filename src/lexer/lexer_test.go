package lexer

import (
	"github.com/nivekithan/monkey/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "{"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, correctToken := range tests {
		tok := lexer.NextToken()

		if tok.Type != correctToken.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected = %q, got = %q", i, correctToken.expectedType, tok.Type)
		}

		if tok.Literal != correctToken.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected = %q, got = %q", i, correctToken.expectedLiteral, tok.Literal)

		}
	}

}
