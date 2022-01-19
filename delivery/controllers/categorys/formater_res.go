package categorys

import "project-e-commerces/entities"

type GetCategoryResponseFormat struct {
	Message string              `json:"message"`
	Data    []entities.Category `json:"data"`
}
