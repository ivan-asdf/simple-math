package lexer

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ivan-asdf/simple-math/token"
)

type Lexer struct {
	input         string
	regexpPattern *regexp.Regexp
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	var groupNamesPatterns [token.LEN]string
	groupNamesPatterns[token.WhatIs] = `\bwhat is\b`
	groupNamesPatterns[token.QuestionMarkKeyword] = `\?`
	groupNamesPatterns[token.Plus] = `\bplus\b`
	groupNamesPatterns[token.Minus] = `\bminus\b`
	groupNamesPatterns[token.MultipliedBy] = `\bmultiplied by\b`
	groupNamesPatterns[token.DividedBy] = `\bdivided by\b`
	groupNamesPatterns[token.Cubed] = `\bcubed\b`
	groupNamesPatterns[token.Squared] = `\bsquared\b`
	groupNamesPatterns[token.Number] = `\b\d+\b`
	groupNamesPatterns[token.Word] = `[^\s\?]+`

	var patternStrings []string
	for groupName, pattern := range groupNamesPatterns {
		groupPattern := fmt.Sprintf(`(?P<%d>%s)`, groupName, pattern)
		patternStrings = append(patternStrings, groupPattern)
	}
	fmt.Println(patternStrings)
	l.regexpPattern = regexp.MustCompile("(?i)" + strings.Join(patternStrings, "|"))
	fmt.Println(l.regexpPattern)

	return l
}

func (l *Lexer) Lex() []token.Token {
	var tokens []token.Token
	matches := l.regexpPattern.FindAllStringSubmatchIndex(l.input, -1)
	for _, m := range matches {
		groupMatches := m[2:]
		fmt.Println(groupMatches)
		for i, groupIndex := 0, 0; i < len(groupMatches); i, groupIndex = i+2, groupIndex+1 {
			if groupMatches[i] != -1 {
				fmt.Println("INDEX: ", i, "GROUP:, ", groupIndex)
				startIndex, endIndex := groupMatches[i], groupMatches[i+1]
				value := l.input[startIndex:endIndex]
				fmt.Println(value)
				tokens = append(tokens, token.Token{Value: value, Type: token.TokenType(groupIndex), Begin: startIndex, End: endIndex})
				break
			}
		}
	}
	return tokens
}
