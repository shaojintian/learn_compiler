package ast

/*接口中要返回interface，因为这样可以使得属性private，任何属性都能通过get/set...方法得到，面向interface编程
//struct 和 interface 的区别仅仅是多了属性，但由于要屏蔽属性为private,且属性均能由方法得到，只有方法就好了，故全为interface
所以能传interface就传interface*/

type Node interface {
	GetParent() Node
	GetChildren() []Node
	GetType() NodeType
	GetText() []rune
	AddChild(Node)
}

type MyNode struct {
	parent           Node
	children         []Node
	readOnlyChildren []Node
	typ              NodeType
	text             []rune
}

func NewNode(
	typ NodeType, text []rune,
) *MyNode {

	return &MyNode{
		parent:           nil,
		children:         nil,
		readOnlyChildren: nil,
		typ:              typ,
		text:             text,
	}
}

func (n *MyNode) GetParent() Node {
	return n.parent
}

func (n *MyNode) GetChildren() []Node {
	return n.children
}
func (n *MyNode) GetType() NodeType {
	return n.typ
}
func (n *MyNode) GetText() []rune {
	return n.text
}



func (n *MyNode) AddChild(node Node) {
	n.children = append(n.children,node)
}

func NullNode() Node{
	return NewNode(NUll,nil)
}

