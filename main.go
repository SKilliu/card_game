package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Card game API started"))
	})

	// server started
	log.Fatal(http.ListenAndServe(":8000", nil))
}
