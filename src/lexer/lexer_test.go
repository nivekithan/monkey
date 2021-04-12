package lexer

import (
	"github.com/nivekithan/monkey/src/token"
	"testing"
)

type testToken = []struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestSingleWordToken(t *testing.T) {
	input := `=+(){}<>-/*!,;|=`

	tests := testToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERRISk, "*"},
		{token.BANG, "!"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.ILLEGAL, "|"},
		{token.ASSIGN, "="},
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

func TestDoubleWordsToken(t *testing.T) {
	input := `
	==
	!=
	`

	tests := testToken{
		{token.EQ, "=="},
		{token.NOT_EQ, "!="},
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

func TestKeyWordsToken(t *testing.T) {
	input := `
	let
	false
	true
	fn
	if 
	else
	return
	`
	tests := testToken{
		{token.LET, "let"},
		{token.FALSE, "false"},
		{token.TRUE, "true"},
		{token.FUNCTION, "fn"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.RETURN, "return"},
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

func TestMediumNextToken(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	
	let add  = fn (x, y) {
		x + y;
	};
	
	let result = add(five, ten);`

	tests := testToken{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
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
