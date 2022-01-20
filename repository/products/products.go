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

func (pr *productRepository) GetAllProduct() ([]entities.Product, error) {
	products := []entities.Product{}

	err := pr.db.Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (pr *productRepository) GetProductByID(product_id int) (entities.Product, error) {
	product := entities.Product{}

	err := pr.db.Where("id = ?", product_id).Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	err := pr.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) UpdateProduct(product_id int, product entities.Product) (entities.Product, error) {
	productData := entities.Product{}

	err := pr.db.Where("id = ?", product_id).Find(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	productData.Name = product.Name
	productData.Price = product.Price
	productData.Category_id = product.Category_id

	err = pr.db.Save(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	return productData, nil
}

func (pr *productRepository) UpdateStockProduct(product_id, qty int) (entities.Product, error) {
	productData := entities.Product{}
	stockData := entities.Stock{}

	err := pr.db.Where("id = ?", product_id).Find(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	productData.Stock = productData.Stock + qty

	err = pr.db.Save(&productData).Error

	if err != nil || productData.ID == 0 {
		return productData, err
	}

	stockData.Product_id = product_id
	stockData.Qty = qty

	err = pr.db.Create(&stockData).Error

	if err != nil {
		return productData, err
	}

	return productData, nil
}

func (pr *productRepository) DeleteProduct(product_id int) (entities.Product, error) {
	product := entities.Product{}

	err := pr.db.Where("id = ?", product_id).Delete(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (pr *productRepository) GetHistoryStockProduct(product_id int) ([]entities.Stock, error) {
	stock := []entities.Stock{}

	err := pr.db.Where("product_id = ?", product_id).Find(&stock).Error

	if err != nil {
		return stock, err
	}

	return stock, nil
}
