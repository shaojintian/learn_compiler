package lexer

import (
	"errors"
	"learn_compiler/my_log"

	"regexp"
)

//init token
//eg:'int a = 1 age >= 45 intA = 67  1+2*3/6'
func (l *Lexer) InitToken(ch rune) Token {
	//initialize a token
	tok := Token{
		typ:  INIT,
		text: []rune{ch},
	}
	//init lexer
	l.status = INIT
	switch {
	case IsAlpha(ch):
		tok.typ = Identifier //变量
	case IsDigital(ch):
		tok.typ = IntLiteral //整型数字
	case ch == '=':
		tok.typ = Assign
	case ch == '>':
		tok.typ = GT
	case ch == '<':
		tok.typ = Less
	case ch == '/':
		tok.typ = Slash
	case ch == '*':
		tok.typ = Star
	case ch == '+':
		tok.typ = Plus
	case ch == '-':
		tok.typ = Minus
	case ch == '(':
		tok.typ = LeftParen
	case ch == ')':
		tok.typ = RightParen



	default:
		my_log.LogPrint(errors.New("illegal token"))

	}
	l.status = tok.typ
	return tok
}

//'int     a    =    1'
func (l *Lexer) Lex() error {
	var token Token
	var ch rune
	//skip whitespaces in beginning
	l.SkipSpace()
	for ; l.index < len(l.source); l.index++ {

		ch = l.source[l.index]

		switch l.status {
		case INIT:
			token = l.InitToken(ch)
		case Identifier:
			if (IsAlpha(ch)) {
				token.text = append(token.text, ch)
				l.lexIdentifierAndKeyWords(token.text, token)

			} else if ch == ' ' {
				l.tokens = append(l.tokens, token)
				l.SkipSpace()
				ch = l.source[l.index]
				token = l.InitToken(ch)
			}

		case INT:
			l.SkipSpace()
			ch = l.source[l.index]
			token = l.InitToken(ch)

		case IntLiteral:
			if IsDigital(ch) {
				token.text = append(token.text, ch)

			} else if ch == ' ' {
				l.tokens = append(l.tokens, token)
				l.SkipSpace()
				ch = l.source[l.index]
				token = l.InitToken(ch)
			} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
				l.tokens = append(l.tokens, token)
				token = l.InitToken(ch)

			}

		case Assign:
			if ch == '=' {

			}
			l.tokens = append(l.tokens, token)
			l.SkipSpace()
			ch = l.source[l.index]
			token = l.InitToken(ch)

		case GT:
			if (ch == '=') {
				token.typ = GE
				l.status = GE
				token.text = append(token.text, '=')
				l.tokens = append(l.tokens, token)
			} else if (ch == ' ') {
				l.SkipSpace()
				ch = l.source[l.index]

				token = l.InitToken(ch)
			}
		case GE:
			l.SkipSpace()
			ch = l.source[l.index]
			token = l.InitToken(ch)
		case Plus:
			if (ch == '+') {

			} else if (ch == '=') {

			} else {
				l.tokens = append(l.tokens, token)
				token = l.InitToken(ch)
			}
		case Minus:
			if (ch == '=') {

			} else {
				l.tokens = append(l.tokens, token)
				token = l.InitToken(ch)
			}

		case Star:
			if (ch == '=') {

			} else {
				l.tokens = append(l.tokens, token)
				token = l.InitToken(ch)
			}
		case Slash:
			if (ch == '=') {

			} else {
				l.tokens = append(l.tokens, token)
				token = l.InitToken(ch)
			}

		}

		if l.index == len(l.source)-1 {
			l.tokens = append(l.tokens, token)
		}
	}
	return nil
}

func (l *Lexer) lexIdentifierAndKeyWords(s []rune, token Token) {
	if l.source[l.index+1] == ' ' {
		if status, ok := KEYWORDS[string(s)]; ok {
			l.status = status
			token.typ = status
			l.tokens = append(l.tokens, token)

		}
	}
}

func (l *Lexer) SkipSpace() {
	var ch rune
	for {
		ch = l.source[l.index]
		//fmt.Println(string(ch))
		if ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r' {
			l.index++
		} else {
			break
		}
	}

}

func IsAlpha(ch rune) bool {
	ok, err := regexp.MatchString(`[a-zA-Z_]`, string(ch))
	if err != nil {
		my_log.LogPrint(err)
	}
	return ok
}

func IsDigital(ch rune) bool {
	ok, err := regexp.MatchString(`[0-9]`, string(ch))
	if err != nil {
		my_log.LogPrint(err)
	}
	return ok
}

