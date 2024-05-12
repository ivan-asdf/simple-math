package parser

import (
	"fmt"
	"strconv"

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
	return &p.tokens[p.currentTokenIndex - 1]
}

func (p *Parser) parseInitialKeyword() error {
	t := p.getNextToken()
	if t.Type != token.WhatIs {
    return SyntaxError{expected: `"What is" keyword`}
	}

	return nil
}

func (p *Parser) parseExpression(prev *Expr) (*Expr, error) {
	if prev == nil {
		number, err := p.parseNumber()
		if err != nil {
      return nil, err
		}
		return &Expr{Prev: prev, Op: "", Value: number}, nil
	}

	op, err := p.parseOp()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	number, err := p.parseNumber()
	if err != nil {
		return nil, err
	}
	return &Expr{Prev: prev, Op: op, Value: number}, nil
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

func (p *Parser) parseOp() (string, error) {
	t := p.getNextToken()
	if t == nil {
    return "", SyntaxError{expected: "operation", after: p.peekPrevToken()}
	}
	switch t.Type {
	case token.Plus:
		return OP_PLUS, nil
	case token.Minus:
		return OP_MINUS, nil
	}
  return "", SyntaxError{expected: "operation", got: t}
}

const OP_PLUS = "+"
const OP_MINUS = "-"

type Expr struct {
	Prev  *Expr
	Op    string
	Value int
}

func (e Expr) String() string {
	if e.Prev == nil {
		return fmt.Sprintf("%d", e.Value)
	}
	return fmt.Sprintf("(%v", e.Prev) + fmt.Sprintf(" %s %d)", e.Op, e.Value)
}

func (e *Expr) Evaluate() int {
	if e.Prev == nil {
		return e.Value
	}
	return e.Prev.Evaluate() + e.Value
}

func (p *Parser) Parse() error {
	fmt.Println()
	err := p.parseInitialKeyword()
	if err != nil {
		return err
	}

	var lastExpr *Expr
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
	fmt.Println("EVAL RESULT: ", lastExpr.Evaluate())
	return nil
}
