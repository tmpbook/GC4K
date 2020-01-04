package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// APIAuth 对 API 鉴权
func APIAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		// 对 API 调用进行鉴权
		fmt.Println("api auth here")
		c.Next()
	}
}

// UserAuth 对用户鉴权
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		// 认证用户，第一次可以访问数据库，后面记得使用 redis 做缓存
		fmt.Println("user auth here")
		c.Next()
	}
}
