package main

import (
	"dev-pipe-sandbox/operations"
	"log"
	"net/http"
	"strings"
)

func main() {

	repo := operations.GetRepo()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "serve/index.html")
	})

	http.HandleFunc("/source/branch/", func(w http.ResponseWriter, r *http.Request) {
		branchName := strings.TrimPrefix(r.URL.Path, "/source/branch/")
		if branchName == "" {
			http.Error(w, "Branch name is required", http.StatusBadRequest)
			return
		}
		operations.SourceFromBranch(w, repo, branchName)
	})

	http.HandleFunc("/branch/", func(w http.ResponseWriter, r *http.Request) {
		operations.HandleBranch(w, r, repo)
	})

	http.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		// Just echo back the code to be rendered in the middle display
		code := r.FormValue("code")
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(code))
	})

	http.HandleFunc("/merge", func(w http.ResponseWriter, r *http.Request) {
		// Stub for merge logic
		w.Write([]byte("<p style='color: #4caf50;'>Merged successfully (mock)!</p>"))
	})

	log.Println("Server running at http://localhost:3232")
	if err := http.ListenAndServe(":3232", nil); err != nil {
		log.Fatal(err)
	}
}
