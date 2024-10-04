package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("/app"))
	http.Handle("/osrs/", http.StripPrefix("/osrs/", fs))

	http.HandleFunc("/osrs", serveOSRS)

	http.HandleFunc("/", handleRequest)

	log.Println("Server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/osrs/") {
		http.Redirect(w, r, "/osrs", http.StatusMovedPermanently)
		return
	}
	http.NotFound(w, r)
}

func serveOSRS(w http.ResponseWriter, r *http.Request) {
	htmlPath := "/app/osrs.html"
	if _, err := os.Stat(htmlPath); os.IsNotExist(err) {
		log.Printf("File not found: %s", htmlPath)
		http.NotFound(w, r)
		return
	}
	log.Printf("Serving file: %s", htmlPath)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, htmlPath)
}
