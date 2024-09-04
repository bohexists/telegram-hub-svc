package binance

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"

	"github.com/bohexists/telegram-hub-svc/configs"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// CheckPricesAndNotify checks the prices of the specified alerts and sends notifications if necessary
func CheckPricesAndNotify(client *telegram.Client, chatID int64) {
	binanceClient := binance.NewClient("", "") // Initialize the Binance client

	for {
		for _, alert := range configs.PriceAlerts {
			//Retrieve the current price from Binance
			ctx := context.Background()
			prices, err := binanceClient.NewListPricesService().Symbol(alert.Symbol).Do(ctx)
			if err != nil {
				log.Printf("Ошибка при получении цены для %s: %v", alert.Symbol, err)
				continue
			}

			price, _ := strconv.ParseFloat(prices[0].Price, 64)
			fmt.Printf("Текущая цена %s: %f\n", alert.Symbol, price)

			// Check if the price is within the specified range
			if price <= alert.MinPrice || price >= alert.MaxPrice {
				message := fmt.Sprintf("Alert! %s price is now %f", alert.Symbol, price)
				// Use the Telegram client to send the message
				err := client.SendMessage(chatID, message)
				if err != nil {
					log.Printf("Ошибка при отправке сообщения: %v", err)
				}
			}
		}

		// Wait for 10 seconds before checking again
		time.Sleep(10 * time.Second) // Get prices every 10 seconds
	}
}
