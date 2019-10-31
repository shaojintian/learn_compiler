package lexer

import (
	"fmt"
	"testing"
)

func TestIsAlpha(t *testing.T) {

	fmt.Println(IsAlpha('a'))
}

func TestLexer(t *testing.T) {

	source := " 1+2*3/6"
	lexer, err := NewLexer([]rune(source))
	if err != nil {
		panic(err)
	}

	err = lexer.Lex()
	if err != nil {
		panic(err)
	}
	for _, token := range lexer.tokens {
		fmt.Printf("{%v : %s}\n", TokenType2Str[token.typ], string(token.text))
	}

}
