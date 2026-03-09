package service

import (
	"context"
	"mginmall/conf"
	"mginmall/dao"
	"mginmall/model"
	"mginmall/pkg/e"
	"mginmall/pkg/utils"
	"mginmall/serializer"
	"mime/multipart"
	"strings"
	"time"

	"gopkg.in/mail.v2"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` //前端验证
}
type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	// 1->绑定邮箱,2->解绑邮箱,3->修改密码
}

type ValidEmailService struct {
	Token string `json:"token" form:"token"`
}

type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
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
		Avatar:   "avatar.jpg",
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

func (service *UserService) Login(ctx context.Context) serializer.Response {
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

	token, err := utils.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)

	if service.NickName != "" {
		user.NickName = service.NickName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}

}

func (service *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	code := e.Success
	var user *model.User
	var err error
	UserDao := dao.NewUserDao(ctx)
	user, err = UserDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	path, err := UploadAvatarToLocalStatic(file, uId, user.UserName)
	if err != nil {
		code = e.ErrorUpLoadFail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = UserDao.UpdateUserById(uId, user)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (service *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var address string
	var notice *model.Notice // 绑定邮箱、修改密码、解绑邮箱时使用通知模板发送邮件
	// 生成邮件校验 token，用于绑定/解绑邮箱或修改密码的后续验证
	token, err := utils.GenerateEmailToken(uId, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 拼接邮件验证链接地址
	address = conf.Config.Email.ValidEmail + token //发送方
	mailStr := notice.Text
	// 替换模板中的占位符，生成最终邮件内容
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()
	m.SetHeader("From", conf.Config.Email.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "Kirep")
	m.SetBody("text/html", mailText)
	d := mail.NewDialer(conf.Config.Email.SmtpHost, 465, conf.Config.Email.SmtpEmail, conf.Config.Email.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// 通过 SMTP 发送邮件
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *ValidEmailService) Vaild(ctx context.Context, token string) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.Success

	if token == "" {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "token 不能为空",
		}
	} else {
		claims, err := utils.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  "token 已过期",
			}
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if operationType == 1 {
		user.Email = email
	} else if operationType == 2 {
		user.Email = ""
	} else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (service *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, service.Key),
		Msg:    e.GetMsg(code),
	}
}
