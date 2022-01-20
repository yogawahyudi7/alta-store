package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID        uint
	Name      string
	ProductID []Product
}
