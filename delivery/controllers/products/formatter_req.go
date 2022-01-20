package products

type CreateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
	Category_id int    `json:"category_id" form:"category_id"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Category_id int    `json:"category_id" form:"category_id"`
}

type UpdateStockProductRequest struct {
	Qty int `json:"qty" form:"qty"`
}
