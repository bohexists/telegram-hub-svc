package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiURL = "https://api.telegram.org/bot"
)

type Client struct {
	Token  string
	Client *http.Client
}

// NewClient creates a new Telegram API client
func NewClient(token string) *Client {
	return &Client{
		Token:  token,
		Client: &http.Client{},
	}
}

// request sends an HTTP request to the Telegram API and returns the result
func (c *Client) request(method string, data interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s/%s", apiURL, c.Token, method)

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData := &bytes.Buffer{}
	_, err = responseData.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("telegram API returned error: %s", responseData.String())
	}

	return responseData.Bytes(), nil
}
