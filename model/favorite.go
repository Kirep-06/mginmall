package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeginKey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeginKey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeginKey:BossId"`
	BossID    uint    `gorm:"not null"`
}
