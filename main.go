package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SKilliu/card_game/models"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Card game API started"))
	})

	game := models.NewGame(2, "x1")

	game.AddRound()

	fmt.Println(game)

	// server started
	log.Fatal(http.ListenAndServe(":8000", nil))
}
