package categorys

import (
	"net/http"
	"project-e-commerces/delivery/common"
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
	res, err := c.Repo.GetAllCategory()

	if err != nil || len(res) == 0 {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "successful operation",
		"data":    res,
	})
}

func (c *CategoryController) GetCategoryByID(e echo.Context) error {
	category_id, _ := strconv.Atoi(e.Param("id"))

	res, err := c.Repo.GetCategoryByID(category_id)

	if err != nil || res.ID == 0 {
		return e.JSON(http.StatusNotFound, common.NewNotFoundResponse())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "successful operation",
		"data":    res,
	})
}

func (c *CategoryController) CreateCategory(e echo.Context) error {
	input := CreateCategoryRequest{}

	e.Bind(&input)

	category := entities.Category{}

	category.Name = input.Name

	res, err := c.Repo.CreateCategory(category)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Successful Operation",
		"data":    res,
	})
}

func (c *CategoryController) UpdateCategory(e echo.Context) error {
	input := CreateCategoryRequest{}

	e.Bind(&input)

	category := entities.Category{}

	category_id, _ := strconv.Atoi(e.Param("id"))

	category.Name = input.Name

	_, err := c.Repo.UpdateCategory(category_id, category)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}

func (c *CategoryController) DeleteCategory(e echo.Context) error {
	category_id, _ := strconv.Atoi(e.Param("id"))

	_, err := c.Repo.DeleteCategory(category_id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
	}

	return e.JSON(http.StatusOK, common.NewSuccessOperationResponse())
}
