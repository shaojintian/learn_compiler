package lexer

import "fmt"

func Main(){
	source :=  "int s = 123 age >= 45 intA = 67 1+2*3/6"
	lexer ,err := NewLexer([]rune(source))
	if err != nil{
		panic(err)
	}


	tokens,err := lexer.lex()
	if err != nil{
		panic(err)
	}
	for _,token := range tokens{
		fmt.Printf("{%v : %v}",token.typ,token.text)
	}


}
