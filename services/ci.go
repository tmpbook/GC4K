package services

import "github.com/tmpbook/GC4K/models"

// DoCreateCI 真正创建 CI 的操作
func DoCreateCI() {
	models.CreateCI()
}
