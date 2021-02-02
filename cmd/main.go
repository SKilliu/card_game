package main

import (
	"github.com/SKilliu/card_game/internal/server"
	"github.com/sirupsen/logrus"
)

const pathToConfigFile = "./static/envs.yaml"

// @title Card Game API
// @version 1.0
// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {
	log := &logrus.Logger{}
	logger := logrus.NewEntry(log)

	server.Init(logger)
	server.Start()
}

//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//	file, err := os.Open("./static/html/page.html")
//	if err != nil {
//		return
//	}
//	defer file.Close()
//
//	t, err := template.ParseFiles("./static/html/page.html")
//	if err != nil {
//		return
//	}
//
//	//t, err := template.New("webpage").Parse(tmpl)
//	//check(err)
//
//	data := struct {
//		Title string
//		Items []string
//	}{
//		Title: "My page",
//		Items: []string{
//			"My photos",
//			"My blog",
//		},
//	}
//
//	err = t.Execute(w, data)
//	if err != nil {
//		return
//	}
//})
//
//game := models.NewGame(2, "x1")
//
//game.AddRound()
//
//fmt.Println(game)
//
//// server started
//log.Fatal(http.ListenAndServe(":8000", nil))
