package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/ivan-asdf/simple-math/internal/api"
)

type Client struct {
	client   *resty.Client
	url      string
	endpoint string
}

func NewClient(url string, endpoint string) *Client {
	return &Client{
		client:   resty.New(),
		url:      url,
		endpoint: endpoint,
	}
}

func (c *Client) requestURL() string {
	return c.url + c.endpoint
}

func (c *Client) printCliPrompt() {
	switch c.endpoint {
	case api.EvaluateEndpoint:
		fmt.Print("evaluate > ")
	case api.ValidateEndpoint:
		fmt.Print("validate > ")
	default:
	}
}

func formatResponse(body []byte) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "  ")
	if err != nil {
		return fmt.Errorf("%s\n\n Failed to parse json body: %w", body, err).Error()
	}
	return prettyJSON.String()
}

func (c *Client) makePostRequest(input string) (string, error) {
	request := api.Request{Expression: input}

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(c.requestURL())
	if err != nil {
		return detailedErrorInformation(resp), err
	}

	return formatResponse(resp.Body()), nil
}

func (c *Client) makeErrorsGetRequest() (string, error) {
	resp, err := c.client.R().
		Get(c.requestURL())
	if err != nil {
		return detailedErrorInformation(resp), err
	}

	return formatResponse(resp.Body()), nil
}

func (c *Client) Run() {
	if c.endpoint == api.ErorrsEndpoint {
		result, err := c.makeErrorsGetRequest()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		fmt.Println(result)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter text (Ctrl+D to end):\n\n")
	c.printCliPrompt()
	for scanner.Scan() {
		input := scanner.Text()
		result, err := c.makePostRequest(input)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		fmt.Print(result, "\n\n")
		c.printCliPrompt()
	}
}
