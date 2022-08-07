package controller

import "github.com/gin-gonic/gin"

//定义增删改查
type ICategoryController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Show(c *gin.Context)
	Delete(c *gin.Context)
}

//为了让方法名能够复用
type CategoryController struct {
}
