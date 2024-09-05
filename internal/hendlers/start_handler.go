package handlers

import (
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
	"log"
)

// StartHandler отвечает на команду /start
func StartHandler(client *telegram.Client, chatID int64) {
	// Приветственное сообщение
	message := "Hello! Welcome to our bot."

	// Отправка сообщения пользователю
	err := client.SendMessage(chatID, message)
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
}
