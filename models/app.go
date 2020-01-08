package models

import (
	"github.com/jinzhu/gorm"
)

// App组件模型
type App struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `gorm:"index"`
}
