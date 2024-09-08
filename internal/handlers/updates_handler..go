package handlers

import (
	"log"
	"time"

	"github.com/bohexists/telegram-hub-svc/internal/binance"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// HandleBotUpdates processes updates from the Telegram bot
func ProcessBotUpdates(client *telegram.Client) {
	go func() {
		for {
			updates, err := client.GetUpdates(0)
			if err != nil {
				log.Fatalf("Error getting updates: %v", err)
			}

			for _, update := range updates {
				if update.Message != nil {
					chatID := update.Message.Chat.ID
					user := update.Message.From
					RouteMessage(client, chatID, update.Message.Text, user)
				}
			}

			time.Sleep(10 * time.Second)
		}
	}()

	// Запуск функции проверки цен и отправки уведомлений
	go binance.CheckPricesAndNotify(client)
}
