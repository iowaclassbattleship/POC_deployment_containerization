package router

import (
	"poc/post/controller"

	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.GET("/post", controller.GetPost)
	api.POST("/post", controller.CreatePost)
	api.PUT("/post", controller.UpdatePost)
	api.DELETE("/post", controller.DeletePost)

	r.NoRoute(controller.NotFound)

	r.Run(":8081")
}
