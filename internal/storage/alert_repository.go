package storage

import (
	"database/sql"
	"log"

	"github.com/bohexists/telegram-hub-svc/db"
)

type CryptoAlert struct {
	ChatID   int64
	Symbol   string
	MinPrice *float64
	MaxPrice *float64
	Enabled  bool
}

func GetAllCryptoAlerts() ([]CryptoAlert, error) {
	alerts := []CryptoAlert{}
	rows, err := db.GetDB().Query("SELECT chat_id, symbol, min_price, max_price, enabled FROM crypto_alerts WHERE enabled = true")
	if err != nil {
		log.Printf("Error getting crypto alerts: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var alert CryptoAlert
		var minPrice sql.NullFloat64
		var maxPrice sql.NullFloat64
		if err := rows.Scan(&alert.ChatID, &alert.Symbol, &minPrice, &maxPrice, &alert.Enabled); err != nil {
			log.Printf("Error scanning crypto alert: %v", err)
			continue
		}
		if minPrice.Valid {
			alert.MinPrice = &minPrice.Float64
		}
		if maxPrice.Valid {
			alert.MaxPrice = &maxPrice.Float64
		}
		alerts = append(alerts, alert)
	}

	return alerts, nil
}
