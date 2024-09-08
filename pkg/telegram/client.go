package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const apiURL = "https://api.telegram.org/bot"

type Client struct {
	Token        string
	HttpClient   *http.Client
	LastUpdateID int64
}

func NewClient(token string) *Client {
	return &Client{
		Token:      token,
		HttpClient: &http.Client{},
	}
}

func (c *Client) SendMessage(chatID int64, text string) error {
	data := SendMessageData{
		ChatID: chatID,
		Text:   text,
	}
	_, err := c.request("sendMessage", data)
	if err != nil {
		log.Printf("Error sending message: %v", err)
		return err
	}
	return nil
}

func (c *Client) GetUpdates() ([]Update, error) {
	data := GetUpdatesData{
		Offset:  c.LastUpdateID + 1,
		Limit:   100,
		Timeout: 30,
	}
	respData, err := c.request("getUpdates", data)
	if err != nil {
		return nil, err
	}

	var updatesResponse struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	err = json.Unmarshal(respData, &updatesResponse)
	if err != nil {
		return nil, err
	}

	if !updatesResponse.Ok {
		return nil, fmt.Errorf("failed to get updates from Telegram API")
	}

	if len(updatesResponse.Result) > 0 {
		c.LastUpdateID = updatesResponse.Result[len(updatesResponse.Result)-1].UpdateID
	}

	return updatesResponse.Result, nil
}

func (c *Client) HandleUpdates() {
	for {
		updates, err := c.GetUpdates()
		if err != nil {
			log.Printf("Error getting updates: %v", err)
			continue
		}

		for _, update := range updates {
			if update.Message != nil {
				// Process update using your message router or handler
			}
		}
	}
}

func (c *Client) request(method string, data interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s%s/%s", apiURL, c.Token, method)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBuffer := &bytes.Buffer{}
	_, err = responseBuffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("telegram API returned error: %s", responseBuffer.String())
	}

	// Correctly retrieve the bytes from the buffer
	return responseBuffer.Bytes(), nil
}
