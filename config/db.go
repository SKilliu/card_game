package config

import (
	"fmt"

	"github.com/SKilliu/card_game/db"
	"github.com/caarlos0/env"
)

type DB struct {
	Host     string `env:"CARD_GAME_API_DB_HOST,required"`
	Port     int64  `env:"CARD_GAME_API_DB_PORT,required"`
	User     string `env:"CARD_GAME_API_DB_USER,required"`
	Password string `env:"CARD_GAME_API_DB_PASSWORD,required"`
	Name     string `env:"CARD_GAME_API_DB_NAME,required"`
	SSL      string `env:"CARD_GAME_API_DB_SSL" envDefault:"disable"`
}

func (c *ConfigImpl) DB() db.QInterface {
	if c.db != nil {
		return c.db
	}

	c.Lock()
	defer c.Unlock()

	dbConf := &DB{}
	if err := env.Parse(dbConf); err != nil {
		panic(err)
	}

	conn, err := db.New(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Name, dbConf.SSL))
	if err != nil {
		panic(err)
	}

	c.db = conn

	return c.db
}
