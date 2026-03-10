package api

import (
	"mginmall/pkg/utils"
	"mginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatedProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	createdProductSevice := service.ProductService{}
	if err := c.ShouldBind(&createdProductSevice); err == nil {
		res := createdProductSevice.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("CreatedProduct bind error:%v", err)
	}
}
