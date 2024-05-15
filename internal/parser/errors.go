package parser

import (
	"errors"
	"fmt"

	"github.com/ivan-asdf/simple-math/internal/token"
)

var NonMathQuestionError = errors.New("non-math question(no numbers found in question)")

type SyntaxError struct {
	expected string
	got      *token.Token
	after    *token.Token
}

func positionString(t *token.Token) string {
	return fmt.Sprintf(`"%s" at %d-%d`, t.Value, t.Begin, t.End)
}

func (e SyntaxError) Error() string {
	if e.got != nil {
		return fmt.Sprintf(`Syntax error: expected %s, got %s`, e.expected, positionString(e.got))
	}
	if e.after != nil {
		return fmt.Sprintf(`Syntax error: expected %s, after %s`, e.expected, positionString(e.after))
	}
	return fmt.Sprintf(`Syntax error: expected %s`, e.expected)
}

type UnsupportedOperationError struct {
	op *token.Token
}

func (e UnsupportedOperationError) Error() string {
	return fmt.Sprintf(`Unsupported error: Unsupported operation %s`, positionString(e.op))
}
