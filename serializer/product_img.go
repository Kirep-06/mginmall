package serializer

import (
	"mginmall/conf"
	"mginmall/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductID: item.ProductID,
		ImgPath:   conf.Config.Path.Host + conf.Config.Service.HttpPort + conf.Config.Path.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImgs = append(productImgs, product)
	}
	return
}
