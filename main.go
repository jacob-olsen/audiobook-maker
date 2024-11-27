package main

import (
	metahandler "audiobook-maker/metaHandler"
	"net/http"
)

func main() {
	metahandler.MakeTabels()
}

func LoadUrlPaths() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
}
