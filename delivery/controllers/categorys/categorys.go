package categorys

import (
	"net/http"
	"project-e-commerces/entities"
	"project-e-commerces/repository/categorys"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	Repo categorys.CategoryInterface
}

func NewCategoryControllers(cc categorys.CategoryInterface) *CategoryController {
	return &CategoryController{Repo: cc}
}

func (c *CategoryController) GetAllCategory(e echo.Context) error {
	res := GetCategoryResponseFormat{}

	categorys, err := c.Repo.GetAllCategory()

	if err != nil || len(categorys) == 0 {
		res.Message = "no categorys data found"
		res.Data = categorys

		return e.JSON(http.StatusNotFound, res)
	}

	res.Message = "success get all category"
	res.Data = append(res.Data, categorys...)

	return e.JSON(http.StatusOK, res)
}

func (c *CategoryController) GetCategoryByID(e echo.Context) error {
	res := GetCategoryResponseFormat{}

	category_id, _ := strconv.Atoi(e.Param("id"))

	// category_id, err := strconv.Atoi(e.Param("id"))

	// if err != nil {
	// 	res.Message = "can't get the id"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	category, err := c.Repo.GetCategoryByID(category_id)

	if err != nil || category.ID == 0 {
		res.Message = "category not found"
		res.Data = nil

		return e.JSON(http.StatusNotFound, res)
	}

	res.Message = "success get category"
	res.Data = append(res.Data, category)

	return e.JSON(http.StatusOK, res)
}

func (c *CategoryController) CreateCategory(e echo.Context) error {
	res := GetCategoryResponseFormat{}

	input := CreateCategoryRequest{}

	e.Bind(&input)

	// err := e.Bind(&input)

	// if err != nil {
	// 	res.Message = fmt.Sprint("can't get the input", err.Error())
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	// err = e.Validate(&input)

	// if err != nil {
	// 	// res.Message = "there is data that has not been filled in"
	// 	res.Message = fmt.Sprint("errornya disini ", err.Error())
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	category := entities.Category{}

	category.Name = input.Name

	newCategory, err := c.Repo.CreateCategory(category)

	if err != nil {
		res.Data = nil
		res.Message = "can't create category"

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success create category"
	res.Data = append(res.Data, newCategory)

	return e.JSON(http.StatusOK, res)
}

func (c *CategoryController) UpdateCategory(e echo.Context) error {
	res := GetCategoryResponseFormat{}

	input := CreateCategoryRequest{}

	e.Bind(&input)

	category := entities.Category{}

	category_id, _ := strconv.Atoi(e.Param("id"))

	// category_id, err := strconv.Atoi(e.Param("id"))

	// if err != nil {
	// 	res.Message = "can't get the id"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	// err = e.Validate(&input)

	// if err != nil {
	// 	res.Message = "there is data that has not been filled in"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	category.Name = input.Name

	updatedCategory, err := c.Repo.UpdateCategory(category_id, category)

	if err != nil {
		res.Message = "can't update category"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success update category"
	res.Data = append(res.Data, updatedCategory)

	return e.JSON(http.StatusOK, res)
}

func (c *CategoryController) DeleteCategory(e echo.Context) error {
	res := GetCategoryResponseFormat{}

	category_id, _ := strconv.Atoi(e.Param("id"))

	// category_id, err := strconv.Atoi(e.Param("id"))

	// if err != nil {
	// 	res.Message = "can't get the id"
	// 	res.Data = nil

	// 	return e.JSON(http.StatusBadRequest, res)
	// }

	_, err := c.Repo.DeleteCategory(category_id)

	if err != nil {
		res.Message = "can't delete category"
		res.Data = nil

		return e.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success delete category"
	res.Data = nil

	return e.JSON(http.StatusOK, res)
}
