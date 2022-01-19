package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID                    uint
	Name                  string
	Stock                 int
	Price                 int
	Category_id           uint
	Detail_transaction_ID []Detail_transaction
	Detail_cart_ID        []Detail_cart
}
