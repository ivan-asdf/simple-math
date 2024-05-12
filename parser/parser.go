package parser

import (
	"fmt"
	"strconv"

	"github.com/ivan-asdf/simple-math/eval"
	"github.com/ivan-asdf/simple-math/token"
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
		return &eval.Expr{Prev: prev, Op: eval.OP_NONE, Value: number}, nil
	}

	op, err := p.parseOp()
	if err != nil {
		return nil, err
	}
	number, err := p.parseNumber()
	if err != nil {
		return nil, err
	}
	return &eval.Expr{Prev: prev, Op: op, Value: number}, nil
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
		return eval.OP_NONE, SyntaxError{expected: "operation", after: p.peekPrevToken()}
	}
	switch t.Type {
	case token.Plus:
		return eval.OP_PLUS, nil
	case token.Minus:
		return eval.OP_MINUS, nil
	case token.MultipliedBy:
		return eval.OP_MULTI, nil
	case token.DividedBy:
		return eval.OP_DIV, nil
	case token.Cubed:
		fallthrough
	case token.Squared:
		return eval.OP_NONE, UnsupportedOperationError{op: t}
	}
	return eval.OP_NONE, SyntaxError{expected: "operation", got: t}
}
func (p *Parser) Parse() error {
	fmt.Println()
	err := p.parseInitialKeyword()
	if err != nil {
		return err
	}

	var lastExpr *eval.Expr
	for {
		expr, err := p.parseExpression(lastExpr)
		if err != nil {
			return err
		}
		fmt.Println(expr)
		if expr == nil {
			break
		}
		lastExpr = expr

		t := p.peekNextToken()
		if t == nil {
			return SyntaxError{expected: `termination keyword "?"`, after: p.peekPrevToken()}
		} else if t.Type == token.QuestionMarkKeyword {
			break
		}
	}
	fmt.Println()
	fmt.Println(lastExpr)
	result, err := lastExpr.Evaluate()
	if err != nil {
		return err
	}
	fmt.Println("EVAL RESULT: ", result)
	return nil
}
