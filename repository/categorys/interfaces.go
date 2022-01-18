package categorys

import (
	"project-e-commerces/entities"
)

type CategoryInterface interface {
	GetAllCategory() ([]entities.Category, error)
	GetCategoryByID(id int) (entities.Category, error)
	CreateCategory(category entities.Category) (entities.Category, error)
	UpdateCategory(category_id int, category entities.Category) (entities.Category, error)
	DeleteCategory(category_id int) (entities.Category, error)
}
