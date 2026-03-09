package service

import (
	"context"
	"mginmall/dao"
	"mginmall/pkg/e"
	"mginmall/serializer"
)

type CarouselService struct {
}

func (service *CarouselService) List(c context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(c)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
