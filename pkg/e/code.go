package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400
	Unauthorized  = 401

	//user模块错误
	ErrorExistUser             = 30001
	ErrorFailEncryption        = 30002
	ErrorNotExistUser          = 30003
	ErrorNotCompare            = 30004
	ErrorAuthToken             = 30005
	ErrorAuthCheckTokenTimeout = 30006
	ErrorUpLoadFail            = 30007
	ErrorAuthCheckTokenFail    = 30008

	//
	ErrorSendEmail = 40001

	//product模块错误
)
