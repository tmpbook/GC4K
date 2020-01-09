package controllers

import (
	//"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tmpbook/GC4K/services"
	"github.com/tmpbook/GC4K/utils"
)

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	isExist := services.DoAuth(username, password)

	data := make(map[string]interface{})
	code := 1

	if isExist {
		token, err := utils.GenerateToken(username, password)
		if err != nil {
			code = 2
		} else {
			data["token"] = token

			code = 0
		}

	} else {
		code = 3
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}
