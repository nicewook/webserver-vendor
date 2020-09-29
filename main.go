package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicewook/webserver-vendor/models"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.Run()
}
