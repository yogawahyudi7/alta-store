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

func (cr *CartsRepository) Insert(newCart entities.Cart) (entities.Cart, error) {
	cr.db.Save(&newCart)
	return newCart, nil
}

func (cr *CartsRepository) InsertProduct(cartID uint, newProduct entities.Detail_cart) (entities.Detail_cart, error) {
	cr.db.Save(&newProduct)
	return newProduct, nil
}

func (cr *CartsRepository) UpdateProduct(cartID, productID uint, qty int) (entities.Detail_cart, error) {
	detail_cart := entities.Detail_cart{}
	cr.db.Where("cart_id=? AND product_id=?", cartID, productID).Find(&detail_cart)

	detail_cart.Qty += qty
	cr.db.Save(&detail_cart)
	return detail_cart, nil
}

func (cr *CartsRepository) DeleteProduct(cartID, productID uint) (entities.Detail_cart, error) {
	carts := entities.Detail_cart{}
	cr.db.Find(&carts, "cart_id=? AND product_id=?", cartID, productID).Delete(&carts)
	return carts, nil
}
