package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ivan-asdf/simple-math/cmd/common"
	"github.com/ivan-asdf/simple-math/internal/api"
	"github.com/ivan-asdf/simple-math/internal/cli"
)

const UsageString = `
Usage: ./cli [options]

Options:
  -mode/-m string
    Mode of operation: eval, validate, or error (default "eval")
  -port/-p :{port} (default ":55555")
`

func main() {
	var help bool
	flag.BoolVar(&help, "help", false, "Print usage explanation")
	flag.BoolVar(&help, "h", false, "")

	var mode string
	flag.StringVar(&mode, "mode", "eval", `Mode of operation: eval, validate or errors`)
	flag.StringVar(&mode, "m", "eval", "")

	var port string
	flag.StringVar(&port, "port", common.DefaultPort, "To which port to connect")
	flag.StringVar(&port, "p", common.DefaultPort, "")

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
	cc := cli.NewCliClient("http://localhost"+port, endpoint)
	cc.Run()
}
