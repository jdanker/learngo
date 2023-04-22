package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	 Title string `json:"title"`
	 Author string `json:"author"`
	 Pages int `json:"pages"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 

	book := Book{Title: "The Alchemist", Author: "Paulo Coelho", Pages: 197}

	json.NewEncoder(w).Encode(book) 
}


func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("COntent-Type", "text/html")

	w.Write([]byte("<h1 color=steelblue> Hello World </h1>"))
}

func main() {
	http.HandleFunc("/hello", hello) // set endpoint for hello
	http.HandleFunc("/books", getBooks) // set endpoint for books

	log.Fatal(http.ListenAndServe(":8901", nil))



}