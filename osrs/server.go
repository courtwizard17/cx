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
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		altPath := strings.TrimPrefix(path, "/osrs")
		altFilePath := filepath.Join(".", altPath)
		log.Printf("File not found, trying alternative path: %s", altFilePath)
		if _, err := os.Stat(altFilePath); os.IsNotExist(err) {
			log.Printf("File not found: %s", altFilePath)
			http.NotFound(w, r)
			return
		}
		filePath = altFilePath
	}

	log.Printf("Serving file: %s", filePath)
	http.ServeFile(w, r, filePath)
}
