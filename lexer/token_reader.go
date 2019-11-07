package lexer

import (
	"errors"
	"fmt"
	"learn_compiler/my_log"
)

type TokenReader interface {
	Read() Token      //从stream读下一个Token，且取出
	Peek() Token      //从stream读下一个Token,但不取出
	Unread()          //恢复一个Token到stream
	GetPosition() int //获取stream当前读取位置
	SetPosition(int)  //设置token stream读取位置
	GetToks() []Token
	PrintToks()//打印toks
}

type MyTokenReader struct {
	tokens []Token
	pos    int
}

func NewTokenReader(tokens []Token) *MyTokenReader {
	return &MyTokenReader{
		tokens: tokens,
		pos:    0,
	}
}



func (r *MyTokenReader) Read() Token {
	if r.pos < len(r.tokens) {
		tok := r.tokens[r.pos]
		r.pos++
		return tok
	}

	return NullToken()
}

func (r *MyTokenReader) Peek() Token {
	if r.pos < len(r.tokens) {
		tok := r.tokens[r.pos]
		return tok
	}

	return NullToken()

}
func (r *MyTokenReader) Unread()  {
	//0------n-1
	if r.pos >= 1{
		r.pos--
	}else{
		my_log.LogFatal(errors.New("unread token stream err"))
	}

}
func (r *MyTokenReader) GetPosition() int {
	return r.pos
}
func (r *MyTokenReader) SetPosition(pos int) {
	if pos >= 0 && pos < len(r.tokens){
		r.pos = pos
	}else{
		my_log.LogFatal(errors.New("setPos token stream err"))
	}
}
func (r *MyTokenReader)GetToks()[]Token{
	return r.tokens
}
func (r *MyTokenReader)PrintToks(){
	for _,tok := range r.GetToks(){
		fmt.Printf("tok:{%v,%v}\n",TokenType2Str[tok.GetTyp()],string(tok.GetText()))
	}
}