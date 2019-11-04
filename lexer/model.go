package lexer

type Token struct {
	typ  TokenType
	text []rune
}

// Lexer
type Lexer struct {
	source []rune
	index  int
	status TokenType
	tokens []*Token
}

func NewToken(typ TokenType,text []rune) *Token{
	return &Token{
		typ:typ,
		text:text,
	}
}



// initialize  function of lexer
func NewLexer(source []rune) (*Lexer, error) {
	tokens := make([]*Token, 0)
	return &Lexer{
		source: source,
		index:  0,
		status: INIT,
		tokens: tokens,
	}, nil
}


