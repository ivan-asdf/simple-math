package parser

//
// import (
// 	"fmt"
// 	"log"
// 	"regexp"
// 	"strconv"
//
// 	"github.com/ivan-asdf/simple-math/token"
// )
//
// // var WHAT_KEYWORD_REGEX *regexp.Regexp = regexp.MustCompile("(?i)what")
// // const IS_KEYWORD = "is"
//
// type Parser struct {
// 	currentTokenIndex int
// 	tokens            []token.Token
//
// 	whatKeywordRegexp *regexp.Regexp
// 	isKeywordRegexp   *regexp.Regexp
// 	plusOpRegexp      *regexp.Regexp
// 	minusOpRegexp     *regexp.Regexp
// 	// plusOpRegexp      *regexp.Regexp
// 	// plusOpRegexp      *regexp.Regexp
// }
//
// func NewParser(tokens []token.Token) *Parser {
// 	return &Parser{
// 		tokens:            tokens,
// 		whatKeywordRegexp: regexp.MustCompile("(?i)what"),
// 		isKeywordRegexp:   regexp.MustCompile("(?i)is"),
// 		plusOpRegexp:      regexp.MustCompile("(?i)plus"),
// 		minusOpRegexp:     regexp.MustCompile("(?i)minus"),
// 	}
// }
//
// func (p *Parser) getNextToken() *token.Token {
// 	if p.currentTokenIndex >= len(p.tokens) {
// 		return nil
// 	}
// 	index := p.currentTokenIndex
// 	p.currentTokenIndex++
// 	return &p.tokens[index]
// }
//
// func (p *Parser) parseInitialKeyword() error {
// 	token1 := p.getNextToken()
// 	if token1.Type != token.WordToken {
// 		return fmt.Errorf("Syntax error: expected word token")
// 	}
// 	if !p.whatKeywordRegexp.MatchString(token1.Value) {
// 		return fmt.Errorf("Syntax error: expected keyword \"what\"")
// 	}
//
// 	token2 := p.getNextToken()
// 	if token2.Type != token.WordToken {
// 		return fmt.Errorf("Syntax error: expected word token")
// 	}
// 	if !p.isKeywordRegexp.MatchString(token2.Value) {
// 		return fmt.Errorf("Syntax error: expected keyword \"is\"")
// 	}
//
// 	return nil
// }
//
// // add error return
// func (p *Parser) parseExpression(prev *Expr) *Expr {
// 	if prev == nil {
// 		number, err := p.parseNumber()
//     if err != nil {
//       // return err
//     }
// 		return &Expr{Prev: prev, Op: "", Value: number}
// 	}
//
// 	op, err := p.parseOp()
// 	if err != nil {
// 		// return err
//     return nil
// 	}
//   number, err := p.parseNumber()
// 	if err != nil {
// 		// return err
//     return nil
// 	}
// 	return &Expr{Prev: prev, Op: op, Value: number}
// }
//
// func (p *Parser) parseNumber() (int, error) {
// 	t := p.getNextToken()
//   // if t == nil {
// 		// return 0, fmt.Errorf("Syntax error: expected number")
//   // }
// 	if t == nil || t.Type != token.NumberToken {
// 		return 0, fmt.Errorf("Syntax error: expected number")
// 	}
// 	value, err := strconv.Atoi(t.Value)
// 	if err != nil {
// 		log.Fatal(err) // handle later
// 	}
// 	return value, nil
// }
//
// func (p *Parser) parseOp() (string, error) {
// 	t := p.getNextToken()
//   // if t == nil {
// 		// return "", fmt.Errorf("Syntax error: expected operation string")
//   // }
// 	if t == nil || t.Type != token.WordToken {
// 		return "", fmt.Errorf("Syntax error: expected operation string")
// 	}
// 	if p.plusOpRegexp.MatchString(t.Value) {
// 		return OP_PLUS, nil
// 	}
// 	if p.minusOpRegexp.MatchString(t.Value) {
// 		return OP_MINUS, nil
// 	}
// 	return "", fmt.Errorf("Syntax error: invalid operation string: %s", t.Value)
// }
//
// const OP_PLUS = "+"
// const OP_MINUS = "-"
//
// type Expr struct {
// 	Prev  *Expr
// 	Op    string
// 	Value int
// }
//
// func (e Expr) String() string {
// 	if e.Prev == nil {
// 		return fmt.Sprintf("%d", e.Value)
// 	}
// 	return fmt.Sprintf("(%v", e.Prev) + fmt.Sprintf(" %s %d)", e.Op, e.Value)
// }
//
// func (e *Expr) Evaluate() int {
// 	if e.Prev == nil {
// 		return e.Value
// 	}
// 	return e.Prev.Evaluate() + e.Value
// }
//
// func (p *Parser) Parse() error {
// 	err := p.parseInitialKeyword()
// 	if err != nil {
// 		return err
// 	}
//
// 	var lastExpr *Expr
// 	for {
// 		expr := p.parseExpression(lastExpr)
//     // fmt.Println(expr)
// 		if expr == nil {
// 			break
// 		}
// 		lastExpr = expr
// 	}
// 	fmt.Println()
// 	fmt.Println(lastExpr)
// 	fmt.Println(lastExpr.Evaluate())
// 	return nil
// }
