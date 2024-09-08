package handlers

import (
	"log"

	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// HelpHandler sends a help message to the user
func HelpHandler(client *telegram.Client, chatID int64) {
	message := `Here are the available commands:
/start - Start interacting with the bot
/status - Get the current status of monitored cryptocurrencies
/help - Show this help message`

	err := client.SendMessage(chatID, message)
	if err != nil {
		log.Printf("Error sending help message: %v", err)
	}
}
