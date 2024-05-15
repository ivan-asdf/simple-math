package parser

import (
	"testing"

	"github.com/ivan-asdf/simple-math/internal/eval"
	"github.com/ivan-asdf/simple-math/internal/lexer"
	"github.com/ivan-asdf/simple-math/internal/token"
	"github.com/stretchr/testify/assert"
)

func genTokens(input string) []token.Token {
	l := lexer.NewLexer()
	tokens := l.Lex(input)
	return tokens
}

func evaluate(t *testing.T, e *eval.Expr) int {
	result, err := e.Evaluate()
	assert.Nil(t, err)
	return result
}

func parse(input string) (*eval.Expr, error) {
	p := NewParser(genTokens(input))
	return p.Parse()
}

func TestParseInitialKeyword(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("What is"))
	err := p.parseInitialKeyword()
	assert.Nil(t, err)

	p = NewParser(genTokens("What is what is"))
	err = p.parseInitialKeyword()
	assert.Nil(t, err)
	err = p.parseInitialKeyword()
	assert.Nil(t, err)
}

func TestParseInitialKeywordFail(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("dasd"))
	err := p.parseInitialKeyword()
	assert.EqualErrorf(t, err, "Syntax error: expected \"What is\" keyword", "Errors should match")
}

func TestParseNumber(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("5"))
	number, err := p.parseNumber()
	assert.Nil(t, err)
	assert.Equal(t, 5, number)

	p = NewParser(genTokens("1 2 3 "))
	number, err = p.parseNumber()
	assert.Nil(t, err)
	assert.Equal(t, 1, number)
	number, err = p.parseNumber()
	assert.Nil(t, err)
	assert.Equal(t, 2, number)
	number, err = p.parseNumber()
	assert.Nil(t, err)
	assert.Equal(t, 3, number)
}

func TestParseNumberFail(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("word"))
	number, err := p.parseNumber()
	assert.EqualErrorf(t, err, "Syntax error: expected number, got \"word\" at 0-4", "Errors should match")
	assert.Equal(t, 0, number)
}

func TestParseOp(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("plus"))
	op, err := p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpPlus, op)

	p = NewParser(genTokens("minus"))
	op, err = p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpMinus, op)

	p = NewParser(genTokens("multiplied by"))
	op, err = p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpMulti, op)

	p = NewParser(genTokens("divided by"))
	op, err = p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpDiv, op)

	p = NewParser(genTokens("plus minus multiplied by divided by"))
	op, err = p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpPlus, op)
	op, err = p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpMinus, op)
	op, err = p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpMulti, op)
	op, err = p.parseOp()
	assert.Nil(t, err)
	assert.Equal(t, eval.OpDiv, op)
}

func TestParseOpFail(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("asd"))
	op, err := p.parseOp()
	assert.EqualErrorf(t, err, "Syntax error: expected operation, got \"asd\" at 0-3", "Errors should match")
	assert.Equal(t, eval.OpNone, op)
}

func TestParseExpression(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("4 plus 3"))

	expr1, err := p.parseExpression(nil)
	assert.Nil(t, err)
	expected_expr1, err := eval.NewExpr(nil, eval.OpNone, 4)
	assert.Nil(t, err)
	assert.Equal(t, expected_expr1, expr1)

	expr2, err := p.parseExpression(expr1)
	assert.Nil(t, err)
	expected_expr2, err := eval.NewExpr(expr1, eval.OpPlus, 3)
	assert.Nil(t, err)
	assert.Equal(t, expected_expr2, expr2)
}

func TestParseExpressionFail(t *testing.T) {
	t.Parallel()

	p := NewParser(genTokens("dasda"))

	expr, err := p.parseExpression(nil)
	assert.EqualErrorf(t, err, "Syntax error: expected number, got \"dasda\" at 0-5", "Errors should match")
	assert.Nil(t, expr)
}

func TestSimpleExpression(t *testing.T) {
	t.Parallel()

	input := "What is 5?"
	p := NewParser(genTokens(input))
	result, err := p.Parse()
	assert.Nil(t, err)
	assert.Equal(t, 5, evaluate(t, result))
}

func TestSimpleExpression2(t *testing.T) {
	t.Parallel()

	input := "What is 5 plus 5 multiplied by 5?"
	p := NewParser(genTokens(input))
	result, err := p.Parse()
	assert.Nil(t, err)
	assert.Equal(t, 50, evaluate(t, result))
}

func TestSyntaxErrorFail(t *testing.T) {
	t.Parallel()

	input := "What is 5"
	expr, err := parse(input)
	assert.EqualErrorf(t, err, "Syntax error: expected termination keyword \"?\", after \"5\" at 8-9", "Errors should match")
	assert.Nil(t, expr)

	input = "What 5"
	expr, err = parse(input)
	assert.EqualErrorf(t, err, "Syntax error: expected \"What is\" keyword", "Errors should match")
	assert.Nil(t, expr)

	input = "What is 5 plus?"
	expr, err = parse(input)
	assert.EqualErrorf(t, err, "Syntax error: expected number, got \"?\" at 14-15", "Errors should match")
	assert.Nil(t, expr)

	input = "What is 5 airplane?"
	expr, err = parse(input)
	assert.EqualErrorf(t, err, "Syntax error: expected operation, got \"airplane\" at 10-18", "Errors should match")
	assert.Nil(t, expr)

	input = "What is 5."
	expr, err = parse(input)
	assert.EqualErrorf(t, err, "Syntax error: expected operation, got \".\" at 9-10", "Errors should match")
	assert.Nil(t, expr)
}

func TestUnsupportedOperationErrorFail(t *testing.T) {
	t.Parallel()

	input := "What is 5 cubed"
	expr, err := parse(input)
	assert.EqualErrorf(t, err, "Unsupported error: Unsupported operation \"cubed\" at 10-15", "Errors should match")
	assert.Nil(t, expr)

	input = "What is 5 squared"
	expr, err = parse(input)
	assert.EqualErrorf(t, err, "Unsupported error: Unsupported operation \"squared\" at 10-17", "Errors should match")
	assert.Nil(t, expr)
}

func TestNonMathQuestionErrorFails(t *testing.T) {
	t.Parallel()

	input := "What is the President of the United States?"
	expr, err := parse(input)
	assert.EqualErrorf(t, err, "non-math question(no numbers found in question)", "Errors should match")
	assert.Nil(t, expr)

	input = "Wt dsa dasd,xzvzx,xcasdsadw q"
	expr, err = parse(input)
	assert.EqualErrorf(t, err, "non-math question(no numbers found in question)", "Errors should match")
	assert.Nil(t, expr)
}
