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
	host     string
	port     string
	endpoint string
}

func NewCliClient(host string, port string, endpoint string) *Client {
	return &Client{
		client:   resty.New(),
		host:     host,
		port:     port,
		endpoint: endpoint,
	}
}

func (c *Client) requestURL() string {
	return "http://" + c.host + ":" + c.port + c.endpoint
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
		return err.Error()
	}
	return prettyJSON.String()
}

func (c *Client) makePostRequest(input string) string {
	request := api.Request{Expression: input}

	resp, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(c.requestURL())
	if err != nil {
		return detailedErrorInformation(err, resp)
	}

	return formatResponse(resp.Body())
}

func (c *Client) makeGetRequest() string {
	resp, err := c.client.R().
		Get(c.requestURL())
	if err != nil {
		return detailedErrorInformation(err, resp)
	}

	return formatResponse(resp.Body())
}

func (c *Client) Run() {
	if c.endpoint == api.ErorrsEndpoint {
		result := c.makeGetRequest()
		fmt.Println(result)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter text (Ctrl+D to end):\n\n")
	c.printCliPrompt()
	for scanner.Scan() {
		result := c.makePostRequest(scanner.Text())
		fmt.Print(result, "\n\n")
		c.printCliPrompt()
	}
}
