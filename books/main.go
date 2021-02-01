package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", GetBooksHandler)
	http.HandleFunc("/save", SaveBooksHandler)

	http.ListenAndServe(":8080", nil)
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {

	videos := getBooks()

	videoBytes, err := json.Marshal(videos)

	if err != nil {
		panic(err)
	}

	w.Write(videoBytes)
}

func SaveBooksHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		var books []book
		err = json.Unmarshal(body, &books)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Bad request!")
		}

		saveBooks(books)

	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not Supported!")
	}
}
