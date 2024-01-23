package main

import "net/http"

// The routes() method returns a servemux containing our routes.
func (app *application) routes() *http.ServeMux {
	// Initialize a new ServeMux, and register all handlers to corresponding URL pattern.
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for all URL paths that start with "/static/".
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/get", app.getSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	return mux
}
