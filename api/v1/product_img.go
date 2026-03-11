package api

import (
	"mginmall/pkg/utils"
	"mginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductImg(c *gin.Context) {
	listProductImgService := service.ListProductImgService{}

	if err := c.ShouldBind(&listProductImgService); err == nil {
		res := listProductImgService.ListProductImg(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("Search Product error: %v", err)
	}
}
