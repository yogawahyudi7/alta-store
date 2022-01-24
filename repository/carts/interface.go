package carts

import "project-e-commerces/entities"

type CartInterface interface {
	Get(cartID uint) ([]entities.Detail_cart, error)
	Insert(newCart entities.Cart) (entities.Cart, error)
	InsertProduct(newItem entities.Detail_cart) (entities.Detail_cart, error)
	DeleteProduct(cartID, productID uint) (entities.Detail_cart, error)
}
