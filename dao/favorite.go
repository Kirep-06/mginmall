package dao

import (
	"context"
	"mginmall/model"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao *FavoriteDao) ListFavorite(uId uint) (resp []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=?", uId).Find(&resp).Error
	return
}

func (dao *FavoriteDao) IsFavoriteExists(pid uint, uid uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id = ? AND user_id=?", pid, uid).Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (dao *FavoriteDao) CreateFavorite(in *model.Favorite) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Create(&in).Error
	return
}

func (dao *FavoriteDao) DeleteFavorite(uid uint, pid uint) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=? AND product_id=?", uid, pid).Delete(&model.Favorite{}).Error
	return
}
