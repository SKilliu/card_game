package s3

import "github.com/caarlos0/env"

var configuration Configuration

type Configuration struct {
	AccessKey string `env:"CARD_GAME_S3_ACCESS_KEY"`
	SecretKey string `env:"CARD_GAME_S3_SECRET_KEY"`
	Url       string `env:"CARD_GAME_S3_URL"`
	Bucket    string `env:"CARD_GAME_S3_BUCKET"`
}

func loadConfigFromEnvs() {

	if err := env.Parse(&configuration); err != nil {
		logger.WithError(err).Error("failed to load s3 config from envs")
		panic(err)
	}

	return
}
