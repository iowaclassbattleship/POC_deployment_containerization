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
			writeMessage(c, http.StatusInternalServerError, err.Error())
			return
		}

		writeResult(c, http.StatusOK, posts)
		return
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		writeMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := db.GetPostByID(objectID)
	if err != nil {
		writeMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := []models.Post{post}

	writeResult(c, http.StatusOK, result)
}

func CreatePost(c *gin.Context) {
	var requestBody models.PostRequestBody

	err := c.BindJSON(&requestBody)
	if err != nil {
		writeMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	objectID, err := db.CreatePost(requestBody)
	if err != nil {
		writeMessage(c, http.StatusInternalServerError, err.Error())
	}

	writeMessage(c, http.StatusOK, fmt.Sprintf("Post created successfully with id %s", objectID))
}

func UpdatePost(c *gin.Context) {
	var requestBody models.PostRequestBody

	err := c.BindJSON(&requestBody)
	if err != nil {
		writeMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	objectID, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		writeMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err = db.UpdatePost(objectID, requestBody)
	if err != nil {
		writeMessage(c, http.StatusInternalServerError, err.Error())
	}

	writeMessage(c, http.StatusOK, fmt.Sprintf("Post updated successfully with id %s", objectID))
}

func DeletePost(c *gin.Context) {
	objectID, err := primitive.ObjectIDFromHex(c.Query("id"))
	if err != nil {
		writeMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err = db.DeletePost(objectID)
	if err != nil {
		writeMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	writeMessage(c, http.StatusOK, fmt.Sprintf("Post with id %s deleted successfully", objectID))
}

func NotFound(c *gin.Context) {
	writeMessage(c, http.StatusNotFound, "Ressource could not be found")
}
