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
	http.HandleFunc("/book/{bookId}", func(w http.ResponseWriter, r *http.Request) {
		BookId := r.PathValue("bookId")
		bookData := metahandler.GetBook(BookId)

		if bookData.Id == 0 {
			temp, err := template.ParseFiles("html/bookInfoErr.html", "html/templet.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, BookId)
		} else {
			temp, err := template.ParseFiles("html/bookInfo.html", "html/templet.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, bookData)
		}
	})
	http.HandleFunc("/book/{bookId}/page/{pageId}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("read page")
		temp, err := template.ParseFiles("html/bookList.html", "html/templet.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, "jacob")
	})
}
