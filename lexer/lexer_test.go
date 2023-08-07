package lexer

import (
	"jojo/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		const five = 5
		const ten = 10

		fn add(x, y) {
   		x + y
		}

		var result = add(five, ten)
 	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.CONST, "const"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		// {token.SEMICOLON, ";"},
		{token.CONST, "const"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		// {token.SEMICOLON, ";"},
		// {token.VAR, "var"},
		// {token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		// {token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		// {token.SEMICOLON, ";"},
		{token.VAR, "var"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		// {token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
