package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/osrs", serveOSRS)
	http.Handle("/osrs/", http.StripPrefix("/osrs/", http.FileServer(http.Dir("."))))

	log.Println("Server starting on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func serveOSRS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "osrs.html")
}
