package eval

import (
	"fmt"
)

type Op int

const (
	OpPlus Op = iota
	OpMinus
	OpMulti
	OpDiv
	OpNone
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
	case OpPlus:
		return "+"
	case OpMinus:
		return "-"
	case OpMulti:
		return "*"
	case OpDiv:
		return "/"
	case OpNone:
		return "NONE"
	default:
		return ""
	}
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
	case OpPlus:
		return result + e.Value, nil
	case OpMinus:
		return result - e.Value, nil
	case OpMulti:
		return result * e.Value, nil
	case OpDiv:
		if e.Value == 0 {
			return 0, fmt.Errorf("Evaluation error: division by zero")
		}
		return result / e.Value, nil
	case OpNone:
		return 0, fmt.Errorf(`Evaluation error: "None" operation %s`, getOpString(e.Op))
	default:
		return 0, fmt.Errorf("Evaluation error: unknown operation %s", getOpString(e.Op))
	}
}
