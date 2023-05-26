package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			tmpl := template.Must(template.ParseFiles("templates/book.html"))
			tmpl.Execute(w, item)
			return
		}
	}

	http.NotFound(w, r)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = fmt.Sprintf("%d", len(books)+1)
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
