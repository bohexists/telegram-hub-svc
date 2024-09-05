package handlers

import (
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
	"log"
)

// StartHandler processes the /start command
func StartHandler(client *telegram.Client, chatID int64) {
	message := "Hello! Welcome to our bot."
	err := client.SendMessage(chatID, message)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}
