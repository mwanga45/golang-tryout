package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID        string   `json:"id,omitempty"`
	Title     string   `json:"title,omitempty`
	Author    string   `json:"Author,omitempty"`
	Publisher *Company `json:"Publisher,omitempty"`
}

type Company struct {
	Name    string `json: "name,omitempty"`
	Address string `json: "Address,omitempty"`
}

var books []Book

// for print all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// for select are book
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// for select createbooks

func CreateBook(w http.ResponseWriter, r *http.Request) {
	
    w.Header().Set("Content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}
func main() {

	Router := mux.NewRouter()

	books = append(books, Book{ID: "1", Title: "Book One", Author: "John Doe", Publisher: &Company{Name: "Publisher One", Address: "Address One"}})
	books = append(books, Book{ID: "2", Title: "Book Two", Author: "Jane Smith", Publisher: &Company{Name: "Publisher Two", Address: "Address Two"}})
	Router.HandleFunc("/books", getBooks).Methods("GET")
	Router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	Router.HandleFunc("/books", CreateBook).Methods("POST")
	Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/plain")
		fmt.Print(w, "hello devprazoo")
	})
   
	fmt.Println("Server is running at http://localhost:8000")
	http.ListenAndServe(":8000", Router)

}
