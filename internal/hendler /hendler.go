package handlers

import (
	"log"
	"time"

	"github.com/bohexists/telegram-hub-svc/internal/binance"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// HandleBotUpdates управляет получением обновлений и отправкой уведомлений активным пользователям
func HandleBotUpdates(client *telegram.Client) {
	// Храним активных пользователей (chat IDs), которые взаимодействовали с ботом
	activeUsers := make(map[int64]struct{})

	// Периодически проверяем обновления от бота
	go func() {
		for {
			updates, err := client.GetUpdates(0)
			if err != nil {
				log.Fatalf("Error getting updates: %v", err)
			}

			// Проверяем, кто из пользователей взаимодействовал с ботом
			for _, update := range updates {
				if update.Message != nil {
					chatID := update.Message.Chat.ID
					activeUsers[chatID] = struct{}{}
					log.Printf("Chat ID detected: %d", chatID)
				}
			}

			// Ждем перед следующей проверкой обновлений
			time.Sleep(10 * time.Second)
		}
	}()

	// Периодически проверяем цены и отправляем уведомления активным пользователям
	go func() {
		for {
			for chatID := range activeUsers {
				binance.CheckPricesAndNotify(client, chatID)
			}
			// Ждем перед следующей проверкой цен
			time.Sleep(10 * time.Second)
		}
	}()
}
