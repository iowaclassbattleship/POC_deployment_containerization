package controller

import (
	"fmt"
	"log"
	"net/http"
	"poc/post/db"
	"poc/post/models"

	"github.com/gin-gonic/gin"
)

func GetPostByID(c *gin.Context) {
	id := c.Query("id")

	resultJSON(http.StatusOK, db.GetPostByID(id))
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
		c.JSON(http.StatusBadRequest, messageJSON(http.StatusBadRequest, "corrupted JSON body"))
		log.Fatal(err)
	}

	db.CreatePost(requestBody)

	c.JSON(http.StatusOK, messageJSON(http.StatusOK, "Post created successfully"))
}

func DeletePost(c *gin.Context) {
	id := c.Query("id")

	deletedId, err := db.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, messageJSON(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, messageJSON(http.StatusOK, fmt.Sprintf("Post %d deleted successfully", deletedId)))
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, messageJSON(http.StatusNotFound, "Ressource could not be found"))
}
