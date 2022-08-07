package middleware

import (
	"Register-Login-Project/common"
	"Register-Login-Project/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		tokenString := c.Request.Header.Get("Authorization")
		//log.Println("token", tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(tokenString, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		//tokenString = tokenString[7:1]
		//log.Println("token", parts[1])
		//log.Println("header", parts[0])
		token, claims, err := common.ParseToken(parts[1])
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		//验证通过后获取claim中的userid
		userID := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)
		//用户不存在
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "用户不存在",
			})
			c.Abort()
			return
		}
		//用户存在 将USER信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}
