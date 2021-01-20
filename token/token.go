package token

//Type of Token
type Type string

//Token data structre (cf. notes (1) for further development)
type Token struct {
	Type    Type
	Literal string
}

const (
	Illegal   = "Illegal"
	EOF       = "EOF"
	Ident     = "Ident"
	Int       = "Int"
	Assign    = "="
	Plus      = "+"
	Comma     = ","
	Semicolon = ";"
	LParen    = "("
	RParen    = ")"
	LBrace    = "{"
	RBrace    = "}"
	Function  = "Function"
	Let       = "Let"
)
