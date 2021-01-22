package parser

import (
	"rodrigobaraglia/monkeyimpl/ast"
	lx "rodrigobaraglia/monkeyimpl/lexer"
	t "rodrigobaraglia/monkeyimpl/token"
)

/*Parser is composed of a lexer l and two tokens curToken and peekToken.
The first is the current token being parsed, the other is the next token to be parsed.*/
type Parser struct {
	l         *lx.Lexer
	curToken  t.Token
	peekToken t.Token
}

/*NewParser creates and initializes a Parser. It takes a lexer l as
an argument to initialize the field by the same name and calls nextToken()
twice to initialize fields curToken and peekToken.*/
func NewParser(l *lx.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/*ParseProgram returns the abstract syntax tree of a program*/
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curTokenIs(t.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case t.Let:
		return p.parseLetStatement()

	default:
		return nil
	}
}

/*parseLetStatement checks that the following context of a let statement corresponds to the proper form of assignement syntax and ends with a semicolon*/
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	/*Check that next token is an identifier*/
	if !p.expectPeek(t.Ident) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	/*Check that next token is an assignement operator*/
	if !p.expectPeek(t.Assign) {
		return nil
	}

	/*Skip expressions until a semicolon is found*/
	for !p.curTokenIs(t.Semicolon) {
		p.nextToken()
	}

	return stmt

}

func (p *Parser) curTokenIs(tokType t.Type) bool {
	return p.curToken.Type == tokType
}

func (p *Parser) peekTokenIs(tokType t.Type) bool {
	return p.peekToken.Type == tokType
}

func (p *Parser) expectPeek(tokType t.Type) bool {
	if !p.peekTokenIs(tokType) {
		return false
	}
	p.nextToken()
	return true
}
