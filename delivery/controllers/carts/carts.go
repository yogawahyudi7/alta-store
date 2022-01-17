package carts

import "project-e-commerces/repository/carts"

type CartsController struct {
	Repo carts.CartInterface
}

func NewCartsControllers(tsrep carts.CartInterface) *CartsController {
	return &CartsController{Repo: tsrep}
}
