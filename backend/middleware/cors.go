package middleware

import "github.com/gin-gonic/gin"

// Cors 中间件用于处理跨域请求
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许跨域访问的域名，可以根据需求进行修改
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// 检查请求方法是否为预检请求（OPTIONS），如果是则返回成功，不继续处理
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		// 继续处理请求
		c.Next()
	}
}
