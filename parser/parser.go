package parser

import (
	"rodrigobaraglia/monkeyimpl/ast"
	lx "rodrigobaraglia/monkeyimpl/lexer"
	t "rodrigobaraglia/monkeyimpl/token"
)

type Parser struct {
	l         *lx.Lexer
	curToken  t.Token
	peekToken t.Token
}

func NewParser(l *lx.Lexer) *Parser {
	p := &Parser{l: l}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
