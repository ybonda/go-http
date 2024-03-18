package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func getBooks() (books []book) {

	fileBytes, err := os.ReadFile("./books.json")

	if err != nil {
		panic(err) // panicing
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

	err = os.WriteFile("./new-books.json", bookBytes, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Book saved!")

}
