package lexer

type TokenType int

type Token struct  {
	typ TokenType
	text []rune

}
// Lexer
type Lexer struct {
	source []rune
	index 	int
}
// initialize  function of lexer
func NewLexer(source []rune) (*Lexer,error){
	return &Lexer{
		source:source,
		index:0,
	},nil
}

