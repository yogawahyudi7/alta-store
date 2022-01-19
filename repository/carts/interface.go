package carts

import "project-e-commerces/entities"

type CartInterface interface {
	Gets() ([]entities.Cart, error)
	Insert(newCart entities.Cart) (entities.Cart, error)
	InsertProduct(cartID uint, newItem entities.Detail_cart) (entities.Detail_cart, error)
	UpdateProduct(cartID, productID uint, qty int) (entities.Detail_cart, error)
	DeleteProduct(cartID, productID uint) (entities.Detail_cart, error)
}
