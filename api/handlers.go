package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/url-shortener/internal/postgres"
	"github.com/keivanipchihagh/url-shortener/internal/shortener"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

// Creates a short URL, stores it in Redis and returns the result as JSON
func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	// Vadlidate the request body
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a short link and store it in Redis
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	postgres.StoreUrlMapping(shortUrl, creationRequest.LongUrl)

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"short_url": host + shortUrl,
	})
}

// Redirects the user to the initial URL
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := postgres.RetrieveOriginalUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
