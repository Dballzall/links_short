package handler

import (
	"net/http"

	"github.com/Dballzall/links-short/shortener"
	"github.com/Dballzall/links-short/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

// CreateShortUrl handles the creation of a short URL
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Short URL created successfully",
		"short_url": shortUrl,
		"long_url":  creationRequest.LongUrl,
		"user_id":   creationRequest.UserId,
	})
}

// HandleShortUrlRedirect handles the redirection from the short URL to the original URL
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	// Skip favicon.ico requests
	if shortUrl == "favicon.ico" {
		c.Status(http.StatusNotFound)
		return
	}

	initialUrl, err := store.RetrieveInitialUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Short URL not found",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, initialUrl)
}

// GetRecentUrls handles the retrieval of recent URLs
func GetRecentUrls(c *gin.Context) {
	urls, err := store.GetRecentUrls(5) // Get last 5 URLs
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, urls)
}
