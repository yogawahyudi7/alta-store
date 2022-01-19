package categorys

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCategory(t *testing.T) {
	t.Run("1-success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/categorys")

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.GetAllCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data

		name := data[0].Name

		assert.Equal(t, name, "Category Alpha")
		assert.Equal(t, "success get all category", response.Message)
	})

	t.Run("2-error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/categorys")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.GetAllCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Category(nil), response.Data)
		assert.Equal(t, "no categorys data found", response.Message)
	})
}

func TestGetCategoryByID(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/categorys/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.GetCategoryByID(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data[0]

		name := data.Name

		assert.Equal(t, name, "Category Alpha")
		assert.Equal(t, "success get category", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/?", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)

		context.SetPath("/categorys")
		context.SetParamNames("id")
		context.SetParamValues("2")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.GetCategoryByID(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Category([]entities.Category(nil)), response.Data)
		assert.Equal(t, "category not found", response.Message)
	})
}

func TestCreateCategory(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Category{
			Name: "Category Alpha",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/categorys")

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.CreateCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data

		fmt.Println(response)

		name := data[0].Name

		assert.Equal(t, "Category Alpha", name)
		assert.Equal(t, "success create category", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name": "Category Alpha",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/categorys")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.CreateCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response)

		assert.Equal(t, "can't create category", response.Message)
	})
}

func TestUpdateCategory(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Category Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/categorys")
		context.SetParamNames("id")
		context.SetParamValues("1")

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.UpdateCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		data := response.Data

		fmt.Println(response)

		name := data[0].Name

		assert.Equal(t, "Category Alpha new", name)
		assert.Equal(t, "success update category", response.Message)
	})
	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Category Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/categorys")
		context.SetParamNames("id")
		context.SetParamValues("1")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.UpdateCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Category([]entities.Category(nil)), response.Data)
		assert.Equal(t, "can't update category", response.Message)
	})
}

func TestDeleteCategory(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/categorys")
		context.SetParamNames("id")
		context.SetParamValues("1")

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.DeleteCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "success delete category", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/categorys")
		context.SetParamNames("id")
		context.SetParamValues("10")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.DeleteCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "can't delete category", response.Message)
	})
}

type mockCategoryRepository struct{}

func (m mockCategoryRepository) GetAllCategory() ([]entities.Category, error) {
	return []entities.Category{
		{ID: 1, Name: "Category Alpha"},
	}, nil
}

func (m mockCategoryRepository) GetCategoryByID(category_id int) (entities.Category, error) {
	return entities.Category{
		ID: 1, Name: "Category Alpha"}, nil
}

func (m mockCategoryRepository) CreateCategory(entities.Category) (entities.Category, error) {
	return entities.Category{
		ID: 1, Name: "Category Alpha"}, nil
}

func (m mockCategoryRepository) UpdateCategory(category_id int, category entities.Category) (entities.Category, error) {
	return entities.Category{
		ID: 1, Name: "Category Alpha new"}, nil
}

func (m mockCategoryRepository) DeleteCategory(category_id int) (entities.Category, error) {
	return entities.Category{
		ID: 1, Name: "Category Alpha"}, nil
}

type mockFalseCategoryRepository struct{}

func (m mockFalseCategoryRepository) GetAllCategory() ([]entities.Category, error) {
	return nil, errors.New("no data")
}

func (m mockFalseCategoryRepository) GetCategoryByID(id int) (entities.Category, error) {
	return entities.Category{
		ID: 0, Name: ""}, errors.New("can't get category")
}

func (m mockFalseCategoryRepository) CreateCategory(entities.Category) (entities.Category, error) {
	return entities.Category{
		ID: 0, Name: ""}, errors.New("error create category")
}

func (m mockFalseCategoryRepository) UpdateCategory(category_id int, category entities.Category) (entities.Category, error) {
	return entities.Category{
		ID: 0, Name: ""}, errors.New("error update category")
}

func (m mockFalseCategoryRepository) DeleteCategory(category_id int) (entities.Category, error) {
	return entities.Category{
		ID: 0, Name: ""}, errors.New("error delete category")
}
