package controller

import (
	"poc/post/models"

	"github.com/gin-gonic/gin"
)

func writeMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"message":    message,
	})
}

func writeResult(c *gin.Context, statusCode int, result []models.Post) {
	c.JSON(statusCode, gin.H{
		"statusCode": statusCode,
		"result":     result,
	})
}
