package main

import (
	"dev-pipe-sandbox/operations"
	"log"
	"net/http"
)

func main() {

	repo := operations.GetRepo()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "serve/index.html")
	})

	http.HandleFunc("/source", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "serve/source.html")
	})

	log.Println("Server running at http://localhost:3232")
	if err := http.ListenAndServe(":3232", nil); err != nil {
		log.Fatal(err)
	}
}
