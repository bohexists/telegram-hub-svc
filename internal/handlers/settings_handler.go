package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/bohexists/telegram-hub-svc/pkg/telegram"
)

// SettingsHandler allows users to update their monitoring settings
func SettingsHandler(client *telegram.Client, chatID int64, text string) {
	// Parse the command arguments
	args := strings.Split(text, " ")

	if len(args) < 2 {
		message := "Please provide a cryptocurrency symbol to monitor. Example: /settings BTCUSDT"
		err := client.SendMessage(chatID, message)
		if err != nil {
			log.Printf("Error sending settings message: %v", err)
		}
		return
	}

	symbol := args[1] // Retrieve the symbol from the command arguments

	// Check if the symbol is valid
	message := fmt.Sprintf("You are now monitoring %s. Use the /status command to view your current settings.", symbol)
	err := client.SendMessage(chatID, message)
	if err != nil {
		log.Printf("Error sending settings message: %v", err)
	}
}
