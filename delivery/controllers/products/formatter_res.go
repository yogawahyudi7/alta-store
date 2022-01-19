package products

import "project-e-commerces/entities"

type GetProductResponseFormat struct {
	Message string             `json:"message"`
	Data    []entities.Product `json:"data"`
}

type GetStockProductResponseFormat struct {
	Message string           `json:"message"`
	Data    []entities.Stock `json:"data"`
}
