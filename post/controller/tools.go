package controller

import (
	"poc/post/models"

	"github.com/gin-gonic/gin"
)

func messageJSON(statusCode uint, message string) gin.H {
	return gin.H{
		"statusCode": statusCode,
		"message":    message,
	}
}

func resultJSON(statusCode uint, result models.Post) gin.H {
	return gin.H{
		"statusCode": statusCode,
		"result":     result,
	}
}
