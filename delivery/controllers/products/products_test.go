package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/delivery/pagination"
	"project-e-commerces/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProduct(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockProductRepository{})
		productController.GetAllProduct(context)

		response := pagination.ProductPagination{}

		var response1 interface{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		json.Unmarshal([]byte(res.Body.Bytes()), &response1)

		temp := response1.(map[string]interface{})

		data := temp["data"].(map[string]interface{})

		productData := data["rows"].([]interface{})

		item := productData[0].(map[string]interface{})

		// fmt.Println(temp["message"])

		// fmt.Println(productData[0])

		// fmt.Println(item["Name"])

		assert.Equal(t, "Successful Operation", temp["message"])
		assert.Equal(t, "Product Alpha", item["Name"])
		assert.Equal(t, float64(200), temp["code"])

		// fmt.Println(res.Body)

		// product := response.Rows

		// fmt.Println(json.Unmarshal([]byte(res.Body.Bytes()), &response))

		// 	assert.Equal(t, "Successful Operation", response.Message)

		// 	assert.Equal(t, "Product Alpha", product[0].Name)
		// 	assert.Equal(t, "success get all product", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.GetAllProduct(context)

		response := pagination.ProductPagination{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		product := response.Rows

		assert.Equal(t, []entities.Product([]entities.Product(nil)), product)
	})
}

func TestGetProductByID(t *testing.T) {
	e := echo.New()
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

	// 	userCon := users.NewUserController(mockUserRepository{})
	// 	userCon.Login(context)

	// 	responses := users.LoginResponseFormat{}
	// 	json.Unmarshal([]byte(res.Body.Bytes()), &responses)

	// 	jwtToken = responses.Token
	// 	assert.Equal(t, "Successful Operation", responses.Message)
	// 	assert.Equal(t, 200, res.Code)

	// })

	t.Run("success-case", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))
		// fmt.Println(jwtToken)

		productController := NewProductControllers(mockProductRepository{})
		productController.GetProductByID(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)

		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("2")

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.GetProductByID(context)
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Product([]entities.Product(nil)), response.Data)
		assert.Equal(t, "Not Found", response.Message)
	})
}

func TestGetStockHistoryProduct(t *testing.T) {
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
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products/stock/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockProductRepository{})
		productController.GetHistoryStockProduct(context)

		response := GetStockProductResponseFormat{}

		var response1 interface{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		json.Unmarshal([]byte(res.Body.Bytes()), &response1)

		temp := response1.(map[string]interface{})

		// data := temp["data"].(map[string]interface{})
		data := temp["data"].([]interface{})

		// fmt.Println(data[0])

		item := data[0].(map[string]interface{})

		qty := item["Qty"]

		// fmt.Println(item1)

		// productData := data["rows"].([]interface{})

		// item := productData[0].(map[string]interface{})

		// fmt.Println("qty", item["Qty"])

		// qty := response.Data[0].Qty

		assert.Equal(t, float64(1), qty)
		assert.Equal(t, "Successful Operation", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products/stock/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.GetHistoryStockProduct(context)

		response := GetStockProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Stock(nil), response.Data)
		assert.Equal(t, "Not Found", response.Message)
	})
}

func TestUpdateStockProduct(t *testing.T) {
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

		requestBody, _ := json.Marshal(entities.Stock{
			Product_id: 1,
			Qty:        1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products/stock/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockProductRepository{})
		productController.UpdateStockProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// data := response.Data

		// stock := data[0].Stock

		// assert.Equal(t, 2, stock)
		assert.Equal(t, "Successful Operation", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(entities.Stock{
			Product_id: 2,
			Qty:        1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products/stocks")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.UpdateStockProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Product([]entities.Product(nil)), response.Data)
	})
}

func TestCreateProduct(t *testing.T) {
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

		requestBody, _ := json.Marshal(entities.Product{
			Name:        "Product Alpha",
			Stock:       1,
			Price:       10000,
			Category_id: 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockProductRepository{})
		productController.CreateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// data := response.Data

		// name := data[0].Name

		// assert.Equal(t, "Product Alpha", name)
		assert.Equal(t, "Successful Operation", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"Name": "Product Alpha",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/products")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.CreateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Internal Server Error", response.Message)
	})
}

func TestUpdateProduct(t *testing.T) {
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
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Product Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockProductRepository{})
		productController.UpdateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// data := response.Data

		// name := data[0].Name

		// assert.Equal(t, "Product Alpha new", name)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name": "Product Alpha new",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.UpdateProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, []entities.Product([]entities.Product(nil)), response.Data)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
}

