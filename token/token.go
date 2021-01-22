package token

//Type of Token
type Type string

//Token data structre (cf. notes (1) for further development)
type Token struct {
	Type    Type
	Literal string
}

const (
	//Illegal tags an illegal token
	Illegal Type = "Illegal"
	//EOF tags the end of the file or input
	EOF Type = "EOF"
	//Ident tags an identifier token
	Ident Type = "Ident"
	//Int tags an ingeger
	Int Type = "Int"
	//Assign tags an assignement operator token
	Assign Type = "="
	//Plus tags a plus operator token
	Plus Type = "+"
	//Minus tags a minus operator token
	Minus Type = "-"
	//Bang tags a negation operator token
	Bang Type = "!"
	//Asterisk tags a multipication operator token
	Asterisk Type = "*"
	//Slash tags a division operator token
	Slash Type = "/"
	//LT tags a "lesser than" operator token
	LT Type = "<"
	//GT tags a "greater than" operator token
	GT Type = ">"
	//Comma tags a comma token
	Comma Type = ","
	//Semicolon tags a semicolon token
	Semicolon Type = ";"
	//LParen tags an opening parentheses token
	LParen Type = "("
	//RParen tags a closing parentheses token
	RParen Type = ")"
	//LBrace tags an opening brace token
	LBrace Type = "{"
	//RBrace tags a closing brace token
	RBrace Type = "}"
	//Function tags a function definition keyword token
	Function Type = "Function"
	//Let tags a variable definition keyword token
	Let Type = "Let"
	//True tags a boolean 1 token
	True Type = "True"
	//False tags a boolean 0 token
	False Type = "False"
	//If tags a conditional main branch keyword token
	If Type = "If"
	//Else tags a conditional alternate branch keyword token
	Else Type = "Else"
	//Return tags a return value token
	Return Type = "Return"
	//Eq tags an equality boolean operator token
	Eq Type = "=="
	//NotEq tags an inequality boolean operator token
	NotEq Type = "!="
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
