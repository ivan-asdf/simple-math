package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-asdf/simple-math/lexer"
	"github.com/ivan-asdf/simple-math/parser"
)

type Request struct {
	Expression string `json:"expression"`
}

type EvaluateResponse struct {
	Result int    `json:"result"`
	Error  string `json:"error,omitempty"`
}

func eval(input string) (int, error) {
	// input := "What iS plUs 4?"
	fmt.Println(input)
	fmt.Println()

	tokens := lexer.NewLexer(input).Lex()
	fmt.Println(tokens)

	parser := parser.NewParser(tokens)
	expr, err := parser.Parse()
	if err != nil {
		return 0, err
	}

	return expr.Evaluate()
}

func Evaluate(c *gin.Context) {
	var evalRequest Request
	if err := c.BindJSON(&evalRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	result, err := eval(evalRequest.Expression)

	evalResponse := EvaluateResponse{Result: result, Error: err.Error()}
	c.JSON(http.StatusOK, evalResponse)
}

type ValidateRespone struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

func validate(input string) error {
	// input := "What iS plUs 4?"
	fmt.Println(input)
	fmt.Println()

	tokens := lexer.NewLexer(input).Lex()
	fmt.Println(tokens)

	parser := parser.NewParser(tokens)
	_, err := parser.Parse()
	return err
}

func Validate(c *gin.Context) {
	var validateRequest Request
	if err := c.BindJSON(&validateRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err := validate(validateRequest.Expression)
	var validateResponse ValidateRespone
	if err != nil {
		validateResponse.Valid = false
		validateResponse.Reason = err.Error()
	} else {
		validateResponse.Valid = true
	}

	c.JSON(http.StatusOK, validateResponse)
}
