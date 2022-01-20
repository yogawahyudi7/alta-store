package entities

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID                    uint
	Total_price           int
	Total_qty             int
	User_id               uint
	PaymentID             uint
	Status                string
	Detail_transaction_ID []Detail_transaction
}

type Detail_transaction struct {
	gorm.Model
	ID             uint
	Transaction_id uint
	Product_id     uint
	Product_qty    int
	Price          int
}

type Payment struct {
	gorm.Model
	ID           uint
	Payment_type string
	Link         string
	Transaction  Transaction
}
