package calculator

import (
	"errors"
	"fmt"
	"learn_compiler/ast"
	"learn_compiler/lexer"
	"learn_compiler/my_log"
	"strconv"
	"strings"
)

type Calculator interface {
	Prog(*lexer.MyTokenReader) ast.Node           //AST根节点
	parse([]rune) ast.Node                        //lexer分析后得到tokens,返回AST根节点
	intDeclare(*lexer.MyTokenReader) ast.Node     //int 变量声明表达式
	additive(*lexer.MyTokenReader) ast.Node       //加法表达式
	multiplicative(*lexer.MyTokenReader) ast.Node //乘法表达式
	primary(reader *lexer.MyTokenReader) ast.Node //基础表达式
	evaluate(node ast.Node, indent string) int    // 求某个node的值，并写出过程
	evaluateAll([]rune) int                       //求脚本的结果
}

type MyCalculator struct {
}

func NewMyCal() *MyCalculator {
	return &MyCalculator{

	}
}

func (c *MyCalculator) Prog(reader *lexer.MyTokenReader) ast.Node {
	node := ast.NewNode(ast.Program, []rune("Calculator"))

	child := c.additive(reader)
	if child.GetType() != ast.NUll {
		node.AddChild(child)
	}
	return node
}

func (c *MyCalculator) parse(source []rune) ast.Node {
	var err error
	//lexer
	Lexer := lexer.NewLexer(source)
	err = Lexer.Lex()

	if err != nil {
		my_log.LogFatal(err)
	}

	toks := Lexer.GetToks()
	tokReader := lexer.NewTokenReader(toks)
	//ast root
	root := c.Prog(tokReader)

	return root

}
func (c *MyCalculator) intDeclare(r *lexer.MyTokenReader) ast.Node {
	return ast.NewNode(ast.Program,nil)
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
		case lexer.IntLiteral: //整数
			node = ast.NewNode(ast.IntLiteral, tok.GetText())
		case lexer.Identifier: //变量
			node = ast.NewNode(ast.Identifier, tok.GetText())
		case lexer.LeftParen:
			node = c.additive(r) //root
			if node.GetType() != ast.NUll {
				tok = r.Peek()
				if tok.GetTyp() == lexer.RightParen {
					tok = r.Read()
				} else {
					my_log.LogPrint(errors.New("expect rightParen"))
				}
			} else {
				my_log.LogPrint(errors.New("expect right additive syntax"))
			}

		}
	}

	return node

}
func (c *MyCalculator) multiplicative(r *lexer.MyTokenReader) ast.Node {
	child1 := c.primary(r)
	node := child1
	tok := r.Peek()
	if tok.GetTyp() != lexer.NULL && node.GetType() != ast.NUll {
		if tok.GetTyp() == lexer.Star {
			tok = r.Read()
			child2 := c.multiplicative(r)
			if child2.GetType() != ast.NUll {
				node = ast.NewNode(ast.Multiplicative, tok.GetText())
				node.AddChild(child1)
				node.AddChild(child2)
			} else {
				my_log.LogPrint(errors.New("expect right multiplicative syntax"))
			}
		}
	}

	return node
}
func (c *MyCalculator) evaluate(node ast.Node, indent string) int {
	var result int
	fmt.Println(indent, "Calculating: ", ast.NodeType2Str[node.GetType()])
	switch node.GetType() {
	case ast.Program:
		for _, child := range node.GetChildren() {
			result = c.evaluate(child, indent+"\t")
		}
		
	case ast.Additive:
		child1 := node.GetChildren()[0]
		value1 := c.evaluate(child1, indent+"\t")
		child2 := node.GetChildren()[1]
		value2 := c.evaluate(child2, indent+"\t")
		if strings.Compare(string(node.GetText()),"+") == 0{
			result = value1 + value2
		} else {
			result = value1 - value2
		}

	case ast.Multiplicative:
		child1 := node.GetChildren()[0]
		value1 := c.evaluate(child1, indent+"\t")
		child2 := node.GetChildren()[1]
		value2 := c.evaluate(child2, indent+"\t")
		if strings.Compare(string(node.GetText()),"*") == 0{
			result = value1 * value2
		} else {
			result = value1 / value2
		}
		 
	case ast.IntLiteral:
		var err error
		if result,err = strconv.Atoi(string(node.GetText())); err != nil{
			my_log.LogFatal(err)
		}

		 
	default:
	}
	fmt.Println(indent, "Result: ", result)
	return result
}

func (c *MyCalculator) evaluateAll(source []rune) int {
	root := c.parse(source)
	ast.DumpAST(root, "")
	return c.evaluate(root, "")

}
