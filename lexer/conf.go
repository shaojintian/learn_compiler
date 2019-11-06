package lexer

type TokenType int

const (
	Plus TokenType = iota
	Minus
	Star
	Slash
	PPlus
	PlusEQ
	GT
	GE
	Less
	LessEQ
	MinusEQ
	MMinus
	StarEQ
	SlashEQ
	EQ
	Assign
	LeftParen  //(
	RightParen	//)

	IntLiteral     //整数
	DoubleLiteral  //浮点数

	Identifier  //variable

	//keyword
	INT
	DOUBLE
	BOOL
	NIL
	RUNE



	NULL
	EOF
	INIT
)

var KEYWORDS = map[string]TokenType{
	"int":    INT,
	"double": DOUBLE,
	"bool":   BOOL,
	"nil":    NIL,
	"rune":   RUNE,
	"init":   INIT,
}

var TokenType2Str = map[TokenType]string{
	Plus:    "Plus",
	Minus:   "Minus",
	Star:    "Star",
	Slash:   "Slash",
	PPlus:   "PPlus",
	PlusEQ:  "PlusEQ",
	GT:      "GT",
	GE:      "GE",
	Less:    "Less",
	LessEQ:  "LessEQ",
	MinusEQ: "MinusEQ",
	MMinus:  "MMinus",
	StarEQ:  "StarEQ",
	SlashEQ: "SlashEQ",
	EQ:      "EQ",
	Assign:  "Assign",
	LeftParen  :"LeftParen",
	RightParen	:"RightParen",

	IntLiteral:    "IntLiteral",    //int type number
	DoubleLiteral: "DoubleLiteral", //double type number

	Identifier: "Identifier", //variable

	//keyword
	INT:    "INT",
	DOUBLE: "DOUBLE",
	BOOL:   " BOOL",
	NIL:    "NIL",
	RUNE:   "RUNE",

	EOF:  "EOF",
	INIT: "INIT",
	NULL: "NULL",
}

