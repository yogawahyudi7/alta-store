package products

import "project-e-commerces/entities"

type ProductInterface interface {
	GetAllProduct() ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	UpdateProduct(product_id int, product entities.Product) (entities.Product, error)
	UpdateStockProduct(product_id, qty int) (entities.Product, error)
	DeleteProduct(product_id int) (entities.Product, error)
	GetHistoryStockProduct(product_id int) ([]entities.Stock, error)
}
