package e

var MsgFlags = map[int]string{
	Success:                    "ok",
	Error:                      "fail",
	InvalidParams:              "参数错误",
	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorNotExistUser:          "用户不存在",
	ErrorNotCompare:            "密码错误",
	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeout: "token过期",
	Unauthorized:               "未认证",
	ErrorUpLoadFail:            "上传失败",
	ErrorAuthCheckTokenFail:    "toke解析失败",
	ErrorSendEmail:             "邮件发送失败",
}

// GetMsg
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
