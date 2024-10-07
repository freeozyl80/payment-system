package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the homepage!")
	})

	r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")
		fmt.Fprintf(w, "User ID: %s\n", userID)
	})

	http.ListenAndServe(":3000", r)
}
