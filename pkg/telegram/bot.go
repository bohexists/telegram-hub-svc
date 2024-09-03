package telegram

import (
	"encoding/json"
	"fmt"
)

// SendMessageData structure for sending messages
type SendMessageData struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// SendMessage sends a message to the specified chat
func (c *Client) SendMessage(chatID int64, text string) error {
	data := SendMessageData{
		ChatID: chatID,
		Text:   text,
	}

	_, err := c.request("sendMessage", data)
	if err != nil {
		return err
	}

	return nil
}

// GetUpdatesData structure for getting updates
type GetUpdatesData struct {
	Offset  int64 `json:"offset"`
	Limit   int   `json:"limit"`
	Timeout int   `json:"timeout"`
}

// Update represents an update from Telegram
type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message"`
}

// GetUpdates retrieves updates (messages/commands)
func (c *Client) GetUpdates(offset int64) ([]Update, error) {
	data := GetUpdatesData{
		Offset:  offset,
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

	return updatesResponse.Result, nil
}
