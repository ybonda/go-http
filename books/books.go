package main

import (
	"encoding/json"
	"io/ioutil"
)

type book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func getBooks() (books []book) {

	fileBytes, err := ioutil.ReadFile("./books.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &books)

	if err != nil {
		panic(err)
	}

	return books
}

func saveBooks(books []book) {
	bookBytes, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./new-books.json", bookBytes, 0644)
	if err != nil {
		panic(err)
	}

}
