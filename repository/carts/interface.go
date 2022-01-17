package carts

import "project-e-commerces/entities"

type CartInterface interface {
	Gets() ([]entities.Cart, error)
}
