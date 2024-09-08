package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/bohexists/telegram-hub-svc/db"
	"github.com/bohexists/telegram-hub-svc/internal/handlers"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"

	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db.InitDB()

	// Initialize a new client with the bot token from environment variable
	client := telegram.NewClient(os.Getenv("TELEGRAM_BOT_TOKEN"))

	// Start the bot and handle updates
	handlers.ProcessBotUpdates(client)

	// Keep the program running
	select {}
}
