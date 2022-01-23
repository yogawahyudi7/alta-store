package categorys

import (
	"bytes"
	"encoding/json"
	"errors"
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
		assert.Equal(t, "successful operation", response.Message)
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
		assert.Equal(t, "Internal Server Error", response.Message)
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
		var response1 interface{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		json.Unmarshal([]byte(res.Body.Bytes()), &response1)

		// fmt.Println(response1)

		temp := response1.(map[string]interface{})

		data := temp["data"].(map[string]interface{})

		name := data["Name"]

		// categoryData := data[0].(map[string]interface{})

		// fmt.Println(categoryData["Name"])

		// data := response.Data[0]

		// name := data.Name

		assert.Equal(t, name, "Category Alpha")
		assert.Equal(t, "successful operation", response.Message)
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
		assert.Equal(t, "Not Found", response.Message)
	})
}

func TestCreateCategory(t *testing.T) {
	// e := echo.New()
	// jwtToken := ""
	// t.Run("login", func(t *testing.T) {
	// 	reqBody, _ := json.Marshal(map[string]string{
	// 		"email":    "mimin@mail.com",
	// 		"password": "mimin123",
	// 	})

	// 	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	// 	res := httptest.NewRecorder()

	// 	req.Header.Set("Content-Type", "application/json")
	// 	context := e.NewContext(req, res)
	// 	context.SetPath("/users/login")
	// 	context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

	// 	userCon := users.NewUserController(mockUserRepository{})
	// 	userCon.Login(context)

	// 	responses := users.LoginResponseFormat{}
	// 	json.Unmarshal([]byte(res.Body.Bytes()), &responses)

	// 	jwtToken = responses.Token
	// 	assert.Equal(t, "Successful Operation", responses.Message)
	// 	assert.Equal(t, 200, res.Code)

	// })

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
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.CreateCategory(context)

		response := GetCategoryResponseFormat{}

		var response1 interface{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		json.Unmarshal([]byte(res.Body.Bytes()), &response1)

		// fmt.Println(response1)

		temp := response1.(map[string]interface{})

		message := temp["message"].(string)

		// data := response.Data

		// name := data[0].Name

		// assert.Equal(t, "Category Alpha", name)
		assert.Equal(t, "Successful Operation", message)
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
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.CreateCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Internal Server Error", response.Message)
	})
}

func TestUpdateCategory(t *testing.T) {
	// e := echo.New()
	// jwtToken := ""
	// t.Run("login", func(t *testing.T) {
	// 	reqBody, _ := json.Marshal(map[string]string{
	// 		"email":    "mimin@mail.com",
	// 		"password": "mimin123",
	// 	})

	// 	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	// 	res := httptest.NewRecorder()

	// 	req.Header.Set("Content-Type", "application/json")
	// 	context := e.NewContext(req, res)
	// 	context.SetPath("/users/login")
	// 	// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

	// 	userCon := users.NewUserController(mockUserRepository{})
	// 	userCon.Login(context)

	// 	responses := users.LoginResponseFormat{}
	// 	json.Unmarshal([]byte(res.Body.Bytes()), &responses)

	// 	jwtToken = responses.Token
	// 	assert.Equal(t, "Successful Operation", responses.Message)
	// 	assert.Equal(t, 200, res.Code)

	// })

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
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.UpdateCategory(context)

		response := GetCategoryResponseFormat{}

		var response1 interface{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		json.Unmarshal([]byte(res.Body.Bytes()), &response1)

		// fmt.Println(response1)

		temp := response1.(map[string]interface{})

		message := temp["message"].(string)

		// assert.Equal(t, "Category Alpha new", name)
		assert.Equal(t, "Successful Operation", message)
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
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.UpdateCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Category([]entities.Category(nil)), response.Data)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
}

func TestDeleteCategory(t *testing.T) {
	// e := echo.New()
	// jwtToken := ""
	// t.Run("login", func(t *testing.T) {
	// 	reqBody, _ := json.Marshal(map[string]string{
	// 		"email":    "mimin@mail.com",
	// 		"password": "mimin123",
	// 	})

	// 	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
	// 	res := httptest.NewRecorder()

	// 	req.Header.Set("Content-Type", "application/json")
	// 	context := e.NewContext(req, res)
	// 	context.SetPath("/users/login")
	// 	// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

	// 	userCon := users.NewUserController(mockUserRepository{})
	// 	userCon.Login(context)

	// 	responses := users.LoginResponseFormat{}
	// 	json.Unmarshal([]byte(res.Body.Bytes()), &responses)

	// 	jwtToken = responses.Token
	// 	assert.Equal(t, "Successful Operation", responses.Message)
	// 	assert.Equal(t, 200, res.Code)

	// })

	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/categorys")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		categoryController := NewCategoryControllers(mockCategoryRepository{})
		categoryController.DeleteCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/categorys")
		context.SetParamNames("id")
		context.SetParamValues("10")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})
		categoryController.DeleteCategory(context)

		response := GetCategoryResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Internal Server Error", response.Message)
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

// type mockUserRepository struct{}

// func (mur mockUserRepository) Login(email string) (entities.User, error) {
// 	return entities.User{ID: 1, Email: "mimin@mail.com", Password: "mimin123", Role: "admin"}, nil
// }

// func (mur mockUserRepository) Register(newUser entities.User) (entities.User, error) {
// 	return entities.User{ID: 1, Email: "mimin@mail.com", Password: "mimin123", Role: "admin"}, nil
// }

// func (mur mockUserRepository) Get(user_id int) (entities.User, error) {
// 	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
// }
// func (mur mockUserRepository) Create(newUser entities.User) (entities.User, error) {
// 	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
// }
// func (mur mockUserRepository) Update(newUser entities.User) (entities.User, error) {
// 	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
// }
// func (mur mockUserRepository) Delete(user_id int) error {
// 	return nil
// }
