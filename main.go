package main

import (
	"fmt"
	"log"

	"github.com/ivan-asdf/simple-math/lexer"
	"github.com/ivan-asdf/simple-math/parser"
)

func main() {
  input := "What iS 5 plUs 4?"
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


