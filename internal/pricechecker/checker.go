package pricechecker

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"

	"github.com/bohexists/telegram-hub-svc/configs"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

func CheckPricesAndNotify(client *telegram.Client) {
	binanceClient := binance.NewClient("", "") // Initialize the Binance client

	// Get the chat ID from environment variable
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	// Convert chatID to int64 as expected by the SendMessage method
	chatIDInt64, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		log.Fatalf("Invalid TELEGRAM_CHAT_ID: %v", err)
	}

	for {
		for _, alert := range configs.PriceAlerts {
			// Get the current price from Binance API
			ctx := context.Background() // Create a context
			prices, err := binanceClient.NewListPricesService().Symbol(alert.Symbol).Do(ctx)
			if err != nil {
				log.Printf("Error fetching price for %s: %v", alert.Symbol, err)
				continue
			}

			price, _ := strconv.ParseFloat(prices[0].Price, 64)
			fmt.Printf("Current price of %s: %f\n", alert.Symbol, price)

			// Check if the price is below or above the threshold
			if price <= alert.MinPrice || price >= alert.MaxPrice {
				message := fmt.Sprintf("Alert! %s price is now %f", alert.Symbol, price)
				err := client.SendMessage(chatIDInt64, message)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}
			}
		}

		// Wait for a specific amount of time before checking the prices again
		time.Sleep(10 * time.Second) // Example: wait for 10 minutes
	}
}
