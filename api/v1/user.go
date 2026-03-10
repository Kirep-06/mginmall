package api

import (
	"mginmall/pkg/utils"
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
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("UserRegister bind error: %v", err)
	}
}

func UserLogin(c *gin.Context) {
	var UserLogin service.UserService

	if err := c.ShouldBind(&UserLogin); err == nil {
		res := UserLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("UserLogin bind error: %v", err)
	}
}

func UserUpdate(c *gin.Context) {
	var sendEmail service.UserService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("UserUpdate bind error: %v", err)
	}
}

func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size

	uploadAvatar := service.UserService{}
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("UploadAvatar bind error: %v", err)
	}

}

func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("SendEmail bind error: %v", err)
	}

}

func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService

	if err := c.ShouldBind(&validEmail); err == nil {
		token := validEmail.Token
		if token == "" {
			token = c.Query("token")
		}
		res := validEmail.Vaild(c.Request.Context(), token)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("ValidEmail bind error: %v", err)
	}
}

func ShowMoney(c *gin.Context) {
	var showMoney service.ShowMoneyService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoney); err == nil {
		res := showMoney.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("ShowMoney bind error: %v", err)
	}
}
