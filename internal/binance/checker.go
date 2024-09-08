package binance

import (
	"context"
	"fmt"
	"github.com/bohexists/telegram-hub-svc/internal/storage"
	"log"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"

	"github.com/bohexists/telegram-hub-svc/configs"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// CheckPricesAndNotify checks the prices of the specified alerts and sends notifications if necessary
func CheckPricesAndNotify(client *telegram.Client) {

	binanceClient := binance.NewClient("", "") // Initialize the Binance client

	for {
		// Retrieve all users from the database
		users, err := storage.GetAllUsers()
		if err != nil {
			log.Printf("Error retrieving users: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}
		// Check the prices of each alert
		for _, alert := range configs.PriceAlerts {
			ctx := context.Background()
			prices, err := binanceClient.NewListPricesService().Symbol(alert.Symbol).Do(ctx)
			if err != nil {
				log.Printf("Error retrieving prices for %s: %v", alert.Symbol, err)
				continue
			}
			// Check if the price is within the specified range
			price, _ := strconv.ParseFloat(prices[0].Price, 64)

			// Send a notification if the price is outside the range
			for _, user := range users {
				if price <= alert.MinPrice || price >= alert.MaxPrice {
					message := fmt.Sprintf("Alert! %s price is now %f", alert.Symbol, price)
					client.SendMessage(user.ChatID, message)
				}
			}
		}

		time.Sleep(15 * time.Second) // Pause for 10 seconds before checking again
	}
}
