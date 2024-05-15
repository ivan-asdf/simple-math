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

func (h *Handler) Evaluate(c *gin.Context) {
	var evalRequest Request
	if err := c.BindJSON(&evalRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

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

		c.JSON(httpCode, evalResponse)
		return
	}

	evalResponse.Result = result
	c.JSON(httpCode, evalResponse)
}

type ValidateRespone struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

func (h *Handler) Validate(c *gin.Context) {
	var validateRequest Request
	if err := c.BindJSON(&validateRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	err := h.s.Validate(validateRequest.Expression)

	var validateResponse ValidateRespone
	if err != nil {
		validateResponse.Valid = false
		validateResponse.Reason = err.Error()
	} else {
		validateResponse.Valid = true
	}

	c.JSON(http.StatusOK, validateResponse)
}

func (h *Handler) Errors(c *gin.Context) {
	c.JSON(http.StatusOK, h.s.Errors())
}
