package controllers

import (
	//"log"

	"github.com/gin-gonic/gin"

	"github.com/tmpbook/GC4K/services"
	"github.com/tmpbook/GC4K/types"
)

// 用户注册
func UserRegister(c *gin.Context) {
	service := services.UserRegister{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, types.ErrorResponse(err))
	}

}

// 用户认证
func UserAuth(c *gin.Context) {
	service := services.UserLogin{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, types.ErrorResponse(err))
	}
}
