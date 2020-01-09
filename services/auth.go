package services

import (
	"github.com/tmpbook/GC4K/models"
)

func DoAuth(username, password string) bool {

	user := models.User{Username: username, Password: password}

	return user.CheckAuth()

}
