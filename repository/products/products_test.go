package products

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/repository/categorys"
	"project-e-commerces/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProduct(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Category{})

	repo := NewProductRepo(db)
	repoCategory := categorys.NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		_, _ = repo.CreateProduct(mockProduct)

		productData, _ := repo.GetAllProduct()

		assert.Equal(t, mockProduct.Name, productData[0].Name)
		assert.Equal(t, 1, int(productData[0].ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		productData, _ := repo.GetAllProduct()

		assert.Equal(t, []entities.Product{}, productData)
	})
}

func TestGetProductByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Category{})

	repo := NewProductRepo(db)
	repoCategory := categorys.NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		productData, _ := repo.GetProductByID(int(createProductData.ID))

		assert.Equal(t, mockProduct.Name, productData.Name)
		assert.Equal(t, 1, int(productData.ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.Migrator().DropTable(&entities.Category{})
		productData, _ := repo.GetProductByID(1)

		assert.Equal(t, "", productData.Name)
		assert.Equal(t, 0, int(productData.ID))
	})
}

func TestCreateProduct(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Category{})

	repo := NewProductRepo(db)
	repoCategory := categorys.NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		assert.Equal(t, 1, int(createProductData.ID))
		assert.Equal(t, mockProduct.Name, createProductData.Name)
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.Migrator().DropTable(&entities.Category{})
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		assert.Equal(t, int(0), int(createProductData.ID))
	})
}

func TestUpdateProduct(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Category{})

	repo := NewProductRepo(db)
	repoCategory := categorys.NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		mockUpdateProduct := entities.Product{Name: "Product Alpha new", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		updateProductData, _ := repo.UpdateProduct(int(createProductData.ID), mockUpdateProduct)

		assert.Equal(t, 1, int(updateProductData.ID))
		assert.Equal(t, mockUpdateProduct.Name, updateProductData.Name)
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.Migrator().DropTable(&entities.Category{})

		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		mockUpdateProduct := entities.Product{Name: "Product Alpha new", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		updateProductData, _ := repo.UpdateProduct(int(createProductData.ID), mockUpdateProduct)

		assert.Equal(t, "", updateProductData.Name)
		assert.Equal(t, int(createProductData.ID), int(updateProductData.ID))

	})
}

func TestUpdateStockProduct(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Stock{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Stock{})

	repo := NewProductRepo(db)
	repoCategory := categorys.NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		mockCreateStockProduct := entities.Stock{Product_id: 1, Qty: 1}

		fmt.Println(createProductData)

		createStockProductData, _ := repo.UpdateStockProduct(mockCreateStockProduct.Product_id, mockCreateStockProduct.Qty)

		assert.Equal(t, mockProduct.Stock+mockCreateStockProduct.Qty, createStockProductData.Stock)
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.Migrator().DropTable(&entities.Stock{})

		mockCreateStockProduct := entities.Stock{Product_id: 1, Qty: 1}
		createStockProductData, _ := repo.UpdateStockProduct(mockCreateStockProduct.Product_id, mockCreateStockProduct.Qty)

		assert.Equal(t, nil, createStockProductData)

	})
}

func TestDeleteProduct(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Category{})

	repo := NewProductRepo(db)
	repoCategory := categorys.NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		productData, _ := repo.DeleteProduct(int(createProductData.ID))

		assert.Equal(t, 0, int(productData.ID))
		assert.Equal(t, "", productData.Name)
	})

	t.Run("error-case", func(t *testing.T) {
		db.AutoMigrate(&entities.Product{})
		db.AutoMigrate(&entities.Category{})

		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		_, _ = repo.CreateProduct(mockProduct)

		productData, _ := repo.DeleteProduct(2)

		assert.Equal(t, "", productData.Name)
	})
}

func TestGetHistoryStockProduct(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Stock{})
	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.Stock{})

	repo := NewProductRepo(db)
	repoCategory := categorys.NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category 1"}
		createCategoryData, _ := repoCategory.CreateCategory(mockCategory)

		mockProduct := entities.Product{Name: "Product Alpha", Price: 10000, Stock: 1, Category_id: createCategoryData.ID}
		createProductData, _ := repo.CreateProduct(mockProduct)

		mockCreateStockProduct := entities.Stock{Product_id: 1, Qty: 1}

		_, _ = repo.UpdateStockProduct(mockCreateStockProduct.Product_id, mockCreateStockProduct.Qty)

		stockProductData, _ := repo.GetHistoryStockProduct(int(createProductData.ID))

		assert.Equal(t, mockCreateStockProduct.Qty, stockProductData[0].Qty)
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Product{})
		db.Migrator().DropTable(&entities.Stock{})

		mockCreateStockProduct := entities.Stock{Product_id: 1, Qty: 1}
		_, _ = repo.UpdateStockProduct(mockCreateStockProduct.Product_id, mockCreateStockProduct.Qty)

		stockProductData, _ := repo.GetHistoryStockProduct(int(1000))

		assert.Equal(t, 0, stockProductData[0].Qty)

	})
}
