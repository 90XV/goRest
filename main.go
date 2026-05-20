package main

import (
	"log"
	"net/http"
)

func main() {
	// Set up routes
	http.HandleFunc("/posts", getPosts)
	http.HandleFunc("/posts/", getPost) // This will handle both GET and other methods for specific posts

	// For demonstration, we'll create separate handlers for different methods
	// In a real app, you'd use a router like gorilla/mux
	go func() {
		// POST /posts
		http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				createPost(w, r)
			} else {
				getPosts(w, r)
			}
		})

		// PUT /posts/{id}
		http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPut {
				updatePost(w, r)
			} else if r.Method == http.MethodDelete {
				deletePost(w, r)
			} else {
				getPost(w, r)
			}
		})
	}()

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
