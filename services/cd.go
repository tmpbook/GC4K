package services

import "github.com/tmpbook/GC4K/models"

// DoCreateCD 真正创建 CD 的操作
func DoCreateCD() {
	models.CreateCD()
}
