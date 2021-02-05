package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetStartPage(c echo.Context) error {
	file, err := os.Open("./web/static/html/page.html")
	if err != nil {
		return err
	}
	defer file.Close()

	t, err := template.ParseFiles("./web/static/html/page.html")
	if err != nil {
		return err
	}

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(c.Response(), data)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
