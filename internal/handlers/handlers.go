package handlers

import (
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
	"log"
)

// RouteMessage routes the incoming message to the correct handler
func RouteMessage(client *telegram.Client, chatID int64, text string) {
	switch text {
	case "/start":
		StartHandler(client, chatID)
	case "/help":
		HelpHandler(client, chatID)
	case "/status":
		HandleStatusCommand(client, chatID)
	default:
		DefaultHandler(client, chatID, text)
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
