package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            uint
	Name          string
	Email         string `gorm:"unique"`
	Password      string
	Role          string `gorm:"default:member"`
	TransactionID []Transaction
	CartID        uint
}
