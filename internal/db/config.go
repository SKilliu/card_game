package db

import "github.com/caarlos0/env"

var configuration Configuration

type Configuration struct {
	Name     string `env:"CARD_GAME_API_DB_NAME,required"`
	Host     string `env:"CARD_GAME_API_DB_HOST,required"`
	Port     int    `env:"CARD_GAME_API_DB_PORT,required"`
	User     string `env:"CARD_GAME_API_DB_USER,required"`
	Password string `env:"CARD_GAME_API_DB_PASSWORD,required"`
	SSL      string `env:"CARD_GAME_API_DB_SSL"`
}

func loadConfigFromEnvs() {

	if err := env.Parse(&configuration); err != nil {
		logger.WithError(err).Error("failed to get db data from env")
		panic(err)
	}

	return
}
