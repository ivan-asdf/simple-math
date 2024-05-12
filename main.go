package main

import (
	"fmt"

	"github.com/ivan-asdf/simple-math/lexer"
)

func main() {
  input := "What iS 5 plUs 4 das, das?"
  fmt.Println(input)
  fmt.Println()
  tokens := lexer.NewLexer(input).Lex()
  fmt.Println(tokens)
  // fmt.Println(tokens)
  // parser := parser.NewParser(tokens)
  // err := parser.Parse()
  // if err != nil {
  //   log.Fatal(err)
  // }
}


