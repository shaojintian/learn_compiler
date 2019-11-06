package calculator

import (
	"learn_compiler/ast"
	"log"
	"testing"
)

func TestCalculate(t *testing.T) {
	//x := ast.NewNode(ast.Program, []rune("calculator"))
	source1 := []rune("2+3*5") //17
	source2 := []rune("2+3*5*2") //32c
	source3 := []rune("2+3") //5
	source4 := []rune("(2+3)") //5

	cal := NewMyCal()
	if cal.evaluateAll(source1)!=17{
		log.Fatalln()
	}
	if cal.evaluateAll(source2)!=32{
		log.Fatalln()
	}
	if cal.evaluateAll(source3)!=5{
		log.Fatalln()
	}
	if cal.evaluateAll(source4)!=5{
		log.Fatalln()
	}

}

func TestMyCalculator_Parse(t *testing.T) {
	cal := NewMyCal()
	source1 := []rune("2+3*5") //17
	root := cal.parse(source1)
	ast.DumpAST(root,"")
}