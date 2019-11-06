package calculator

import (
	"errors"
	"learn_compiler/ast"
	"learn_compiler/lexer"
	"learn_compiler/my_log"
)

type Calculator interface {
	prog(*lexer.MyTokenReader) ast.Node           //AST根节点
	parse([]rune) ast.Node                     //lexer分析后得到tokens,返回AST根节点
	intDeclare(*lexer.MyTokenReader) ast.Node     //int 变量声明表达式
	additive(*lexer.MyTokenReader) ast.Node       //加法表达式
	multiplicative(*lexer.MyTokenReader) ast.Node //乘法表达式
	primary(reader *lexer.MyTokenReader) ast.Node //基础表达式
	evaluate(node ast.Node, indent string) int // 求某个node的值，并写出过程
}

type MyCalculator struct {
}

func (c *MyCalculator) prog(reader *lexer.MyTokenReader) ast.Node {
	node := ast.NewNode(ast.Program, []rune("Calculator"))

	child := c.additive(reader)
	if child.GetType() != ast.NUll {
		node.AddChild(child)
	}
}
func (c *MyCalculator) parse(toks []rune) ast.Node {

}
func (c *MyCalculator) intDeclare(r *lexer.MyTokenReader) ast.Node {

}
func (c *MyCalculator) additive(r *lexer.MyTokenReader) ast.Node {
	child1 := c.multiplicative(r)
	node := child1
	tok := r.Peek()
	if child1.GetType() != ast.NUll && tok.GetTyp() != lexer.NULL {
		if tok.GetTyp() == lexer.Plus || tok.GetTyp() == lexer.Minus {
			tok := r.Read()
			child2 := c.additive(r) //右递归
			if child2.GetType() != ast.NUll {
				node = ast.NewNode(ast.Additive, tok.GetText())
				node.AddChild(child1)
				node.AddChild(child2)
			}

		}
	}
	return node

}
func (c *MyCalculator) primary(r *lexer.MyTokenReader) ast.Node {
	tok := r.Peek()
	node := ast.NullNode()
	if tok.GetTyp() != lexer.NULL {
		tok = r.Read()
		switch tok.GetTyp() {
		case lexer.IntLiteral://整数
			node = ast.NewNode(ast.IntLiteral, tok.GetText())
		case lexer.Identifier://变量
			node = ast.NewNode(ast.Identifier,tok.GetText())
		case lexer.LeftParen:
			node = c.additive(r)//root
			if node.GetType() != ast.NUll{
				tok = r.Peek()
				if tok.GetTyp() == lexer.RightParen{
					tok = r.Read()
				}else {
					my_log.LogPrint(errors.New("expect rightParen"))
				}
			}else {
				my_log.LogPrint(errors.New("expect right additive syntax"))
			}

		}
	}

	return node


}
func (c *MyCalculator) multiplicative(r *lexer.MyTokenReader) ast.Node {

}
func (c *MyCalculator) evaluate(node ast.Node, indent string) int {

}
