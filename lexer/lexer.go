package lexer

import (
	"log"
	"unicode"
	"unicode/utf8"

	"github.com/ivan-asdf/simple-math/token"
)

type Lexer struct {
	input  string
	pos    int
	length int
}

func newLexer(input string) *Lexer {
	return &Lexer{input: input, length: len(input)}
}

func Lex(input string) []token.Token {
	lexer := newLexer(input)
  var tokens []token.Token
	for {
		token := lexer.NextToken()
		if token == nil {
			break
		}
    tokens = append(tokens, *token)
		// fmt.Printf("Token: %v, Value: %s\n", token.Type, token.Value)
	}
  return tokens
}

func (l *Lexer) NextToken() *token.Token {
	for l.pos < l.length {
    // TODO: add error handling for example invalid UTF-8
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
			return &token.Token{Type: token.SymbolToken, Value: string(r)}
		}

		log.Fatal("Lexer error: Unrecognized character")
		l.pos += width
	}

	return nil
}

func (l *Lexer) scanWord() *token.Token {
	start := l.pos
	for l.pos < l.length {
		r, width := utf8.DecodeRuneInString(l.input[l.pos:])
		if !unicode.IsLetter(r) {
			break
		}
		l.pos += width
	}
	return &token.Token{Type: token.WordToken, Value: l.input[start:l.pos]}
}

func (l *Lexer) scanNumber() *token.Token {
	start := l.pos
	for l.pos < l.length {
		r, width := utf8.DecodeRuneInString(l.input[l.pos:])
		if !unicode.IsDigit(r) {
			break
		}
		l.pos += width
	}
	return &token.Token{Type: token.NumberToken, Value: l.input[start:l.pos]}
}
