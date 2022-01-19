package carts

import "project-e-commerces/entities"

type AddItemIntoDetail_CartResponsesFormat struct {
	Message string                 `json:"message"`
	Data    []entities.Detail_cart `json:"data"`
}

type DelItemIntoDetail_CartResponsesFormat struct {
	Message string                 `json:"message"`
	Data    []entities.Detail_cart `json:"data"`
}
