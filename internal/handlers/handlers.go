package handlers

import (
	"log"
	"strings"

	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// RouteMessage routes the incoming message to the correct handler
func RouteMessage(client *telegram.Client, chatID int64, text string, user *telegram.User) {
	switch text {
	case "/start":
		StartHandler(client, chatID, user) // Save user data
	case "/help":
		HelpHandler(client, chatID) // Send help message
	case "/status":
		HandleStatusCommand(client, chatID) // Send status message
	default:
		if strings.HasPrefix(text, "/settings") {
			SettingsHandler(client, chatID, text) // Update user settings
		} else {
			DefaultHandler(client, chatID, text)
		}
	}
}

// DefaultHandler responds to unknown commands
func DefaultHandler(client *telegram.Client, chatID int64, text string) {
	message := "Sorry, I don't understand this command. Try /help"
	err := client.SendMessage(chatID, message)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
