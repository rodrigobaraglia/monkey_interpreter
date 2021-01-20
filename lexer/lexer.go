package lexer

import t "rodrigobaraglia/monkeyimpl/token"

//Lexer stores a given token as input, two integers as current and next position and a byte
//representing the value of the character at current position. Cf notes (2) for further development.
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

//NewLexer creates a new Lexer struct and initializes it's input field to a given string
func NewLexer(input string) *Lexer {
	//Create lexer and initialize input
	l := &Lexer{input: input}

	//Initialize position, readPosition and ch
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	//On creation, position and readPosition = 0
	l.position = l.readPosition
	l.readPosition++
}

//NextToken returns a token
func (l *Lexer) NextToken() t.Token {
	var tok t.Token

	switch l.ch {
	case '=':
		tok = newToken(t.Assign, l.ch)
	case ';':
		tok = newToken(t.Semicolon, l.ch)
	case ',':
		tok = newToken(t.Comma, l.ch)
	case '(':
		tok = newToken(t.LParen, l.ch)
	case ')':
		tok = newToken(t.RParen, l.ch)
	case '{':
		tok = newToken(t.LBrace, l.ch)
	case '}':
		tok = newToken(t.RBrace, l.ch)
	case '+':
		tok = newToken(t.Plus, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = t.EOF
	}

	l.readChar()

	return tok
}

func newToken(tokenType t.Type, ch byte) t.Token {
	return t.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
