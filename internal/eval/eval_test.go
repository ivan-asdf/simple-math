package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePass(t *testing.T) {
	t.Parallel()

	// Single value with no operator
	expr1 := &Expr{prev: nil, op: OpNone, value: 5}
	err := validate(expr1)
	assert.Nil(t, err)

	// Two operands with operator
	expr2 := &Expr{prev: expr1, op: OpPlus, value: 2}
	err = validate(expr2)
	assert.Nil(t, err)
}

func TestValidateFail(t *testing.T) {
	t.Parallel()

	// Invalid operator
	expr1 := &Expr{prev: nil, op: 32312, value: 5}
	err := validate(expr1)
	assert.EqualErrorf(t, err, "invalid expression: 5", "Errors should match")

	// Invalid operator
	expr1 = &Expr{prev: nil, op: -3213, value: 5}
	err = validate(expr1)
	assert.EqualErrorf(t, err, "invalid expression: 5", "Errors should match")

	// Invalid operator
	expr1 = &Expr{prev: nil, op: OpNoneEnd, value: 5}
	err = validate(expr1)
	assert.EqualErrorf(t, err, "invalid expression: 5", "Errors should match")

	// Single operand yet operator present
	expr1 = &Expr{prev: nil, op: OpPlus, value: 5}
	err = validate(expr1)
	assert.EqualErrorf(t, err, "invalid expression: 5", "Errors should match")

	// Two operands yet not operator set
	expr1 = &Expr{prev: nil, op: OpNone, value: 5}
	expr2 := &Expr{prev: expr1, op: OpNone, value: 2}
	err = validate(expr2)
	assert.EqualErrorf(t, err, "invalid expression: (5 NONE 2)", "Errors should match")
}

func TestSimpleEvaluation(t *testing.T) {
	t.Parallel()

	expr1 := &Expr{prev: nil, op: OpNone, value: 4}
	expr2 := &Expr{prev: expr1, op: OpPlus, value: 3}
	result, err := expr2.Evaluate()
	assert.Nil(t, err)
	assert.Equal(t, 7, result)
}

func TestSimpleEvaluation2(t *testing.T) {
	t.Parallel()

	expr1 := &Expr{prev: nil, op: OpNone, value: 3}
	expr2 := &Expr{prev: expr1, op: OpPlus, value: 2}
	expr3 := &Expr{prev: expr2, op: OpMulti, value: 3}
	result, err := expr3.Evaluate()
	assert.Nil(t, err)
	assert.Equal(t, 15, result)
}

func TestExprString(t *testing.T) {
	t.Parallel()

	expr1 := &Expr{prev: nil, op: OpNone, value: 3}
	expr2 := &Expr{prev: expr1, op: OpPlus, value: 2}
	expr3 := &Expr{prev: expr2, op: OpMulti, value: 3}
	assert.Equal(t, "((3 + 2) * 3)", expr3.String())
}

func TestVeryNestedSum(t *testing.T) {
	t.Parallel()

	prevExpr := &Expr{prev: nil, op: OpNone, value: 1}
	for i := 0; i < 1*1000*1000-1; i++ {
		expr := &Expr{prev: prevExpr, op: OpPlus, value: 1}
		prevExpr = expr
	}
	result, err := prevExpr.Evaluate()
	assert.Nil(t, err)
	assert.Equal(t, 1*1000*1000, result)
}

func TestVeryNestedMultiplication(t *testing.T) {
	t.Parallel()

	prevExpr := &Expr{prev: nil, op: OpNone, value: 2}
	for i := 0; i < 31; i++ {
		expr := &Expr{prev: prevExpr, op: OpMulti, value: 2}
		prevExpr = expr
	}
	result, err := prevExpr.Evaluate()
	assert.Nil(t, err)
	assert.Equal(t, 4294967296, result)
}

func TestDivisionByZeroFail(t *testing.T) {
	t.Parallel()

	expr1 := &Expr{prev: nil, op: OpNone, value: 4}
	expr2 := &Expr{prev: expr1, op: OpDiv, value: 0}
	result, err := expr2.Evaluate()
	assert.ErrorIsf(t, err, ErrEvalDivisionByZero, "Errors should match")
	assert.Equal(t, 0, result)
}
