package eval

import "fmt"

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
	return fmt.Sprintf("(%v %s %d)", e.Prev, e.Op, e.Value)
}

func (e *Expr) Evaluate() int {
	if e.Prev == nil {
		return e.Value
	}
	return e.Prev.Evaluate() + e.Value
}


