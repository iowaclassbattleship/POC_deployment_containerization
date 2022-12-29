package router

import (
	"poc/user/controller"

	"github.com/gin-gonic/gin"
)

func Serve() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.POST("/auth", controller.Login)

	r.Run(":8081")
}
