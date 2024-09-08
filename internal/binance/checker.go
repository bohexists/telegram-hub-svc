package binance

import (
	"context"
	"fmt"
	"github.com/bohexists/telegram-hub-svc/configs"
	"log"
	"strconv"
	"time"

	"github.com/adshao/go-binance/v2"

	"github.com/bohexists/telegram-hub-svc/internal/storage"
	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// CheckPricesAndNotify checks the prices of the specified alerts and sends notifications if necessary
func CheckPricesAndNotify(client *telegram.Client, configs configs.Config) {

	binanceClient := binance.NewClient("", "") // Initialize the Binance client

	for {
		log.Println("Fetching alerts from database...")
		// Retrieve all users from the database

		alerts, err := storage.GetAllCryptoAlerts()
		if err != nil {
			log.Printf("Error fetching alerts: %v", err)
			time.Sleep(configs.PriceCheckInterval)
			continue
		}

		log.Printf("Fetched %d alerts", len(alerts))

		// Check the prices of each alert
		for _, alert := range alerts {
			log.Printf("Processing alert for %s", alert.Symbol)
			if !alert.Enabled {
				continue
			}
			ctx := context.Background()
			prices, err := binanceClient.NewListPricesService().Symbol(alert.Symbol).Do(ctx)
			if err != nil {
				log.Printf("Error retrieving prices for %s: %v", alert.Symbol, err)
				continue
			}
			// Check if the price is within the specified range
			price, err := strconv.ParseFloat(prices[0].Price, 64)
			if err != nil {
				log.Printf("Error retrieving prices for %s: %v", alert.Symbol, err)
				continue
			}

			log.Printf("Current price for %s is %f", alert.Symbol, price)

			// Send a notification if the price is outside the range
			if (alert.MinPrice != nil && price <= *alert.MinPrice) || (alert.MaxPrice != nil && price >= *alert.MaxPrice) {

				log.Printf("Alert triggered for %s at price %f", alert.Symbol, price)

				message := fmt.Sprintf("Alert! %s price is now %f", alert.Symbol, price)
				err := client.SendMessage(alert.ChatID, message) // Directly use ChatID stored in the alert
				if err != nil {
					log.Printf("Error sending message for %s: %v", alert.Symbol, err)
				}
			} else {
				log.Printf("No alert triggered for %s at price %f", alert.Symbol, price)
			}
		}
	}

	time.Sleep(configs.PriceCheckInterval) // Pause for 10 seconds before checking again
}
