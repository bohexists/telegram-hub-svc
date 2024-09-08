package configs

import "time"

type Config struct {
	PriceCheckInterval time.Duration
}

var DefaultConfig = Config{
	PriceCheckInterval: 15 * time.Second, // Задаем стандартное значение
}
