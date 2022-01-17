package entities

type Transaction struct {
	ID uint
}

type Detail_transaction struct {
	ID         uint      `gorm:"foreignKey:ID"`
	Product_id []Product `gorm:"foreignKey:ID"`
}
