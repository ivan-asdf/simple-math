package parser

import (
	"fmt"
	"log"
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
	index := p.currentTokenIndex
	return &p.tokens[index]
}

func (p *Parser) parseInitialKeyword() error {
	t := p.getNextToken()
	if t.Type != token.WhatIs {
		return fmt.Errorf(`Syntax error: expected "What is" keyword`)
	}

	return nil
}

// add error return
func (p *Parser) parseExpression(prev *Expr) *Expr {
	if prev == nil {
		number, err := p.parseNumber()
		if err != nil {
			// return err
			fmt.Println("HERE")
			return nil
		}
		return &Expr{Prev: prev, Op: "", Value: number}
	}

	op, err := p.parseOp()
	if err != nil {
		// return err
		fmt.Println(err)
		return nil
	}
	number, err := p.parseNumber()
	if err != nil {
		// return err
		return nil
	}
	return &Expr{Prev: prev, Op: op, Value: number}
}

func (p *Parser) parseNumber() (int, error) {
	t := p.getNextToken()
	// if t == nil {
	// return 0, fmt.Errorf("Syntax error: expected number")
	// }
	if t == nil || t.Type != token.Number {
		return 0, fmt.Errorf("Syntax error: expected number")
	}
	value, err := strconv.Atoi(t.Value)
	if err != nil {
		log.Fatal(err) // handle later
	}
	return value, nil
}

func (p *Parser) parseOp() (string, error) {
	t := p.getNextToken()
	// if t == nil {
	// return "", fmt.Errorf("Syntax error: expected operation string")
	// }
	if t == nil {
		return "", fmt.Errorf("Syntax error: expected operation string, reaced end")
	}
	switch t.Type {
	case token.Plus:
		return OP_PLUS, nil
	case token.Minus:
		return OP_MINUS, nil
	}
	return "", fmt.Errorf("Syntax error: invalid operation string: %s", t.Value)
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
		expr := p.parseExpression(lastExpr)
		fmt.Println(expr)
		if expr == nil {
			break
		}
		lastExpr = expr

		t := p.peekNextToken()
		if t == nil {
			return fmt.Errorf(`Syntax error: expected termination keyword "?"`)
		} else if t.Type == token.QuestionMarkKeyword {
			break
		}
	}
	fmt.Println()
	fmt.Println(lastExpr)
	fmt.Println("EVAL RESULT: ", lastExpr.Evaluate())
	return nil
}
