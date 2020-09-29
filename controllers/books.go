package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicewook/webserver-vendor/models"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
