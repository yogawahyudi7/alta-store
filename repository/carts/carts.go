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

func (cr *CartsRepository) Get(cartID uint) ([]entities.Detail_cart, error) {
	dCarts := []entities.Detail_cart{}
	cr.db.Where("cart_id=?", cartID).Find(&dCarts)
	return dCarts, nil
}

func (cr *CartsRepository) Insert(newCart entities.Cart) (entities.Cart, error) {
	cr.db.Save(&newCart)
	return newCart, nil
}

func (cr *CartsRepository) InsertProduct(newProduct entities.Detail_cart) (entities.Detail_cart, error) {
	temp := entities.Detail_cart{}
	cr.db.Where("cart_id=? AND product_id=?", newProduct.CartID, newProduct.ProductID).Find(&temp)
	if temp.ID != 0 {
		temp.Qty += newProduct.Qty
		temp.TotalPrice += newProduct.TotalPrice
		cr.db.Save(&temp)
		return temp, nil
	}
	cr.db.Save(&newProduct)
	return newProduct, nil

}

func (cr *CartsRepository) DeleteProduct(cartID, productID uint) (entities.Detail_cart, error) {
	carts := entities.Detail_cart{}
	cr.db.Find(&carts, "cart_id=? AND product_id=?", cartID, productID).Delete(&carts)
	return carts, nil
}
