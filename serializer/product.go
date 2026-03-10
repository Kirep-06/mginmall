package serializer

import (
	"mginmall/conf"
	"mginmall/model"
)

type Product struct {
	Id            uint
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	View          uint
	CreatedAt     int64
	Num           uint
	OnSale        bool
	BossId        uint
	BossName      string
	BossAvatar    string
}

func BuildProduct(item *model.Product) Product {
	return Product{
		Id:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryId,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.Config.Path.Host + conf.Config.Service.HttpPort + conf.Config.Path.ProductPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          uint(item.View()),
		CreatedAt:     item.CreatedAt.Unix(),
		Num:           item.Num,
		OnSale:        item.OnSale,
		BossId:        item.BossId,
		BossName:      item.BossName,
		BossAvatar:    item.BossAvatar,
	}
}
