package main

import (
	metahandler "audiobook-maker/metaHandler"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var err error
	metahandler.MakeTabels()

	LoadUrlPaths()

	fmt.Println("all system online")
	err = http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func LoadUrlPaths() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("index")
		temp, err := template.ParseFiles("html/bookList.html", "html/templet.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, "jacob")
	})
}
