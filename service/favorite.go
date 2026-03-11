package service

import (
	"context"
	"mginmall/dao"
	"mginmall/model"
	"mginmall/pkg/e"
	"mginmall/serializer"
	"strconv"
)

type FavoriteService struct {
	ProductID  uint `json:"product_id" form:"product_id"`
	BossID     uint `json:"boss_id" form:"boss_id"`
	FavoriteID uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoriteService) List(c context.Context, uid uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(c)
	code := e.Success
	favorite, err := favoriteDao.ListFavorite(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(c, favorite), uint(len(favorite)))
}

func (service *FavoriteService) Create(c context.Context, uid uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(c)
	code := e.Success
	exist, _ := favoriteDao.IsFavoriteExists(service.ProductID, uid)
	if exist {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(c)
	user, err := userDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	bossDao := dao.NewUserDao(c)
	boss, err := bossDao.GetUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(c)
	product, err := productDao.GetProductByID(service.ProductID)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	favorite := &model.Favorite{
		User:      *user,
		UserID:    uid,
		Product:   *product,
		ProductID: service.ProductID,
		Boss:      *boss,
		BossID:    service.BossID,
	}

	err = favoriteDao.CreateFavorite(favorite)
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
	}
}

func (service *FavoriteService) Delete(c context.Context, uid uint, pId string) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(c)
	code := e.Success
	fid, _ := strconv.Atoi(pId)
	err := favoriteDao.DeleteFavorite(uid, uint(fid))
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
	}
}
