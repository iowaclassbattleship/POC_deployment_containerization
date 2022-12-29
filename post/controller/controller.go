package controller

import (
	"fmt"
	"net/http"
	"poc/post/db"
	"poc/post/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPost(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		posts, err := db.GetPosts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, messageJSON(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, resultJSON(http.StatusOK, posts))
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageJSON(http.StatusBadRequest, err.Error()))
		return
	}

	post, err := db.GetPostByID(objectID)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageJSON(http.StatusInternalServerError, err.Error()))
		return
	}

	result := []models.Post{post}

	c.JSON(http.StatusOK, resultJSON(http.StatusOK, result))
}

func CreatePost(c *gin.Context) {
	var requestBody models.PostRequestBody

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageJSON(http.StatusBadRequest, err.Error()))
		return
	}

	objectID, err := db.CreatePost(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, messageJSON(http.StatusInternalServerError, err.Error()))
	}

	c.JSON(http.StatusOK, messageJSON(http.StatusOK, fmt.Sprintf("Post created successfully with id %s", objectID)))
}

func UpdatePost(c *gin.Context) {
	var requestBody models.PostRequestBody

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageJSON(http.StatusBadRequest, err.Error()))
		return
	}

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, messageJSON(http.StatusOK, "No ID sent with"))
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageJSON(http.StatusBadRequest, err.Error()))
		return
	}

	err = db.UpdatePost(objectID, requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, messageJSON(http.StatusInternalServerError, err.Error()))
	}

	c.JSON(http.StatusOK, messageJSON(http.StatusOK, fmt.Sprintf("Post updated successfully with id %s", objectID)))
}

func DeletePost(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, messageJSON(http.StatusBadRequest, err.Error()))
		return
	}

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
