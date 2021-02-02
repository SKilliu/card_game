package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type QInterface interface {
	DB() *sql.DB
}

// DB wraps dbx interface.
type DB struct {
	db *sql.DB
}

// DB returns db client.
func (d DB) DB() *sql.DB {
	return d.db
}

// New connection opening.
func New(config string) (QInterface, error) {
	db, err := sql.Open("postgres", config)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db: db}, err
}
