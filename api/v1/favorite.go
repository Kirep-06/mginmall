package api

import (
	"mginmall/pkg/utils"
	"mginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFavorite(c *gin.Context) {
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	createdFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&createdFavoriteService); err == nil {
		res := createdFavoriteService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("CreatedFavorite bind error:%v", err)
	}
}

func ListFavorite(c *gin.Context) {
	listFavoriteService := service.FavoriteService{}

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listFavoriteService); err == nil {
		res := listFavoriteService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("List Favorite error: %v", err)
	}
}

func DeleteFavorite(c *gin.Context) {
	deleteFavoriteService := service.FavoriteService{}

	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		utils.LogrusObj.Errorf("Delete Favorite error: %v", err)
	}
}
