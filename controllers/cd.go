package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tmpbook/GC4K/services"
	"github.com/tmpbook/GC4K/types"
)

// CDCreate ...
func CDCreate(c *gin.Context) {
	services.DoCreateCD()
	c.JSON(201, types.ResponseBody{Code: 0, Msg: "cd created"})
}
