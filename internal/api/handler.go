package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-asdf/simple-math/internal/eval"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}

const InternalServerError = "internal server error"

const (
	EvaluateEndpoint = "/evaluate"
	ValidateEndpoint = "/validate"
	ErorrsEndpoint   = "/errors"
)

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST(EvaluateEndpoint, h.Evaluate)
	r.POST(ValidateEndpoint, h.Validate)
	r.GET(ErorrsEndpoint, h.Errors)
}

type Request struct {
	Expression string `json:"expression"`
}

type EvaluateResponse struct {
	Result int    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

type ValidateResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

func (h *Handler) Evaluate(c *gin.Context) {
	var evalRequest Request
	if err := c.BindJSON(&evalRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	evalResponse, httpCode := h.evaluate(evalRequest)
	c.JSON(httpCode, evalResponse)
}

func (h *Handler) evaluate(evalRequest Request) (EvaluateResponse, int) {
	result, err := h.s.Evaluate(evalRequest.Expression)

	var evalResponse EvaluateResponse
	httpCode := http.StatusOK
	if err != nil {
		switch {
		case errors.Is(err, eval.ErrEvalInvalidExpr):
			httpCode = http.StatusInternalServerError
			evalResponse.Error = InternalServerError
		case errors.Is(err, eval.ErrEvalDivisionByZero):
			fallthrough
		default:
			evalResponse.Error = err.Error()
		}

		return evalResponse, httpCode
	}

	evalResponse.Result = result
	return evalResponse, httpCode
}

func (h *Handler) Validate(c *gin.Context) {
	var validateRequest Request
	if err := c.BindJSON(&validateRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	response, httpCode := h.validate(validateRequest)
	c.JSON(httpCode, response)
}

func (h *Handler) validate(validateRequest Request) (ValidateResponse, int) {
	err := h.s.Validate(validateRequest.Expression)

	var validateResponse ValidateResponse
	if err != nil {
		validateResponse.Valid = false
		validateResponse.Reason = err.Error()
	} else {
		validateResponse.Valid = true
	}

	return validateResponse, 200
}

func (h *Handler) Errors(c *gin.Context) {
	c.JSON(http.StatusOK, h.s.Errors())
}
