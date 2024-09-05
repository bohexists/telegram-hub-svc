package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/bohexists/telegram-hub-svc/internal/handlers"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize a new client with the bot token from environment variable
	client := telegram.NewClient(os.Getenv("TELEGRAM_BOT_TOKEN"))

	// Start the bot and handle updates
	handlers.HandleBotUpdates(client)

	// Keep the program running
	select {}
}
