package dao

import (
	"fmt"
	"mginmall/model"
)

func migration() {
	err := _db.Set("gorm:table_option", "charest=utf8mb4").
		AutoMigrate(
			&model.Address{},
			&model.Admin{},
			&model.Carousel{},
			&model.Cart{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.ProductImg{},
			&model.Product{},
			&model.User{},
		)
	if err != nil {
		fmt.Println("err", err)
	}
}
