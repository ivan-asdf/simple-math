package eval

import (
	"errors"
	"testing"
)

func TestSimpleEvaluation(t *testing.T) {
	t.Parallel()
	expr1 := Expr{prev: nil, op: OpNone, value: 5}
	expr2 := Expr{prev: &expr1, op: OpPlus, value: 2}
	result, err := expr2.Evaluate()
	if err != nil {
		t.Errorf("Should have not returned error: %s", err)
	}
	if result != 7 {
		t.Errorf("Result should have been 7 but was %d", result)
	}
}

func TestVeryNestedSum(t *testing.T) {
	t.Parallel()
	var prevExpr *Expr
	for i := 0; i < 1*1000*1000; i++ {
		expr := &Expr{prev: prevExpr, op: OpPlus, value: 1}
		prevExpr = expr
	}
	result, err := prevExpr.Evaluate()
	if err != nil {
		t.Errorf("Should have not returned error: %s", err)
	}
	if result != 1*1000*1000 {
		t.Errorf("Result should have been 1 000 000 but was %d", result)
	}
}

// Test very nested multiplication

// Test evaluation division by zero fail, none operator, unknow operator

func TestValidatePass(t *testing.T) {
	t.Parallel()

	// Single value with no operator
	expr1 := &Expr{prev: nil, op: OpNone, value: 5}
	err := validate(expr1)
	if err != nil {
		t.Errorf("Should not have returned error")
	}

	// Two operands with operator
	expr2 := &Expr{prev: expr1, op: OpPlus, value: 2}
	err = validate(expr2)
	if err != nil {
		t.Errorf("Should not have returned error")
	}
}

func TestValidateFail(t *testing.T) {
	t.Parallel()
	// Single operand yet operator present
	expr0 := &Expr{prev: nil, op: OpPlus, value: 5}
	err := validate(expr0)
	if !errors.Is(err, ErrEvalInvalidExpr) {
		t.Errorf("Should have returned error but didnt")
	}
	// Two operands yet not operator set
	expr1 := &Expr{prev: nil, op: OpNone, value: 5}
	expr2 := &Expr{prev: expr1, op: OpNone, value: 2}
	err = validate(expr2)
	if !errors.Is(err, ErrEvalInvalidExpr) {
		t.Errorf("Should have returned error but didnt")
	}
}

// Test Expr.String()
