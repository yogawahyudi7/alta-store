package products

import (
	"math"
	"project-e-commerces/delivery/pagination"
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
	// productData := entities.Product{}

	// err := pr.db.Where("id = ?", product_id).Find(&productData).Error

	// if err != nil || productData.ID == 0 {
	// 	return productData, err
	// }

	// productData.Name = product.Name
	// productData.Price = product.Price
	// productData.Category_id = product.Category_id

	// err = pr.db.Save(&productData).Error

	// if err != nil || productData.ID == 0 {
	// 	return productData, err
	// }

	// return productData, nil

	productData := entities.Product{}

	err := pr.db.Model(&productData).Where("id = ?", product_id).Updates(product).Error

	if err != nil {
		return productData, err
	}

	return productData, nil
}

func (pr *productRepository) UpdateStockProduct(product_id, qty int) (entities.Product, error) {
	productData := entities.Product{}
	stockData := entities.Stock{}

	pr.db.Where("id = ?", product_id).Find(&productData)

	err := pr.db.Model(&productData).Where("id = ?", product_id).Update("stock", (productData.Stock + qty)).Error

	if err != nil {
		return productData, err
	}

	stockData.Product_id = product_id
	stockData.Qty = qty

	pr.db.Create(&stockData)

	return productData, err
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

func (pr *productRepository) ProductPagination(Pagination pagination.ProductPagination) (interface{}, int, error) {
	products := []entities.Product{}

	var totalRows, totalPages, fromRow, toRow int = 0, 0, 0, 0

	offset := Pagination.Page * Pagination.Limit

	err := pr.db.Limit(Pagination.Limit).Offset(offset).Find(&products).Error

	if err != nil {
		return Pagination, totalPages, err
	}

	Pagination.Rows = products

	totalRows64 := int64(totalRows)

	err = pr.db.Model(&entities.Product{}).Count(&totalRows64).Error

	if err != nil || totalRows64 == 0 {
		return Pagination, (totalPages), err
	}

	Pagination.TotalRows = int(totalRows64)

	totalPages = int(math.Ceil(float64(totalRows64)/float64(Pagination.Limit))) - 1

	if Pagination.Page == 0 {
		fromRow = 1
		toRow = Pagination.Limit
	} else {
		if Pagination.Page <= totalPages {
			fromRow = Pagination.Page*Pagination.Limit + 1
			toRow = (Pagination.Page + 1) * Pagination.Limit
		}
	}

	if toRow > int(totalRows64) {
		toRow = int(totalRows64)
	}

	Pagination.FromRow = fromRow
	Pagination.ToRow = toRow

	return Pagination, totalPages, err
}
