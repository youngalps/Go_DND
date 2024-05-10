package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Define your routes

	// Route for serving the HTML file
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Route for serving JSON response
	router.GET("/api/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to your D&D API!",
		})
	})

	// Start the server
	router.Run(":8080")
}
