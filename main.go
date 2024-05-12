package main

import (
	"fmt"
	"log"

	"github.com/ivan-asdf/simple-math/lexer"
	"github.com/ivan-asdf/simple-math/parser"
)

func main() {
  // input := "What iS plUs 4?"
  // input := "What iS "
  // input := "4?"
  input := "What iS 4 plus 5"
  fmt.Println(input)
  fmt.Println()
  tokens := lexer.NewLexer(input).Lex()
  fmt.Println(tokens)
  parser := parser.NewParser(tokens)
  err := parser.Parse()
  if err != nil {
    log.Fatal(err)
  }
}


