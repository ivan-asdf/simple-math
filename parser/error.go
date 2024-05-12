package parser

import (
	"fmt"

	"github.com/ivan-asdf/simple-math/token"
)

type SyntaxError struct {
	expected string
	got      *token.Token
	after    *token.Token
}

func (e SyntaxError) Error() string {
	if e.got != nil {
		return fmt.Sprintf(
			`Syntax error: expected %s, got "%s" at %d-%d`,
			e.expected,
			// token.GetTokenTypeString(e.got.Type),
			e.got.Value,
			e.got.Begin,
			e.got.End)
	} else if e.after != nil {
		return fmt.Sprintf(
			`Syntax error: expected %s after "%s" at %d-%d`,
			e.expected,
			// token.GetTokenTypeString(e.got.Type),
			e.after.Value,
			e.after.Begin,
			e.after.End)
	} else {
		return fmt.Sprintf(
			`Syntax error: expected %s`,
			e.expected,
		)
	}
}
