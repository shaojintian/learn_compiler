package lexer
//token
type Token struct {
	typ  TokenType
	text []rune
}

func NewToken(typ TokenType, text []rune) *Token {
	return &Token{
		typ:  typ,
		text: text,
	}
}

func NullToken() Token {
	return *NewToken(NULL, nil)

}

func (tok *Token) GetTyp() TokenType {
	return tok.typ
}

func (tok *Token) GetText() []rune {
	return tok.text
}

// Lexer
type Lexer struct {
	source []rune
	index  int
	status TokenType
	tokens []Token
}

// initialize  function of lexer
func NewLexer(source []rune) *Lexer {
	tokens := make([]Token, 0)
	return &Lexer{
		source: source,
		index:  0,
		status: INIT,
		tokens: tokens,
	}
}

func ( l *Lexer) GetToks() []Token{
	return l.tokens
}
