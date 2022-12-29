package router

import (
	"poc/post/controller"

	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()

	r.GET("/api/post/:id", controller.GetPostByID)
	r.DELETE("/api/post/:id", controller.DeletePost)

	r.GET("/api/post", controller.GetPosts)
	r.POST("/api/post", controller.CreatePost)

	r.NoRoute(controller.NotFound)

	r.Run(":8081")
}
