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

func TestNextToken2(t *testing.T) {

	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	`
	// !-/*5;
	// 5 < 10 > 5;

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LParen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.RBrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.LParen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.RParen, ")"},
		{token.Semicolon, ";"},
		{token.EOF, ""},
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
