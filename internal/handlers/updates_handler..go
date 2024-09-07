package handlers

import (
	"log"
	"time"

	"github.com/bohexists/telegram-hub-svc/internal/binance"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// HandleBotUpdates processes updates from the Telegram bot
func HandleBotUpdates(client *telegram.Client) {
	// Dictionary for keeping track of active users
	activeUsers := make(map[int64]struct{})

	// Periodically get updates from the Telegram bot
	go func() {
		for {
			updates, err := client.GetUpdates(0)
			if err != nil {
				log.Fatalf("Error getting updates: %v", err)
			}

			//Check if there are any updates
			for _, update := range updates {
				if update.Message != nil {
					chatID := update.Message.Chat.ID
					user := update.Message.From // Extracting user info
					activeUsers[chatID] = struct{}{}
					log.Printf("Chat ID detected: %d", chatID)
					RouteMessage(client, chatID, update.Message.Text, user) // Passing user info
				}
			}

			// Wait for 10 seconds before checking again
			time.Sleep(10 * time.Second)
		}
	}()

	// Periodically check prices and send notifications
	go func() {
		for {
			for chatID := range activeUsers {
				binance.CheckPricesAndNotify(client, chatID)
			}
			// Wait for 10 seconds before checking again
			time.Sleep(10 * time.Second)
		}
	}()
}
