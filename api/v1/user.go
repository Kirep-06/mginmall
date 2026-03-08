package api

import (
	"mginmall/pkg/e"
	"mginmall/pkg/utils"
	"mginmall/serializer"
	"mginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService

	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserLogin(c *gin.Context) {
	var UserLogin service.UserService

	if err := c.ShouldBind(&UserLogin); err == nil {
		res := UserLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserUpdate(c *gin.Context) {
	var Userupdate service.UserService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&Userupdate); err == nil {
		res := Userupdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// func UploadAvatar(c *gin.Context) {
// 	file, fileHeader, err := c.Request.FormFile("file")
// 	if err != nil || fileHeader == nil {
// 		c.JSON(http.StatusBadRequest, serializer.Response{
// 			Status: e.InvalidParams,
// 			Msg:    "file 不能为空",
// 			Error:  "missing file",
// 		})
// 		return
// 	}
// 	defer file.Close()
// 	fileSize := fileHeader.Size

// 	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
// 	uploadAvatar := service.UserService{}
// 	res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileHeader.Filename, fileSize)
// 	c.JSON(http.StatusOK, res)
// }

func UploadAvatar(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil || fileHeader == nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Status: e.InvalidParams,
			Msg:    "file 不能为空",
			Error:  "missing file",
		})
		return
	}
	defer file.Close()
	fileSize := fileHeader.Size

	uploadAvatar := service.UserService{}
	claims, err := utils.ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, serializer.Response{
			Status: e.ErrorAuthCheckTokenFail,
			Msg:    "token 解析失败",
			Error:  err.Error(),
		})
		return
	}

	res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
	c.JSON(http.StatusOK, res)

}
