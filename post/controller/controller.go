package controller

import (
	"log"
	"net/http"
	"poc/post/db"
	"poc/post/models"

	"github.com/gin-gonic/gin"
)

func GetPostByID(c *gin.Context) {
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"result":     db.GetPostByID(id),
	})
}

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"result":     db.GetPosts(),
	})
}

func CreatePost(c *gin.Context) {
	var requestBody models.RequestCreatePost

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, envelope(http.StatusBadRequest, "corrupted JSON body"))
		log.Fatal(err)
	}

	db.CreatePost(requestBody)

	c.JSON(http.StatusOK, envelope(http.StatusOK, "Post created successfully"))
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, envelope(http.StatusNotFound, "Ressource could not be found"))
}
