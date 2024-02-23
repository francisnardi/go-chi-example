package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// Create a new Chi router
	r := chi.NewRouter()

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, welcome to the homepage!"))
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the about page!"))
	})

	r.Get("/contact", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You can contact us at contact@example.com"))
	})

	// Start the server
	http.ListenAndServe(":8080", r)
}
