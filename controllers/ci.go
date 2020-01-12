package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmpbook/GC4K/services"
	"github.com/tmpbook/GC4K/types"
)

// CICreate ...
func CICreate(c *gin.Context) {
	services.DoCreateCI()
	c.JSON(201, types.ResponseBody{Code: 0, Msg: "ci created"})
}
