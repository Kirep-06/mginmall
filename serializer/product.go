package serializer

import (
	"mginmall/conf"
	"mginmall/model"
)

type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryID    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint   `json:"view"`
	CreatedAt     int64  `json:"created_at"`
	Num           uint   `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) Product {
	return Product{
		ID:         item.ID,
		Name:       item.Name,
		CategoryID: item.CategoryID,
		Title:      item.Title,
		Info:       item.Info,
		//ImgPath:    item.ImgPath,
		ImgPath:       conf.Config.Path.Host + conf.Config.Service.HttpPort + conf.Config.Path.ProductPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          uint(item.View()),
		CreatedAt:     item.CreatedAt.Unix(),
		Num:           item.Num,
		OnSale:        item.OnSale,
		BossID:        item.BossID,
		BossName:      item.BossName,
		BossAvatar:    conf.Config.Path.Host + conf.Config.Service.HttpPort + conf.Config.Path.AvataPath + item.BossAvatar,
	}
}

func BuildProducts(items []model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(&item)
		products = append(products, product)
	}
	return
}
