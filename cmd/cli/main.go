package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ivan-asdf/simple-math/internal/api"
	"github.com/ivan-asdf/simple-math/internal/cli"
)

const UsageString = `
Usage: ./cli [options]

Options:
  -mode/-m string
    Mode of operation: eval, validate, or error (default "eval")
`

func main() {
	var help bool
	flag.BoolVar(&help, "help", false, "Print usage explanation")
	flag.BoolVar(&help, "h", false, "")
	var mode string
	flag.StringVar(&mode, "mode", "eval", `Mode of operation: eval, validate or errors`)
	flag.StringVar(&mode, "m", "", "")
	flag.Parse()

	if help {
		fmt.Print(UsageString)
		return
	}

	var endpoint string
	switch mode {
	case "eval":
		endpoint = api.EvaluateEndpoint
	case "validate":
		endpoint = api.ValidateEndpoint
	case "errors":
		endpoint = api.ErorrsEndpoint
	default:
		fmt.Printf("Invalid mode: %s\n", mode)
		os.Exit(1)
	}
	fmt.Println(mode, endpoint)
	cc := cli.NewCliClient("localhost", "1234", endpoint)
	cc.Run()
}
