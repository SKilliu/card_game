package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/SKilliu/card_game/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("./static/html/page.html")
		if err != nil {
			return
		}
		defer file.Close()

		t, err := template.ParseFiles("./static/html/page.html")
		if err != nil {
			return
		}

		//t, err := template.New("webpage").Parse(tmpl)
		//check(err)

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

		err = t.Execute(w, data)
		if err != nil {
			return
		}
	})

	game := models.NewGame(2, "x1")

	game.AddRound()

	fmt.Println(game)

	// server started
	log.Fatal(http.ListenAndServe(":8000", nil))
}
