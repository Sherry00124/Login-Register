package main

import (
	"Register-Login-Project/controller"
	"Register-Login-Project/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/Register", controller.Register)
	r.POST("/Login", controller.Login)
	r.GET("/info", middleware.JWTAuthMiddleware(), controller.Info) //用户信息
	r.POST("/upload", controller.FileReceive)                       //文件上传

	categoryRoutes := r.Group("/categories")
	categoryRoutes.POST("")
	return r
}
