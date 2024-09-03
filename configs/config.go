package configs

type PriceAlert struct {
	Symbol   string  // Cryptocurrency pair symbol, e.g., "BTCUSDT"
	MinPrice float64 // Minimum price threshold for alert
	MaxPrice float64 // Maximum price threshold for alert
}

var PriceAlerts = []PriceAlert{
	{
		Symbol:   "BTCUSDT",
		MinPrice: 20000.0,
		MaxPrice: 50000.0,
	},
	{
		Symbol:   "ETHUSDT",
		MinPrice: 10000.0,
		MaxPrice: 1,
	},
}
