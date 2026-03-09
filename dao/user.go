package dao

import (
	"context"
	"mginmall/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Where("user_name = ?", userName).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) GetUserById(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uid uint, user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
}
