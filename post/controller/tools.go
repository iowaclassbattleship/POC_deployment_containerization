package controller

import "github.com/gin-gonic/gin"

func envelope(statusCode uint, message string) gin.H {
	return gin.H{
		"statusCode": statusCode,
		"message":    message,
	}
}
