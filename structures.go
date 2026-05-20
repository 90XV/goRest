package main

// Post represents a blog post
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// In-memory database simulation
var posts = []Post{
	{ID: 1, Title: "Hello World", Content: "This is my first post", Author: "John"},
	{ID: 2, Title: "Learning Go", Content: "Go is awesome", Author: "Jane"},
}
