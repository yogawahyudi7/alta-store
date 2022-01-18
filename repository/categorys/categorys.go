package categorys

import (
	"project-e-commerces/entities"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (cr *categoryRepository) GetAllCategory() ([]entities.Category, error) {
	categorys := []entities.Category{}

	err := cr.db.Find(&categorys).Error

	if err != nil {
		return categorys, err
	}

	return categorys, nil
}

func (cr *categoryRepository) GetCategoryByID(category_id int) (entities.Category, error) {
	category := entities.Category{}

	err := cr.db.Where("ID = ?", category_id).Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (cr *categoryRepository) CreateCategory(category entities.Category) (entities.Category, error) {
	err := cr.db.Save(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (cr *categoryRepository) UpdateCategory(category_id int, category entities.Category) (entities.Category, error) {
	categoryData := entities.Category{}

	err := cr.db.Where("id = ?", category_id).Find(&categoryData).Error

	if err != nil || categoryData.ID == 0 {
		return categoryData, err
	}

	categoryData.Name = category.Name

	err = cr.db.Save(&categoryData).Error

	if err != nil || categoryData.ID == 0 {
		return categoryData, err
	}

	return categoryData, nil
}

func (cr *categoryRepository) DeleteCategory(category_id int) (entities.Category, error) {
	category := entities.Category{}

	err := cr.db.Where("id = ?", category_id).Delete(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
