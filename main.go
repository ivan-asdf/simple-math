package main

import (
	"fmt"
	"log"

	"github.com/ivan-asdf/simple-math/lexer"
	"github.com/ivan-asdf/simple-math/parser"
)

func main() {
  tokens := lexer.Lex("What is 5 plus 4?")
  fmt.Println(tokens)
  parser := parser.NewParser(tokens)
  err := parser.Parse()
  if err != nil {
    log.Fatal(err)
  }
}


