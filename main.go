package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "serve/index.html")
	})

	log.Println("Server running at http://localhost:3232")
	if err := http.ListenAndServe(":3232", nil); err != nil {
		log.Fatal(err)
	}
}
