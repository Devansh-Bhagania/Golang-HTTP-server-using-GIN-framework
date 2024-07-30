// main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML files
	r.LoadHTMLGlob("templates/*")

	// Define the routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})

	// Start the server
	r.Run(":8080") // Default listens and serves on 0.0.0.0:8080
}
