package categorys

import (
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategory(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	repo := NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category Alpha"}
		_, _ = repo.CreateCategory(mockCategory)
		categoryData, err := repo.GetAllCategory()

		assert.Equal(t, nil, err)
		assert.Equal(t, mockCategory.Name, categoryData[0].Name)
		assert.Equal(t, 1, int(categoryData[0].ID))
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Category{})
		categoryData, err := repo.GetAllCategory()

		assert.Equal(t, err, err)
		assert.Equal(t, []entities.Category{}, categoryData)
	})
}

func TestGetCategoryByID(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	repo := NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {

		mockCategory := entities.Category{Name: "Category Alpha"}
		res, _ := repo.CreateCategory(mockCategory)
		categoryData, err := repo.GetCategoryByID(int(res.ID))

		assert.Equal(t, mockCategory.Name, categoryData.Name)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, nil, err)
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Category{})
		categoryData, err := repo.GetCategoryByID(1)

		assert.Equal(t, err, err)
		assert.Equal(t, "", categoryData.Name)
		assert.Equal(t, 0, int(categoryData.ID))
	})
}

func TestCreateCategory(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	repo := NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCategory := entities.Category{Name: "Category Alpha"}
		createCategoryData, err := repo.CreateCategory(mockCategory)

		assert.Equal(t, nil, err)
		assert.Equal(t, 1, int(createCategoryData.ID))
		assert.Equal(t, mockCategory.Name, createCategoryData.Name)
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Category{})
		mockCategory := entities.Category{Name: "Category Alpha"}
		res, err := repo.CreateCategory(mockCategory)

		assert.Equal(t, "", "")
		assert.Equal(t, err, err)
		assert.Equal(t, 0, int(res.ID))
	})
}

func TestUpdateCategory(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	repo := NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCreateCategory := entities.Category{Name: "Category Alpha"}
		mockUpdateCategory := entities.Category{Name: "Category Alpha new"}

		createCategoryData, _ := repo.CreateCategory(mockCreateCategory)
		updateCategoryData, err := repo.UpdateCategory(int(createCategoryData.ID), mockUpdateCategory)

		assert.Equal(t, nil, err)
		assert.Equal(t, mockUpdateCategory.Name, updateCategoryData.Name)
	})

	t.Run("success-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Category{})
		mockUpdateCategory := entities.Category{Name: "Category Alpha new"}

		_, err := repo.UpdateCategory(1000, mockUpdateCategory)

		assert.Equal(t, err, err)
	})
}

func TestDeleteCategory(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Category{})
	db.AutoMigrate(&entities.Category{})

	repo := NewCategoryRepo(db)

	t.Run("success-case", func(t *testing.T) {
		mockCreateCategory := entities.Category{Name: "Category Alpha"}

		createCategoryData, _ := repo.CreateCategory(mockCreateCategory)
		categoryData, err := repo.DeleteCategory(int(createCategoryData.ID))

		assert.Equal(t, 0, int(categoryData.ID))
		assert.Equal(t, "", categoryData.Name)
		assert.Equal(t, nil, err)
	})

	t.Run("error-case", func(t *testing.T) {
		db.Migrator().DropTable(&entities.Category{})

		mockCreateCategory := entities.Category{Name: "Category Alpha"}

		createCategoryData, _ := repo.CreateCategory(mockCreateCategory)
		categoryData, err := repo.DeleteCategory(int(createCategoryData.ID))

		assert.Equal(t, "", categoryData.Name)
		assert.Equal(t, err, err)
		assert.Equal(t, int(createCategoryData.ID), int(categoryData.ID))
	})
}
