package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tmpbook/GC4K/utils"
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
	//TO DO
	//第一次访问数据库，之后要缓存起来
	return func(c *gin.Context) {
		var code int

		code = 0
		token := c.Query("token")
		if token == "" {
			code = 1
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 2
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 3
			}
		}

		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"Message": "token error",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
