package lexer

import (
	"fmt"
	"log"
	"unicode"
	"unicode/utf8"
)

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

type Lexer struct {
	input  string
	pos    int
	length int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, length: len(input)}
}

func Lex(input string) {
	lexer := NewLexer(input)
	for {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		fmt.Printf("Token: %v, Value: %s\n", token.Type, token.Value)
	}
}

func (l *Lexer) NextToken() *Token {
	for l.pos < l.length {
		r, width := utf8.DecodeRuneInString(l.input[l.pos:])
		if unicode.IsSpace(r) {
			l.pos += width
			continue
		}

		if unicode.IsLetter(r) {
			return l.scanWord()
		}

		if unicode.IsDigit(r) {
			return l.scanNumber()
		}

		if unicode.IsSymbol(r) || unicode.IsPunct(r) {
			l.pos += width
			return &Token{Type: SymbolToken, Value: string(r)}
		}

		log.Fatal("Lexer error: Unrecognized character")
		l.pos += width
	}

	return nil
}

func (l *Lexer) scanWord() *Token {
	start := l.pos
	for l.pos < l.length {
		r, width := utf8.DecodeRuneInString(l.input[l.pos:])
		if !unicode.IsLetter(r) {
			break
		}
		l.pos += width
	}
	return &Token{Type: WordToken, Value: l.input[start:l.pos]}
}

func (l *Lexer) scanNumber() *Token {
	start := l.pos
	for l.pos < l.length {
		r, width := utf8.DecodeRuneInString(l.input[l.pos:])
		if !unicode.IsDigit(r) {
			break
		}
		l.pos += width
	}
	return &Token{Type: NumberToken, Value: l.input[start:l.pos]}
}
