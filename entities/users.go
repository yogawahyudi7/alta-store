package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            uint
	Name          string
	Email         string `json:"email" gorm:"unique"`
	Password      string `json:"password"`
	Role          string `gorm:"default:member"`
	TransactionID []Transaction
	CartID        uint
}
