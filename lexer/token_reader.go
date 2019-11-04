package lexer

type TokenReader interface {
	Read() Token //从stream读下一个Token，且取出
	Peek() Token //从stream读下一个Token,但不取出
	Unread()     //恢复一个Token到stream
	GetPosition() int //获取stream当前读取位置
	SetPosition(int) 	//设置token stream读取位置

}

type MyTokenReader struct {
	tokens 	[]Token
	pos 	int
}
func NewTokenReader(tokens []Token) *MyTokenReader{
	return &MyTokenReader{
		tokens:tokens,
		pos:0,
	}
}

func(reader *MyTokenReader) Read()Token{
	
}

func(reader *MyTokenReader) Peek()Token{

}
func(reader *MyTokenReader) Unread()Token{

}
func(reader *MyTokenReader) GetPosition() int{

}
func(reader *MyTokenReader) SetPosition(pos int) {

}