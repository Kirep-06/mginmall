package service

import (
	"context"
	"mginmall/dao"
	"mginmall/serializer"
	"strconv"
)

type ListProductImgService struct {
}

func (service *ListProductImgService) ListProductImg(c context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(c)
	id, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImg(uint(id))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))

}
