package products

import (
	"project-e-commerces/entities"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (cr *productRepository) GetAllProduct() ([]entities.Product, error) {
	products := []entities.Product{}

	err := cr.db.Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (cr *productRepository) GetProductByID(product_id int) (entities.Product, error) {
	product := entities.Product{}

	err := cr.db.Where("ID = ?", product_id).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (cr *productRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	err := cr.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (cr *productRepository) UpdateProduct(product_id int, product entities.Product) (entities.Product, error) {
	productData := entities.Product{}

	err := cr.db.Where("id = ?", product_id).Find(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	productData.Name = product.Name
	productData.Price = product.Price
	productData.Category_id = product.Category_id

	err = cr.db.Save(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	return productData, nil
}

func (cr *productRepository) UpdateStockProduct(product_id, qty int) (entities.Product, error) {
	productData := entities.Product{}
	stockData := entities.Stock{}

	err := cr.db.Where("id = ?", product_id).Find(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	productData.Stock = productData.Stock + qty

	err = cr.db.Save(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	stockData.Product_id = product_id
	stockData.Qty = qty

	err = cr.db.Create(&stockData).Error

	if err != nil {
		return productData, err
	}

	return productData, nil
}

func (cr *productRepository) DeleteProduct(product_id int) (entities.Product, error) {
	product := entities.Product{}

	err := cr.db.Where("id = ?", product_id).Delete(&product).Error

	if err != nil || product.ID == 0 {
		return product, err
	}

	return product, nil
}
