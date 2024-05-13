package api

import (
	"fmt"

	"github.com/ivan-asdf/simple-math/lexer"
	"github.com/ivan-asdf/simple-math/parser"
)

type Service struct {
	lexer lexer.Lexer
}

func NewService() Service {
	return Service{
		lexer: *lexer.NewLexer(),
	}
}

func (s *Service) Evaluate(input string) (int, error) {
	// input := "What iS plUs 4?"
	fmt.Println(input)
	fmt.Println()

	tokens := s.lexer.Lex(input)
	fmt.Println(tokens)

	parser := parser.NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		return 0, err
	}
	fmt.Println("HERE ", expr)
	return expr.Evaluate()
}

func (s *Service) Validate(input string) error {
	// input := "What iS plUs 4?"
	fmt.Println(input)
	fmt.Println()

	tokens := s.lexer.Lex(input)
	fmt.Println(tokens)

	parser := parser.NewParser(tokens)
	_, err := parser.Parse()
	return err
}
