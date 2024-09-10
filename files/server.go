package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
