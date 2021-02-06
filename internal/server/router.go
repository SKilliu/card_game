package server

import (
	"net/http"

	"github.com/SKilliu/cardgame/internal/server/handlers"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(logger *logrus.Entry) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := handlers.New(logger)

	e.GET("/", healthz)
	e.GET("/start", handler.GetStartPage)
	e.POST("/user", handler.CreateUser)

	return e
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "Server successfully started")
}
