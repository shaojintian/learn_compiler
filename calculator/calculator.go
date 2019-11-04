package calculator

import (
	"learn_compiler/ast"
	"learn_compiler/lexer"
)

type Calculator interface {
	prog(lexer.TokenReader) ast.Node//AST根节点
	parse([]rune)ast.Node //lexer分析后得到tokens,返回AST根节点
	intDeclare(lexer.TokenReader) ast.Node //int 变量声明表达式
	additive(lexer.TokenReader) ast.Node//加法表达式
	primary(reader lexer.TokenReader) ast.Node//基础表达式


}

type MyCalculator struct {

}


func (c *MyCalculator) prog(reader lexer.TokenReader) ast.Node{
	return ast.NewNode(ast.Program,[]rune("Calculator"))
}