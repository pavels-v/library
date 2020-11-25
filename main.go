package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

//Books is a temporary dummy data storage being later replaced with database table
var Books []Book = []Book{
	Book{Author: "Robert C. Martin", Title: "Clean Code: A Handbook of Agile Software Craftsmanship", Description: "Even bad code can function. But if code isn’t clean, it can bring a development organization to its knees. Every year, countless hours and significant resources are lost because of poorly written code. But it doesn’t have to be that way."},
	Book{Author: "Eric Freeman, Elisabeth Robson, Bert Bates, Kathy Sierra", Title: "Head First Design Patterns", Description: "At any given moment, someone struggles with the same software design problems you have. And, chances are, someone else has already solved your problem. This edition of Head First Design Patterns—now updated for Java 8—shows you the tried-and-true, road-tested patterns used by developers to create functional, elegant, reusable, and flexible software. By the time you finish this book, you’ll be able to take advantage of the best design practices and experiences of those who have fought the beast of software design and triumphed."},
	Book{Author: "Marijn Haverbeke", Title: "Eloquent JavaScript", Description: "This is a book about JavaScript, programming, and the wonders of the digital."},
	Book{Author: "Brian W. Kernighan, Alan Donovan", Title: "The Go Programming Language", Description: "The Go Programming Language is the authoritative resource for any programmer who wants to learn Go. It shows how to write clear and idiomatic Go to solve real-world problems."},
}

var port = flag.Int("port", 8080, "A port for http server to listen")

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/books", getBooksHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getBooks")
	json.NewEncoder(w).Encode(Books)
}

// Book structure including JSON tags
type Book struct {
	Author      string `json:"Author"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
