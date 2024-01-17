package main

import (
	"log"
	"net/http"
)

// Handler for the home ("/") route
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/".
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	_, err := w.Write([]byte("Hello, SnippetBox!"))
	if err != nil {
		return
	}
}

// Handler for the getSnippet ("/snippet/get") route
func getSnippet(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Get a specific snippet!"))
	if err != nil {
		return
	}
}

// Handler for the createSnippet ("/snippet/create") route
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, `Method Not Allowed`, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte(`{"Message": "Create a new Snippet"}`))
	if err != nil {
		return
	}
}

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
