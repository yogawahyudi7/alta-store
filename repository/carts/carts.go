package carts

import (
	"project-e-commerces/entities"

	"gorm.io/gorm"
)

type CartsRepository struct {
	db *gorm.DB
}

func NewCartsRepo(db *gorm.DB) *CartsRepository {
	return &CartsRepository{db: db}
}

func (cr *CartsRepository) Gets() ([]entities.Cart, error) {
	carts := []entities.Cart{}
	cr.db.Find(&carts)
	return carts, nil
}
