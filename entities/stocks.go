package entities

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	Product_id int
	Qty        int
}
