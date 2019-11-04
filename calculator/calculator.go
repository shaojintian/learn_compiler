package calculator

import "learn_compiler/lexer"

type Calculator interface {
	prog(reader lexer.TokenReader)
}

type MyCalculator struct {

}