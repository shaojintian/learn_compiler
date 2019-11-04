package ast

import "fmt"

//打印AST

func DumpAST(node Node,indent string)  {
	fmt.Println(indent,node.GetType()," ",node.GetText())
	
}
