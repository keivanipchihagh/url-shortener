package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keivanipchihagh/url-shortener/internal/postgres"
	"github.com/keivanipchihagh/url-shortener/internal/shortener"
)

// Request model definition
type UrlCreationRequest struct {
	Url        string `json:"url" binding:"required"`
	HashLength int    `json:"hash_length" binding:"required"`
}

// Creates a hash, stores in database and returns the response as JSON (keys: url, hash)
func CreateHash(c *gin.Context) {
	var creationRequest UrlCreationRequest

	// Parse and validate the request body
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a short link and store it in the database
	hash := shortener.GenerateHash(creationRequest.Url, creationRequest.HashLength)
	postgres.StoreUrlMapping(hash, creationRequest.Url)

	c.JSON(http.StatusOK, gin.H{
		"url":  creationRequest.Url,
		"hash": hash,
	})
}

// Retrieves the original URL from database and returns the response as JSON (keus: url, hash)
func GetUrl(c *gin.Context) {
	// hash, err := c.Params.Get("hash")
	hash, err := c.GetQuery("hash")
	if !err {
		c.JSON(http.StatusNotFound, http.NotFoundHandler())
		return
	}

	url := postgres.RetrieveOriginalUrl(hash)

	c.JSON(http.StatusOK, gin.H{
		"url":  url,
		"hash": hash,
	})
}
