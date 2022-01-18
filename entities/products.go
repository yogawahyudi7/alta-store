package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint
	Name        string
	Stock       int
	Price       int
	Category_id []Category `gorm:"foreignKey:ID"`
}
