package router

import (
	"poc/post/controller"

	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	r.GET("/api/post/:id", controller.GetPostByID)

	r.GET("/api/post", controller.GetPosts)
	r.POST("/api/post", controller.CreatePost)

	r.DELETE("/api/post", controller.DeletePost)

	r.NoRoute(controller.NotFound)

	r.Run(":8081")
}
