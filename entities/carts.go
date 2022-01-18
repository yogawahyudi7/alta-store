package entities

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	ID             uint
	Total_Product  int
	Total_price    int
	DateCheckout   time.Time
	User           User
	Detail_cart_ID []Detail_cart
}

type Detail_cart struct {
	gorm.Model
	ID         uint
	CartID     uint
	ProductID  int
	Qty        int
	Price      int
	TotalPrice int
}
