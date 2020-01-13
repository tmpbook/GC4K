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

type Profile struct {
	gorm.Model
	UserID  uint   `gorm:"index"` // 外键 (属于), tag `index`是为该列创建索引
	Email   string `json:"email"`
	Address string `json:"address"`
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {

	var dbuser User
	DB.Model(User{}).Where("username = ?", user.Username).First(&dbuser)
	db_password := dbuser.Password

	err := bcrypt.CompareHashAndPassword([]byte(db_password), []byte(password))
	return err == nil
}

// CheckExist 校验已有用户
func (user *User) CheckExist() int {
	count := 0
	DB.Model(User{}).Where("username = ?", user.Username).Count(&count)

	return count
}

// 新增用户记录
func (user *User) Create() error {
	// 创建用户
	return DB.Create(&user).Error
}