func TestDeleteProduct(t *testing.T) {
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
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("1")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockProductRepository{})
		productController.DeleteProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("id")
		context.SetParamValues("10")
		// context.Request().Header.Set(echo.HeaderAuthorization, fmt.Sprint("Bearer", jwtToken))

		productController := NewProductControllers(mockFalseProductRepository{})
		productController.DeleteProduct(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Internal Server Error", response.Message)
	})
}

func TestExportPDF(t *testing.T) {
	t.Run("success-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/export")

		productController := NewProductControllers(mockProductRepository{})
		productController.ExportPDF(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		// fmt.Println(response.Message)

		assert.Equal(t, "success export pdf", "success export pdf")
	})

	t.Run("error-case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/export")

		productController := NewProductControllers(mockProductRepository{})
		productController.ExportPDF(context)

		response := GetProductResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		fmt.Println(response.Message)

		assert.Equal(t, "can't export pdf", "can't export pdf")
	})
}

type mockProductRepository struct{}

func (m mockProductRepository) GetAllProduct() ([]entities.Product, error) {
	return []entities.Product{
		{ID: 1, Name: "Product Alpha"},
	}, nil
}

func (m mockProductRepository) GetProductByID(product_id int) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha"}, nil
}

func (m mockProductRepository) GetHistoryStockProduct(product_id int) ([]entities.Stock, error) {
	return []entities.Stock{
		{Product_id: 1, Qty: 1}}, nil
}

func (m mockProductRepository) CreateProduct(entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha"}, nil
}

func (m mockProductRepository) UpdateProduct(product_id int, product entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha new"}, nil
}

func (m mockProductRepository) UpdateStockProduct(product_id, qty int) (entities.Product, error) {
	return entities.Product{
		ID: 1, Stock: 2}, nil
}

func (m mockProductRepository) DeleteProduct(product_id int) (entities.Product, error) {
	return entities.Product{
		ID: 1, Name: "Product Alpha"}, nil
}

func (m mockProductRepository) ProductPagination(Pagination pagination.ProductPagination) (interface{}, int, error) {
	return pagination.ProductPagination{
		Limit: 0, Page: 0, TotalRows: 5, FirstPage: "/products?limit=0&page=0", LastPage: "/products?limit=0&page=0",
		PreviousPage: "", NextPage: "/products?limit=0&page=1", FromRow: 1, ToRow: 5,
		Rows: []entities.Product{
			{ID: 1, Name: "Product Alpha", Price: 1000, Stock: 1, Category_id: 1},
			{ID: 2, Name: "Product Beta", Price: 2000, Stock: 2, Category_id: 2},
			{ID: 3, Name: "Product Cherry", Price: 3000, Stock: 3, Category_id: 3},
			{ID: 4, Name: "Product Delta", Price: 4000, Stock: 4, Category_id: 4},
			{ID: 5, Name: "Product Echo", Price: 5000, Stock: 5, Category_id: 5},
		},
	}, 5, nil
}

type mockFalseProductRepository struct{}

func (m mockFalseProductRepository) GetAllProduct() ([]entities.Product, error) {
	return nil, errors.New("no data")
}

func (m mockFalseProductRepository) GetProductByID(id int) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("can't get product")
}

func (m mockFalseProductRepository) GetHistoryStockProduct(product_id int) ([]entities.Stock, error) {
	return []entities.Stock{
		{Product_id: 0, Qty: 0}}, nil
}

func (m mockFalseProductRepository) CreateProduct(entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("error create product")
}

func (m mockFalseProductRepository) UpdateProduct(product_id int, product entities.Product) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("error update product")
}

func (m mockFalseProductRepository) UpdateStockProduct(product_id, qty int) (entities.Product, error) {
	return entities.Product{
		ID: 0, Stock: 0}, errors.New("can't update stock product")
}

func (m mockFalseProductRepository) DeleteProduct(product_id int) (entities.Product, error) {
	return entities.Product{
		ID: 0, Name: ""}, errors.New("error delete product")
}

func (m mockFalseProductRepository) ProductPagination(Pagination pagination.ProductPagination) (interface{}, int, error) {
	return pagination.ProductPagination{Limit: 0, Page: 0}, 0, errors.New("can't get products")
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
