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

	// Serve static files from the "frontend" directory
	r.Static("/frontend", "./frontend")

	// Serve the index.html file at the root path
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	// Endpoint to create a short URL
	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// Endpoint to handle URL redirection
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Endpoint to get recent URLs
	r.GET("/recent", func(c *gin.Context) {
		handler.GetRecentUrls(c)
	})

	// Endpoint to get click statistics
	r.GET("/stats/:shortUrl", func(c *gin.Context) {
		handler.GetUrlClickStats(c)
	})

	// Endpoint to get URL information
	r.GET("/url-info/:shortUrl", func(c *gin.Context) {
		handler.GetUrlInfo(c)
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
