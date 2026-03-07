package service

import (
	"context"
	"mginmall/dao"
	"mginmall/model"
	"mginmall/pkg/e"
	"mginmall/pkg/utils"
	"mginmall/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //前端验证

}

func (service UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}

	//10 000 --->密文加密 对称加密
	utils.Encrypt.SetKey(service.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Email:    "",
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    utils.Encrypt.AesEncoding("10000"),
	}

	//规范，前端传密文，我们要解密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//创建用户
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service UserService) Login(ctx context.Context) serializer.Response {
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "未知错误",
		}
	}

	if !exist {
		code = e.ErrorNotExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册",
		}
	}

	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重新输入",
		}
	}
}
