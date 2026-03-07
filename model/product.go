package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Category      string
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           uint
	BossID        uint
	BossName      string
	BossAvatar    string
}
