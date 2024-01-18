package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	_, err = fmt.Fprintf(w, "Get snippet with ID %d", id)
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
