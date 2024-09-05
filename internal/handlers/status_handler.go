package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/adshao/go-binance/v2"
	"github.com/bohexists/telegram-hub-svc/configs"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// HandleStatusCommand processes the /status command
func HandleStatusCommand(client *telegram.Client, chatID int64) {
	binanceClient := binance.NewClient("", "") // Initialize the Binance client

	message := "Бот следит за следующими криптовалютами:\n"
	for _, alert := range configs.PriceAlerts {
		// Retrieve the current price from Binance
		ctx := context.Background()
		prices, err := binanceClient.NewListPricesService().Symbol(alert.Symbol).Do(ctx)
		if err != nil {
			log.Printf("Ошибка при получении цены для %s: %v", alert.Symbol, err)
			message += fmt.Sprintf("%s: ошибка при получении цены\n", alert.Symbol)
			continue
		}

		price, _ := strconv.ParseFloat(prices[0].Price, 64)
		message += fmt.Sprintf("%s: %f\n", alert.Symbol, price)
	}

	// Send the message to the chat
	err := client.SendMessage(chatID, message)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
	}
}
