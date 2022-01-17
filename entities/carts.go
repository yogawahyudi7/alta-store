package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID          uint
	Product_id  int
	Product_qty int
	Total_price int
}

type Detail_cart struct {
	Cart_id    []Cart    `gorm:"foreignKey:ID"`
	Product_id []Product `gorm:"foreignKey:ID"`
	Qty        int
}
