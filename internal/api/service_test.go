package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceEvaluate(t *testing.T) {
	s := NewService()
	result, err := s.Evaluate("What is 7?")
	assert.Nil(t, err)
	assert.Equal(t, 7, result)

	result, err = s.Evaluate("What is 7 multiplied by 3?")
	assert.Nil(t, err)
	assert.Equal(t, 21, result)
}

func TestServiceEvaluateFail(t *testing.T) {
	s := NewService()
	endpoint := "/evaluate"

	input := "is 7"
	expected_err := "Syntax error: expected \"What is\" keyword"
	result, err := s.Evaluate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 0, result)
	result, err = s.Evaluate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 0, result)
	assert.Equal(t, 2, s.errorsLog[ErrorLogKey{endpoint, input, ErrorTypeSyntaxError}])

	input = "what is 4 cubed"
	expected_err = "Unsupported error: Unsupported operation \"cubed\" at 10-15"
	result, err = s.Evaluate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 0, result)
	result, err = s.Evaluate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 0, result)
	assert.Equal(t, 2, s.errorsLog[ErrorLogKey{endpoint, input, ErrorTypeUnsupportedOperation}])

	input = "what is the weather today?"
	expected_err = "non-math question(no numbers found in question)"
	result, err = s.Evaluate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 0, result)
	result, err = s.Evaluate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 0, result)
	assert.Equal(t, 2, s.errorsLog[ErrorLogKey{endpoint, input, ErrorTypeNonMathQuestion}])

	errors := []map[string]string{
		{
			"endpoint":   endpoint,
			"expression": "is 7",
			"type":       ErrorTypeSyntaxError,
			"frequency":  "2",
		},
		{
			"endpoint":   endpoint,
			"expression": "what is 4 cubed",
			"type":       ErrorTypeUnsupportedOperation,
			"frequency":  "2",
		},
		{
			"endpoint":   endpoint,
			"expression": "what is the weather today?",
			"type":       ErrorTypeNonMathQuestion,
			"frequency":  "2",
		},
	}
	assert.ElementsMatch(t, errors, s.Errors())
}

func TestServiceValidate(t *testing.T) {
	s := NewService()
	_, err := s.Evaluate("What is 7?")
	assert.Nil(t, err)

	_, err = s.Evaluate("What is 7 multiplied by 3?")
	assert.Nil(t, err)
}

func TestServiceValidateFail(t *testing.T) {
	s := NewService()
	endpoint := "/validate"

	input := "is 7"
	expected_err := "Syntax error: expected \"What is\" keyword"
	err := s.Validate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	err = s.Validate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 2, s.errorsLog[ErrorLogKey{endpoint, input, ErrorTypeSyntaxError}])

	input = "what is 4 cubed"
	expected_err = "Unsupported error: Unsupported operation \"cubed\" at 10-15"
	err = s.Validate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	err = s.Validate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 2, s.errorsLog[ErrorLogKey{endpoint, input, ErrorTypeUnsupportedOperation}])

	input = "what is the weather today?"
	expected_err = "non-math question(no numbers found in question)"
	err = s.Validate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	err = s.Validate(input)
	assert.EqualErrorf(t, err, expected_err, "Errors should match")
	assert.Equal(t, 2, s.errorsLog[ErrorLogKey{endpoint, input, ErrorTypeNonMathQuestion}])

	errors := []map[string]string{
		{
			"endpoint":   endpoint,
			"expression": "is 7",
			"type":       ErrorTypeSyntaxError,
			"frequency":  "2",
		},
		{
			"endpoint":   endpoint,
			"expression": "what is 4 cubed",
			"type":       ErrorTypeUnsupportedOperation,
			"frequency":  "2",
		},
		{
			"endpoint":   endpoint,
			"expression": "what is the weather today?",
			"type":       ErrorTypeNonMathQuestion,
			"frequency":  "2",
		},
	}
	assert.ElementsMatch(t, errors, s.Errors())
}
