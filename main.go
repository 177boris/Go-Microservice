//main.go

package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Create a router - use the default option provided by Gin
	router := gin.Default()

	// Load the templates
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

	// Start serving the application - port 8080 by default
	router.Run()

}
