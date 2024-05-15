package cli

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ivan-asdf/simple-math/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestMakeErrorsGetRequest(t *testing.T) {
	t.Parallel()

	endpoint := api.ErorrsEndpoint

	response := `{"endpoint": "/validate","expression": "is 7","type": "Syntax error","frequency": "2"}`
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == endpoint {
			fmt.Fprint(w, response)
		}
	}))
	defer testServer.Close()

	cli := NewClient(testServer.URL, endpoint)
	result, err := cli.makeErrorsGetRequest()

	assert.Nil(t, err)
	expected := `{
  "endpoint": "/validate",
  "expression": "is 7",
  "type": "Syntax error",
  "frequency": "2"
}`
	assert.Equal(t, expected, result)
}

func TestEvaluateMakePostRequest(t *testing.T) {
	t.Parallel()

	endpoint := api.EvaluateEndpoint

	response := `{"result": 5}`
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == endpoint {
			fmt.Fprint(w, response)
		}
	}))
	defer testServer.Close()

	cli := NewClient(testServer.URL, endpoint)
	result, err := cli.makePostRequest("What is 5?")

	assert.Nil(t, err)
	expected := `{
  "result": 5
}`
	assert.Equal(t, expected, result)
}

func TestValidateMakePostRequest(t *testing.T) {
	t.Parallel()

	endpoint := api.ValidateEndpoint

	response := `{"valid": true}`
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == endpoint {
			fmt.Fprint(w, response)
		}
	}))
	defer testServer.Close()

	cli := NewClient(testServer.URL, endpoint)
	result, err := cli.makePostRequest("What is 5?")

	assert.Nil(t, err)
	expected := `{
  "valid": true
}`
	assert.Equal(t, expected, result)
}

func TestRequestFailure(t *testing.T) {
	t.Parallel()

	cli := NewClient("invalid", api.ErorrsEndpoint)

	result, err := cli.makeErrorsGetRequest()
	assert.NotNil(t, err)
	assert.Contains(t, result, "Status Code:")
	assert.Contains(t, result, "RequestAttempt:")
}
