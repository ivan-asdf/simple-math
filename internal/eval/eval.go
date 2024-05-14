package eval

import (
	"errors"
	"fmt"
)

type Op int

const (
	OpNone Op = iota
	OpPlus
	OpMinus
	OpMulti
	OpDiv
	OpNoneEnd
)

var (
	ErrEvalInvalidExpr      = errors.New("invalid expression")
	ErrEvalDivisionByZero   = errors.New("division by zero")
	ErrEvalNoneOperation    = errors.New(`"NONE" operation`)
	ErrEvalUnknownOperation = errors.New("unknown operation")
)

type Expr struct {
	prev  *Expr
	op    Op
	value int
}

func NewExpr(prev *Expr, op Op, value int) (*Expr, error) {
	expr := &Expr{prev: prev, op: op, value: value}
	err := validate(expr)
	return expr, err
}

func validate(e *Expr) error {
	if e.op == OpNone && e.prev == nil {
		return nil
	}
	if e.op > OpNone && e.op < OpNoneEnd && e.prev != nil {
		return nil
	}
	return fmt.Errorf("%w: %v", ErrEvalInvalidExpr, e)
}

func (e Expr) String() string {
	if e.prev == nil {
		return fmt.Sprintf("%d", e.value)
	}
	return fmt.Sprintf("(%v %s %d)", e.prev, getOpString(e.op), e.value)
}

func getOpString(op Op) string {
	switch op {
	case OpPlus:
		return "+"
	case OpMinus:
		return "-"
	case OpMulti:
		return "*"
	case OpDiv:
		return "/"
	case OpNone, OpNoneEnd:
		return "NONE"
	default:
		return ""
	}
}

func (e *Expr) Evaluate() (int, error) {
	if e.prev == nil {
		return e.value, nil
	}

	result, err := e.prev.Evaluate()
	if err != nil {
		return 0, err
	}
	switch e.op {
	case OpPlus:
		return result + e.value, nil
	case OpMinus:
		return result - e.value, nil
	case OpMulti:
		return result * e.value, nil
	case OpDiv:
		if e.value == 0 {
			return 0, ErrEvalDivisionByZero
		}
		return result / e.value, nil
	case OpNone, OpNoneEnd:
		return 0, fmt.Errorf(`%w: %s`, ErrEvalNoneOperation, getOpString(e.op))
	default:
		return 0, fmt.Errorf("%w: %s", ErrEvalUnknownOperation, getOpString(e.op))
	}
}
