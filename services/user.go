package services

import (
	"github.com/tmpbook/GC4K/models"
	"github.com/tmpbook/GC4K/types"
	"github.com/tmpbook/GC4K/utils"
)

// UserRegisterService 管理用户注册服务
type UserRegister struct {
	Username        string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegister) valid(user models.User) *types.ResponseBody {
	if service.PasswordConfirm != service.Password {
		return &types.ResponseBody{
			Code: types.CodeParamErr,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := user.CheckExist()
	if count > 0 {
		return &types.ResponseBody{
			Code: types.CodeParamErr,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegister) Register() types.ResponseBody {
	user := models.User{
		Username: service.Username,
	}

	// 表单验证
	if err := service.valid(user); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return types.Err(
			types.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := user.Create(); err != nil {
		return types.ParamErr("注册失败", err)
	}

	return types.ResponseBody{
		Msg:  "注册成功",
		Data: user.Username,
	}
}

// UserLoginService 管理用户注册服务
type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *UserLogin) Login() types.ResponseBody {
	user := models.User{
		Username: service.Username,
		Password: service.Password,
	}

	count := user.CheckExist()
	if count > 0 {
		if user.CheckPassword(service.Password) == false {
			return types.ParamErr("密码错误", nil)
		} else {
			token, err := utils.GenerateToken(user.Username, user.Password)
			if err != nil {
				return types.ResponseBody{
					Code: types.CodeParamErr,
					Msg:  "token生成错误",
				}
			} else {
				return types.ResponseBody{
					Data: token,
					Msg:  "登陆成功",
				}
			}

		}
	} else {
		return types.ResponseBody{
			Code: types.CodeParamErr,
			Msg:  "账号错误",
		}
	}

}
