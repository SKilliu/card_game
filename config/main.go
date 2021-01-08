package config

import (
	"sync"

	"github.com/SKilliu/card_game/db"

	"github.com/sirupsen/logrus"
)

type Config interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	//Authentication() *Authentication
	DB() db.QInterface
	//S3() *s3.Client
}

// ConfigImpl is implementation of config interface
type ConfigImpl struct {
	sync.Mutex

	http *HTTP
	log  *logrus.Entry
	//jwt  *Authentication
	db db.QInterface
	//s3   *s3.Client
}

// New config creating
func New() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
