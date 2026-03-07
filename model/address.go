package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Name     string `gorm:"type:varchar(20) not null"`
	PhoneNum string `gorm:"type:varchar(11) not null"`
	Address  string `gorm:"type:varchar(50) not null"`
}
