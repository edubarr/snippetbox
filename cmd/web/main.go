package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize a new ServeMux, and register all handlers to corresponding URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/get", getSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")

	// Start a new web server on ":4000" and use the ServeMux as handler.
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
