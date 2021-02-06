package handlers

import (
	"github.com/SKilliu/cardgame/internal/db"
	"github.com/gocraft/dbr/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log *logrus.Entry
	db  *dbr.Connection
}

func New(logger *logrus.Entry) *Handler {
	return &Handler{
		log: logger,
		db:  db.Connection(),
	}
}
