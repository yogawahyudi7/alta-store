package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID          uint
	Total       int
	Total_price int
	Total_qty   int
}

type Detail_transaction struct {
	gorm.Model
	Transaction_id []Transaction `gorm:"foreignKey:ID"`
	Product_id     []Product     `gorm:"foreignKey:ID"`
	Product_qty    int
	Price          int
}
