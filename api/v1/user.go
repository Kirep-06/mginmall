package api

import (
	"mginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userRegister service.UserService

	if err := ctx.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(ctx.Request.Context())
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, err) 
	}

}
