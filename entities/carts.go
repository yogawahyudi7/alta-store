package entities

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID           uint
	Product_id   int
	Product_qty  int
	Total_price  int
	DateCheckout time.Time
}

type Detail_cart struct {
	gorm.Model
	ID        uint
	ProductID int
	Qty       int
}
