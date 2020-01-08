package main

import (
	"github.com/gin-gonic/gin"
	orm "github.com/tmpbook/GC4K/models"
	"github.com/tmpbook/GC4K/router"
	"github.com/tmpbook/GC4K/types"
)

func main() {
	defer orm.DB.Close()
	// 不包含任何 middleware
	r := gin.New()

	// 全局的 middleware
	r.Use(gin.Logger())
	// 任何 panic 都会被捕捉并且返回 500
	r.Use(gin.Recovery())

	// 404 的时候返回自定义 json
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, types.ResponseBody{Code: 404, Message: "page not found"})
	})

	// 注册路由
	router.RegisterRouter(r)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8088")
}
