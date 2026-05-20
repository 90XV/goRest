package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// GET /posts - Get all posts
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// GET /posts/{id} - Get a specific post
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract ID from URL
	idStr := r.URL.Path[len("/posts/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find the post
	for _, post := range posts {
		if post.ID == id {
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

// POST /posts - Create a new post
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newPost Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Assign a new ID (simple increment)
	newPost.ID = len(posts) + 1
	posts = append(posts, newPost)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPost)
}

// PUT /posts/{id} - Update a post
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Path[len("/posts/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedPost Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedPost.ID = id

	// Find and update the post
	for i, post := range posts {
		if post.ID == id {
			posts[i] = updatedPost
			json.NewEncoder(w).Encode(updatedPost)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}

// DELETE /posts/{id} - Delete a post
func deletePost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/posts/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find and delete the post
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[:i], posts[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Post not found", http.StatusNotFound)
}
