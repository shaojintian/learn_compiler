package lexer

import (

	"learn_compiler/my_log"

	"regexp"
)

//'int a = 1'
func(l *Lexer) lex() ([]*Token,error){
	tokens := make([]*Token,0)
	for _,ch := range l.source{
		switch {

		case IsAlpha(ch):
				tok := l.lexIdentifierAndKeyWords(ch)
		case IsDigital(ch):
				tok := l.lex



		}
	}
	return tokens,nil
}

func(l *Lexer) lexIdentifierAndKeyWords(ch rune) *Token{
	tok := &Token{//实例化一个Token，把Token的地址赋值给tok，*tok取到Token的具体属性
		typ:Identifier,
		text:[]rune{ch},
	}


	return tok
}

func (l *Lexer)SkipSpace() error {
	var ch rune
	for {
		ch = l.source[l.index]
		if ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r' {
			l.index++
		}else {
				break
		}
	}
	return nil

}

func IsAlpha(ch rune) bool{
	ok ,err := regexp.MatchString(`[a-zA-Z_]`,string(ch))
	if err != nil{
		my_log.LogPrint(err)
	}
	return ok
}

func IsDigital(ch rune) bool{
	ok,err := regexp.MatchString(`[0-9]`,string(ch))
	if err != nil{
		my_log.LogPrint(err)
	}
	return ok
}

func(l *Lexer) lexIntLiteral(ch rune) bool{

}

