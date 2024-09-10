package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if _, err := os.Stat(filepath.Join("docs", path)); os.IsNotExist(err) {
		// File doesn't exist, return custom message
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "This file doesn't exist.")
		return
	}

	http.ServeFile(w, r, filepath.Join("docs", path))
}
