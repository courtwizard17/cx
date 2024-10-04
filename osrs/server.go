package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Printf("Server starting on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for: %s", r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/osrs")

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	log.Printf("Adjusted path: %s", path)

	if path == "/" {
		log.Printf("Serving osrs.html")
		http.ServeFile(w, r, "osrs.html")
		return
	}

	filePath := filepath.Join(".", path)
	log.Printf("Looking for file: %s", filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Printf("File not found: %s", filePath)
		http.NotFound(w, r)
		return
	}

	log.Printf("Serving file: %s", filePath)
	http.ServeFile(w, r, filePath)
}
