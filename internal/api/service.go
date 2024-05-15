package api

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ivan-asdf/simple-math/internal/lexer"
	"github.com/ivan-asdf/simple-math/internal/parser"
)

type Service struct {
	lexer            lexer.Lexer
	errorsLog        map[ErrorLogKey]int
	errorsLogRWMutex sync.RWMutex
}

type ErrorLogKey struct {
	endpoint   string
	expression string
	errorType  string
}

const (
	ErrorTypeSyntaxError          = "Syntax error"
	ErrorTypeUnsupportedOperation = "Unsupported operation error"
	ErrorTypeNonMathQuestion      = "Non-math question error"
)

func NewService() Service {
	return Service{
		lexer:            *lexer.NewLexer(),
		errorsLog:        make(map[ErrorLogKey]int),
		errorsLogRWMutex: sync.RWMutex{},
	}
}

func (s *Service) SaveError(endpoint string, expression string, err error) {
	var errType string
	switch {
	case errors.As(err, &parser.SyntaxError{}):
		errType = ErrorTypeSyntaxError
	case errors.As(err, &parser.UnsupportedOperationError{}):
		errType = ErrorTypeUnsupportedOperation
	case errors.Is(err, parser.NonMathQuestionError):
		errType = ErrorTypeNonMathQuestion
	default:
		return
	}

	s.errorsLogRWMutex.Lock()
	s.errorsLog[ErrorLogKey{endpoint, expression, errType}]++
	s.errorsLogRWMutex.Unlock()
}

func (s *Service) Evaluate(input string) (int, error) {
	tokens := s.lexer.Lex(input)

	parser := parser.NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		s.SaveError(EvaluateEndpoint, input, err)
		return 0, err
	}
	return expr.Evaluate()
}

func (s *Service) Validate(input string) error {
	tokens := s.lexer.Lex(input)

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
	s.errorsLogRWMutex.RLock()
	for entry, frequency := range s.errorsLog {
		errorFrequencies = append(errorFrequencies, map[string]string{
			"expression": entry.expression,
			"endpoint":   entry.endpoint,
			"frequency":  fmt.Sprint(frequency),
			"type":       entry.errorType,
		})
	}
	s.errorsLogRWMutex.RUnlock()
	return errorFrequencies
}
