package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerValidate(t *testing.T) {
	s := NewService()
	h := NewHandler(&s)

	validateRequest := Request{Expression: "What is 3?"}
	expected := ValidateResponse{Valid: true}
	response, httpCode := h.validate(validateRequest)
	assert.Equal(t, 200, httpCode)
	assert.Equal(t, expected, response)

	validateRequest = Request{Expression: "What is 3 plus plus 3?"}
	expected = ValidateResponse{Valid: false, Reason: "Syntax error: expected number, got \"plus\" at 15-19"}
	response, httpCode = h.validate(validateRequest)
	assert.Equal(t, 200, httpCode)
	assert.Equal(t, expected, response)
}

func TestHandlerEvaluate(t *testing.T) {
	s := NewService()
	h := NewHandler(&s)

	evalRequest := Request{Expression: "What is 3?"}
	expected := EvaluateResponse{Result: 3}
	response, httpCode := h.evaluate(evalRequest)
	assert.Equal(t, 200, httpCode)
	assert.Equal(t, expected, response)

	evalRequest = Request{Expression: "What is 3 plus plus 3?"}
	expected = EvaluateResponse{Result: 0, Error: "Syntax error: expected number, got \"plus\" at 15-19"}
	response, httpCode = h.evaluate(evalRequest)
	assert.Equal(t, 200, httpCode)
	assert.Equal(t, expected, response)

	evalRequest = Request{Expression: "What is 3 divided by 0?"}
	expected = EvaluateResponse{Result: 0, Error: "division by zero"}
	response, httpCode = h.evaluate(evalRequest)
	assert.Equal(t, 200, httpCode)
	assert.Equal(t, expected, response)
}
