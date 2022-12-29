package controller

import (
	"log"
	"net/http"
	"poc/post/db"
	"poc/post/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPostByID(c *gin.Context) {
	_, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    err,
		})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "bing chilling",
	})
}

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "bong chilling",
	})
}

func CreatePost(c *gin.Context) {
	var requestBody models.RequestCreatePost

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"statusCode": http.StatusBadRequest,
			"message":    "corrupt JSON request body",
		})
		log.Fatal(err)
	}

	db.CreatePost(requestBody)

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Post created successfully",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"statusCode": http.StatusNotFound,
		"message":    "Page not found",
	})
}
