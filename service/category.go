package service

import (
	"context"
	"mginmall/dao"
	"mginmall/pkg/e"
	"mginmall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(c context.Context) serializer.Response {
	categoryDao := dao.NewCategoryDao(c)
	code := e.Success
	categorys, err := categoryDao.ListCategory()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategorys(categorys), uint(len(categorys)))
}
