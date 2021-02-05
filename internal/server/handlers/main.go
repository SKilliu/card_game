package handlers

import "github.com/sirupsen/logrus"

type Handler struct {
	log *logrus.Entry
}

func New(logger *logrus.Entry) *Handler {
	return &Handler{
		log: logger,
	}
}
