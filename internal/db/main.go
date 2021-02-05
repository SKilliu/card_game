package db

import (
	"fmt"

	"github.com/gocraft/dbr/v2"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var dbConn QInterface

type QInterface interface {
	DB() *dbr.Connection
}

// DB wraps interface.
type DB struct {
	db *dbr.Connection
}

// DB returns db client.
func (d DB) DB() *dbr.Connection {
	return d.db
}

// New connection opening.
func New(config string) (QInterface, error) {
	db, err := dbr.Open("postgres", config, nil)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.WithError(err).Error("failed to ping db")
		return nil, err
	}

	logger.Info("db connection successfully created")

	return &DB{db: db}, err
}

func Init(log *logrus.Entry) {
	setLogger(log)

	loadConfigFromEnvs()

	conn, err := New(configuration.Info())
	if err != nil {
		log.WithError(err).Error("failed to setup db")
		panic(err)
	}

	dbConn = conn
}

func (d Configuration) Info() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSL,
	)
}

func Connection() QInterface {
	return dbConn
}
