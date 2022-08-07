package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//跨域中间件

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")  //*代表所有域名都可以访问 http://localhost:8080
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")   //设置缓存时间
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*") //允许请求的方法GET POST...
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//判断是否为option请求，是则返回200
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next() //中间件继续向右判定
		}
	}
}
