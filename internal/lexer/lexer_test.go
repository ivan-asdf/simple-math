package lexer

import (
	"testing"

	"github.com/ivan-asdf/simple-math/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestLex1(t *testing.T) {
	t.Parallel()

	l := NewLexer()
	input := "WhAt iS 2 plUs 3?word 1word ..word."
	expectedTokens := []token.Token{
		{Value: "WhAt iS", Type: token.WhatIs, Begin: 0, End: 7},
		{Value: "2", Type: token.Number, Begin: 8, End: 9},
		{Value: "plUs", Type: token.Plus, Begin: 10, End: 14},
		{Value: "3", Type: token.Number, Begin: 15, End: 16},
		{Value: "?", Type: token.QuestionMarkKeyword, Begin: 16, End: 17},
		{Value: "word", Type: token.Word, Begin: 17, End: 21},
		{Value: "1word", Type: token.Word, Begin: 22, End: 27},
		{Value: "..word.", Type: token.Word, Begin: 28, End: 35},
	}
	resultTokens := l.Lex(input)
	assert.Equal(t, expectedTokens, resultTokens, "Generated tokens should be equal to expected")
}

func TestLex2(t *testing.T) {
	t.Parallel()

	l := NewLexer()
	input := "WhAtiS   2cubed  cubEd ?"
	expectedTokens := []token.Token{
		{Value: "WhAtiS", Type: token.Word, Begin: 0, End: 6},
		{Value: "2cubed", Type: token.Word, Begin: 9, End: 15},
		{Value: "cubEd", Type: token.Cubed, Begin: 17, End: 22},
		{Value: "?", Type: token.QuestionMarkKeyword, Begin: 23, End: 24},
	}
	resultTokens := l.Lex(input)
	assert.Equal(t, expectedTokens, resultTokens, "Generated tokens should be equal to expected")
}
