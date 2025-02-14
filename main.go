package main

import (
	"fmt"
	"os"

	"github.com/Dballzall/links-short/handler"
	"github.com/Dballzall/links-short/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Home endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	// Endpoint to create a short URL
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Endpoint to handle URL redirection
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Initialize the URL store (e.g., connect to Redis)
	store.InitializeStore()

	// Use the PORT environment variable or default to 9808
	port := os.Getenv("PORT")
	if port == "" {
		port = "9808"
	}

	err := r.Run(":" + port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
