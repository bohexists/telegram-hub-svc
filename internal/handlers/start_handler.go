package handlers

import (
	"log"

	"github.com/bohexists/telegram-hub-svc/internal/storage"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// StartHandler processes the "/start" command, saves user data into the database.
func StartHandler(client *telegram.Client, chatID int64, user *telegram.User) {
	// Try to create a new user in the database
	created, err := storage.CreateUser(storage.User{
		ChatID:    chatID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	})
	if err != nil {
		// If there is an error, send a message back to the user
		client.SendMessage(chatID, "An error occurred while saving your data. Please try again later.")
		log.Printf("Error creating user: %v", err)
		return
	}
	if created {
		// If user was successfully created, send a welcome message
		client.SendMessage(chatID, "Welcome to the bot! Your data has been successfully saved.")
	} else {
		// If the user already exists, send a different message
		client.SendMessage(chatID, "Welcome back! You are already registered.")
	}
}
