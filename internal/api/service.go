package api

import (
	"fmt"

	"github.com/ivan-asdf/simple-math/internal/lexer"
	"github.com/ivan-asdf/simple-math/internal/parser"
)

type Service struct {
	lexer     lexer.Lexer
	errorsLog map[ErrorLogKey]int
}

type ErrorLogKey struct {
	endpoint   string
	expression string
	errorType  string
}

func NewService() Service {
	return Service{
		lexer:     *lexer.NewLexer(),
		errorsLog: make(map[ErrorLogKey]int),
	}
}

func (s *Service) SaveError(endpoint string, expression string, err error) {
	var errType string
	switch err.(type) {
	case parser.SyntaxError:
		errType = "Syntax error"
	case parser.UnsupportedOperationError:
		errType = "Unsupported operation error"
	case parser.NonMathQuestionError:
		errType = "Non-math question error"
	default:
		return
	}

	s.errorsLog[ErrorLogKey{endpoint, expression, errType}]++
}

func (s *Service) Evaluate(input string) (int, error) {
	fmt.Println(input)
	fmt.Println()

	tokens := s.lexer.Lex(input)
	fmt.Println(tokens)

	parser := parser.NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		s.SaveError(EvaluateEndpoint, input, err)
		return 0, err
	}
	return expr.Evaluate()
}

func (s *Service) Validate(input string) error {
	fmt.Println(input)
	fmt.Println()

	tokens := s.lexer.Lex(input)
	fmt.Println(tokens)

	parser := parser.NewParser(tokens)
	_, err := parser.Parse()
	if err != nil {
		s.SaveError(ValidateEndpoint, input, err)
		return err
	}
	return err
}

func (s *Service) Errors() []map[string]string {
	var errorFrequencies []map[string]string
	for entry, frequency := range s.errorsLog {
		errorFrequencies = append(errorFrequencies, map[string]string{
			"expression": entry.expression,
			"endpoint":   entry.endpoint,
			"frequency":  fmt.Sprint(frequency),
			"type":       entry.errorType,
		})
	}
	return errorFrequencies
}
