package eval

import (
	"fmt"
)

type Op int

const (
	OP_PLUS Op = iota
	OP_MINUS
	OP_MULTI
	OP_DIV
	OP_NONE
)

type Expr struct {
	Prev  *Expr
	Op    Op
	Value int
}

func (e Expr) String() string {
	if e.Prev == nil {
		return fmt.Sprintf("%d", e.Value)
	}
	return fmt.Sprintf("(%v %s %d)", e.Prev, getOpString(e.Op), e.Value)
}

func getOpString(op Op) string {
	switch op {
	case OP_PLUS:
		return "+"
	case OP_MINUS:
		return "-"
	case OP_MULTI:
		return "*"
	case OP_DIV:
		return "/"
	case OP_NONE:
		return "NONE"
	}
	return ""
}

func (e *Expr) Evaluate() (int, error) {
	if e.Prev == nil {
		return e.Value, nil
	}

	result, err := e.Prev.Evaluate()
	if err != nil {
		return 0, err
	}
	switch e.Op {
	case OP_PLUS:
		return result + e.Value, nil
	case OP_MINUS:
		return result - e.Value, nil
	case OP_MULTI:
		return result * e.Value, nil
	case OP_DIV:
		if e.Value == 0 {
			return 0, fmt.Errorf("Evaluation error: division by zero")
		}
		return result / e.Value, nil
	default:
		return 0, fmt.Errorf("Evaluation error: unknown operation %s", getOpString(e.Op))
	}
}
