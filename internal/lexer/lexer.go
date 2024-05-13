package lexer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ivan-asdf/simple-math/internal/token"
)

type Lexer struct {
	regexpPattern *regexp.Regexp
}

func NewLexer() *Lexer {
	l := Lexer{}
	var groupPatterns [token.LEN]string
	groupPatterns[token.WhatIs] = `\bwhat is\b`
	groupPatterns[token.QuestionMarkKeyword] = `\?`
	groupPatterns[token.Plus] = `\bplus\b`
	groupPatterns[token.Minus] = `\bminus\b`
	groupPatterns[token.MultipliedBy] = `\bmultiplied by\b`
	groupPatterns[token.DividedBy] = `\bdivided by\b`
	groupPatterns[token.Cubed] = `\bcubed\b`
	groupPatterns[token.Squared] = `\bsquared\b`
	groupPatterns[token.Number] = `\b\d+\b`
	groupPatterns[token.Word] = `[^\s\?]+`

	var patternStrings [token.LEN]string
	for groupIndex, pattern := range groupPatterns {
		groupPattern := fmt.Sprintf(`(?P<%d>%s)`, groupIndex, pattern)
		patternStrings[groupIndex] = groupPattern
	}
	fmt.Println(patternStrings)

	l.regexpPattern = regexp.MustCompile("(?i)" + strings.Join(patternStrings[:], "|"))
	fmt.Println(l.regexpPattern)

	return &l
}

func (l *Lexer) Lex(input string) []token.Token {
	var tokens []token.Token
	matches := l.regexpPattern.FindAllStringSubmatchIndex(input, -1)
	for _, m := range matches {
		groupMatches := m[2:]
		fmt.Println(groupMatches)
		for i, groupIndex := 0, 0; i < len(groupMatches); i, groupIndex = i+2, groupIndex+1 {
			if groupMatches[i] != -1 {
				fmt.Println("INDEX: ", i, "GROUP:, ", groupIndex)
				startIndex, endIndex := groupMatches[i], groupMatches[i+1]
				value := input[startIndex:endIndex]
				fmt.Println(value)
				tokens = append(tokens, token.Token{Value: value, Type: token.TokenType(groupIndex), Begin: startIndex, End: endIndex})
				break
			}
		}
	}
	return tokens
}
