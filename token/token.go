package token

type TokenType int

const (
	WordToken TokenType = iota
	NumberToken
	SymbolToken
)

type Token struct {
	Value string
	Type  TokenType
}


