package token

import "fmt"

type TokenType int

const (
	WhatIs TokenType = iota
	QuestionMarkKeyword
	Plus
	Minus
	MultipliedBy
	DividedBy
	Cubed
	Squared
	Number
	Word
	LEN
)

type Token struct {
	Value string
	Type  TokenType
	Begin int
	End   int
}

func (t Token) String() string {
	return fmt.Sprintf("{%v %s %d-%d}", t.Value, GetTokenTypeString(t.Type), t.Begin, t.End)
}

func GetTokenTypeString(tt TokenType) string {
	switch tt {
	case WhatIs:
		return "WhatIs"
	case QuestionMarkKeyword:
		return "QuestionMark"
	case Plus:
		return "Plus"
	case Minus:
		return "Minux"
	case Number:
		return "Number"
	case Word:
		return "Word"

	}
	return ""
}
