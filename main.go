package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// define a blog post struct
type SocMedPost struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// FOR TESTING. In memory DB
var posts = []SocMedPost{
	{ID: 1, Title: "Hello World", Content: "This is my first post", Author: "John Doe"},
	{ID: 2, Title: "WHAT IS UP", Content: "Yoooooooo", Author: "Jane Doe"},
}

// Create HTTP Handlers | Using Gin

// GET /posts Gets ALL posts
func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

// POST code to add a new USER on the social platform
func postPosts(c *gin.Context) {
	var newSMP SocMedPost

	// Call BindJSON to bind the received JSON to newSMP
	if err := c.BindJSON(&newSMP); err != nil {
		return
	}

	// Add the new SocMedPost to the slice
	posts = append(posts, newSMP)
	c.IndentedJSON(http.StatusCreated, newSMP)
}

func getPostByID(c *gin.Context) {
	//load ID
	id := c.Param("id")
	// Convert it so string so a.ID works later on
	idStr, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID value"})
	}
	// Loop over the list of posts
	// Is looking for a post with specific ID
	for _, a := range posts {
		if a.ID == idStr {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Post not found"})
}

// Main function
func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostByID)
	router.POST("/posts", postPosts)

	router.Run("localhost:8080")
}
