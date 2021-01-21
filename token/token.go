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
	Minus     = "-"
	Bang      = "!"
	Asterisk  = "*"
	Slash     = "/"
	LT        = "<"
	GT        = ">"
	Comma     = ","
	Semicolon = ";"
	LParen    = "("
	RParen    = ")"
	LBrace    = "{"
	RBrace    = "}"
	Function  = "Function"
	Let       = "Let"
	True      = "True"
	False     = "False"
	If        = "If"
	Else      = "Else"
	Return    = "Return"
	Eq        = "=="
	NotEq     = "!="
)

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

/*LookupIdent takes a token literal and looks it up in the keyword map. If the token is not a keyword, it returns
an Ident type tag, else, it returns the token associated with the given keyword.*/
func LookupIdent(ident string) Type {
	tok, ok := keywords[ident]
	if !ok {
		return Ident
	}
	return tok
}
