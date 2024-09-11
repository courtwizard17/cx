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
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Nothing here yet"))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/docs/") {
		path := r.URL.Path[len("/docs/"):]
		filePath := filepath.Join("docs", path)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, filePath)
		return
	}

	http.NotFound(w, r)
}
