package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tmpbook/GC4K/controllers"
	"github.com/tmpbook/GC4K/middlewares"
)

// RegisterRouter 注册路由
func RegisterRouter(r *gin.Engine) {
	r.POST("/register", controllers.UserRegister)
	r.POST("/auth", controllers.UserAuth)
	apiAuthorized := r.Group("/api")
	// 单路由的 middleware
	apiAuthorized.Use(middlewares.APIAuth())
	{
		// nested group
		ci := apiAuthorized.Group("ci")
		cd := apiAuthorized.Group("cd")
		ci.GET("/create", controllers.CICreate)
		cd.GET("/create", controllers.CDCreate)
	}
}
