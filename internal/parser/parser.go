package parser

import (
	"fmt"
	"strconv"

	"github.com/ivan-asdf/simple-math/internal/eval"
	"github.com/ivan-asdf/simple-math/internal/token"
)

type Parser struct {
	currentTokenIndex int
	tokens            []token.Token
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) getNextToken() *token.Token {
	if p.currentTokenIndex >= len(p.tokens) {
		return nil
	}
	index := p.currentTokenIndex
	p.currentTokenIndex++
	return &p.tokens[index]
}

func (p *Parser) peekNextToken() *token.Token {
	if p.currentTokenIndex >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.currentTokenIndex]
}

func (p *Parser) peekPrevToken() *token.Token {
	if p.currentTokenIndex <= 0 {
		return nil
	}
	return &p.tokens[p.currentTokenIndex-1]
}

func (p *Parser) checkIfMathQuestion() error {
	for t := p.getNextToken(); t != nil; t = p.getNextToken() {
		if t.Type == token.Number {
			p.currentTokenIndex = 0
			return nil
		}
	}
	return NonMathQuestionError{}
}

func (p *Parser) parseInitialKeyword() error {
	t := p.getNextToken()
	if t.Type != token.WhatIs {
		return SyntaxError{expected: `"What is" keyword`}
	}

	return nil
}

func (p *Parser) parseExpression(prev *eval.Expr) (*eval.Expr, error) {
	if prev == nil {
		number, err := p.parseNumber()
		if err != nil {
			return nil, err
		}
		return eval.NewExpr(prev, eval.OpNone, number)
	}

	op, err := p.parseOp()
	if err != nil {
		return nil, err
	}
	number, err := p.parseNumber()
	if err != nil {
		return nil, err
	}
	return eval.NewExpr(prev, op, number)
}

func (p *Parser) parseNumber() (int, error) {
	t := p.getNextToken()
	if t == nil {
		return 0, SyntaxError{expected: "number", after: p.peekPrevToken()}
	}
	if t.Type != token.Number {
		return 0, SyntaxError{expected: "number", got: t}
	}
	value, err := strconv.Atoi(t.Value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (p *Parser) parseOp() (eval.Op, error) {
	t := p.getNextToken()
	if t == nil {
		return eval.OpNone, SyntaxError{expected: "operation", after: p.peekPrevToken()}
	}
	switch t.Type {
	case token.Plus:
		return eval.OpPlus, nil
	case token.Minus:
		return eval.OpMinus, nil
	case token.MultipliedBy:
		return eval.OpMulti, nil
	case token.DividedBy:
		return eval.OpDiv, nil
	case token.Cubed:
		fallthrough
	case token.Squared:
		return eval.OpNone, UnsupportedOperationError{op: t}
	default:
		return eval.OpNone, SyntaxError{expected: "operation", got: t}
	}
}

func (p *Parser) Parse() (*eval.Expr, error) {
	err := p.checkIfMathQuestion()
	if err != nil {
		return nil, err
	}
	fmt.Println()
	err = p.parseInitialKeyword()
	if err != nil {
		return nil, err
	}

	var lastExpr *eval.Expr
	for {
		expr, err := p.parseExpression(lastExpr)
		if err != nil {
			return nil, err
		}
		fmt.Println(expr)
		if expr == nil {
			break
		}
		lastExpr = expr

		t := p.peekNextToken()
		if t == nil {
			return nil, SyntaxError{expected: `termination keyword "?"`, after: p.peekPrevToken()}
		} else if t.Type == token.QuestionMarkKeyword {
			break
		}
	}
	return lastExpr, nil
}
