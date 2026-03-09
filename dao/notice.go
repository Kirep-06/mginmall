package dao

import (
	"context"
	"mginmall/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (dao *NoticeDao) GetNoticeById(uid uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", uid).First(&notice).Error
	return
}

func (dao *NoticeDao) CreateNotice(notice *model.Notice) error {
	return dao.DB.Model(&model.Notice{}).Create(&notice).Error
}
