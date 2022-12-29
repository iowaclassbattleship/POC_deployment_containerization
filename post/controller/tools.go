package controller

import (
	"poc/post/models"

	"github.com/gin-gonic/gin"
)

func messageJSON(statusCode int, message string) gin.H {
	return gin.H{
		"statusCode": statusCode,
		"message":    message,
	}
}

func resultJSON(statusCode int, result []models.Post) gin.H {
	return gin.H{
		"statusCode": statusCode,
		"result":     result,
	}
}
