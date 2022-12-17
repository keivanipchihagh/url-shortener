package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func RegisterRoutes() {

	router.POST("/shorten", func(c *gin.Context) {
		CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		HandleShortUrlRedirect(c)
	})
}

// Starts the web server
func Run(host string, port int) {
	err := router.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
