package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	gorm.Model
	Username string  `json:"username"` // 列名为 `username`
	Password string  `json:"password"` // 列名为 `password`
	CnName   string  `json:"cnname"`
	IsActive bool    `json:"isactive"`
	Profile  Profile //与Profile为one2one关系
	Apps     []App   //与App为one2many关系
}

// Profile 描述信息模型
type Profile struct {
	gorm.Model
	UserID  uint   `gorm:"index"` // 外键 (属于), tag `index`是为该列创建索引
	Email   string `json:"email"`
	Address string `json:"address"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
