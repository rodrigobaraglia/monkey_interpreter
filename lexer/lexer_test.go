package lexer

import (
	"rodrigobaraglia/monkeyimpl/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Assign, "="},
		{token.Plus, "+"},
		{token.LParen, "("},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.RBrace, "}"},
		{token.Comma, ","},
		{token.Semicolon, ";"},
	}

	l := NewLexer(input)

	for i, tCase := range tests {
		tok := l.NextToken()

		if tok.Type != tCase.expectedType {
			t.Fatalf("tests [%d] - wrong token type. expected=%q, got=%q",
				i, tCase.expectedType, tok.Type)
		}

		if tok.Literal != tCase.expectedLiteral {
			t.Fatalf("tests [%d] - wrong token literal. expected=%q, got=%q",
				i, tCase.expectedLiteral, tok.Literal)
		}
	}
}
