package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: *s}
}

// type Endpoint string

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

	var errorText string
	if err != nil {
		errorText = err.Error()
	}
	evalResponse := EvaluateResponse{Result: result, Error: errorText}

	c.JSON(http.StatusOK, evalResponse)
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
