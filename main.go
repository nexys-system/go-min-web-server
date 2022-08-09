package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string
	Desc    string
	Content string
}

var Articles = []Article{
	{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.HandleFunc("/articles", returnAllArticles)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "main")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
