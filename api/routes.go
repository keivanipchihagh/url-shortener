package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// Registers the routes
func RegisterRoutes() {
	router.POST("/set", func(c *gin.Context) { CreateHash(c) })
	router.GET("/get", func(c *gin.Context) { GetUrl(c) })
}

// Starts the web server
func Run(host string, port int) {
	err := router.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
