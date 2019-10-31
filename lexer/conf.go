package lexer

const (

	Plus TokenType = iota
	Minus
	Star
	Slash
	PPlus
	PlusEQ
	GT
	GE
	MinusEQ
	MMinus
	StarEQ
	SlashEQ
	EQ

	IntLiteral  	//int type
	DoubleLiteral	//double type

	Identifier		//variable

	//keyword
	INT
	DOUBLE
	BOOL
	NIL
	RUNE

)


