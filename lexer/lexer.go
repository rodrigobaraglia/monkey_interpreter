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

//NextToken returns a token
func (l *Lexer) NextToken() t.Token {
	var tok t.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = t.Token{Type: t.Eq, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(t.Assign, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = t.Token{Type: t.NotEq, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(t.Bang, l.ch)
		}
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
	case '-':
		tok = newToken(t.Minus, l.ch)
	case '*':
		tok = newToken(t.Asterisk, l.ch)
	case '/':
		tok = newToken(t.Slash, l.ch)
	case '<':
		tok = newToken(t.LT, l.ch)
	case '>':
		tok = newToken(t.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = t.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = t.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = t.Int
			return tok
		}
		tok = newToken(t.Illegal, l.ch)

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

//Added '?' character, sugested but not included in T. Ball's original implementation
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '?'
}

func (l *Lexer) readChar() {
	//This if-else block could be replaced by peakChar()
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	//On creation, position and readPosition = 0
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	beggining := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	end := l.position
	return l.input[beggining:end]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	beggining := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	end := l.position
	return l.input[beggining:end]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
