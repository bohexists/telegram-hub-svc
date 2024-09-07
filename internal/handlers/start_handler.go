package handlers

import (
	"log"

	"github.com/bohexists/telegram-hub-svc/internal/storage"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// StartHandler processes the "/start" command, saves user data into the database.
func StartHandler(client *telegram.Client, chatID int64, user *telegram.User) {
	// Create a new User instance to be saved in the database
	dbUser := storage.User{
		ChatID:    chatID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}

	// Save the user to the database
	err := storage.CreateUser(dbUser)
	if err != nil {
		log.Printf("Failed to save user: %v", err)
		return
	}

	// Respond to the user in Telegram
	message := "Welcome to the bot!"
	err = client.SendMessage(chatID, message)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
