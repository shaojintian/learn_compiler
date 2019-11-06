package calculator

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	//x := ast.NewNode(ast.Program, []rune("calculator"))
	source1 := []rune("2+3*5") //17
	source2 := []rune("2+3*5*2") //32
	source3 := []rune("2+3") //5
	source4 := []rune("(2+3)") //5

	cal := NewMyCal()
	fmt.Println(cal.evaluateAll(source1),"---",17)
	fmt.Println(cal.evaluateAll(source2),"---",32)
	fmt.Println(cal.evaluateAll(source3),"---",5)
	fmt.Println(cal.evaluateAll(source4),"---",5)





}
