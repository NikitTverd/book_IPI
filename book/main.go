package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Serve static files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// API routes
	router.HandleFunc("/api/books", GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/api/books", CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", DeleteBook).Methods("DELETE")

	// Serve HTML templates
	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
