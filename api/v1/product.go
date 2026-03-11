package api

import (
	"errors"
	"mginmall/pkg/utils"
	"mginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("CreatedProduct form error:%v", err)
		return
	}
	files := form.File["file"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse(errors.New("file 不能为空")))
		return
	}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	createdProductService := service.ProductService{}
	if err := c.ShouldBind(&createdProductService); err == nil {
		res := createdProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("CreatedProduct bind error:%v", err)
	}
}

func ListProduct(c *gin.Context) {
	listProductService := service.ProductService{}

	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("List Product error: %v", err)
	}
}

func SearchProduct(c *gin.Context) {
	searchProductService := service.ProductService{}

	if err := c.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("Search Product error: %v", err)
	}
}

func ShowProduct(c *gin.Context) {
	showProductService := service.ProductService{}

	if err := c.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("Search Product error: %v", err)
	}
}
