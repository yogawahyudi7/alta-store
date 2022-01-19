package carts

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCarts(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Detail_cart{})
	db.Migrator().DropTable(&entities.Cart{})

	db.AutoMigrate(&entities.Cart{})
	db.AutoMigrate(&entities.Detail_cart{})

	e := echo.New()

	t.Run("PUT /carts/additem/:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
			Qty:       1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/additem/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cartCon := NewCartsControllers(mockCartRepository{})
		cartCon.PutItemIntoDetail_CartCtrl()(context)

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})

	t.Run("DEL /carts/delitem:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/delitem:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cartCon := NewCartsControllers(mockCartRepository{})
		cartCon.DeleteItemFromDetail_CartCtrl()(context)

		responses := DelItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})
}

func TestFalseCarts(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Detail_cart{})
	db.Migrator().DropTable(&entities.Cart{})

	db.AutoMigrate(&entities.Cart{})
	db.AutoMigrate(&entities.Detail_cart{})

	e := echo.New()

	t.Run("PUT /carts/additem/:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
			Qty:       1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/additem/:id")
		context.SetParamNames("id")
		context.SetParamValues("a")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		cartCon.PutItemIntoDetail_CartCtrl()(context)

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Not Found", responses.Message)
		assert.Equal(t, 404, res.Code)

	})
	t.Run("PUT /carts/additem/:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/additem/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		cartCon.PutItemIntoDetail_CartCtrl()(context)

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Internal Server Error", responses.Message)
		assert.Equal(t, 500, res.Code)

	})
	t.Run("PUT /carts/additem/:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"product_id": "1",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/additem/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		cartCon.PutItemIntoDetail_CartCtrl()(context)

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Bad Request", responses.Message)
		assert.Equal(t, 400, res.Code)

	})
	t.Run("DEL /carts/delitem:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/delitem:id")
		context.SetParamNames("id")
		context.SetParamValues("a")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		cartCon.DeleteItemFromDetail_CartCtrl()(context)

		responses := DelItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Not Found", responses.Message)
		assert.Equal(t, 404, res.Code)

	})
	t.Run("DEL /carts/delitem:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/delitem:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		cartCon.DeleteItemFromDetail_CartCtrl()(context)

		responses := DelItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Internal Server Error", responses.Message)
		assert.Equal(t, 500, res.Code)

	})
	t.Run("DEL /carts/delitem:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"product_id": "1",
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/delitem:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		cartCon.DeleteItemFromDetail_CartCtrl()(context)

		responses := DelItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Bad Request", responses.Message)
		assert.Equal(t, 400, res.Code)

	})
}

type mockCartRepository struct{}

func (mcr mockCartRepository) Gets() ([]entities.Cart, error) {
	return []entities.Cart{}, nil
}

func (mcr mockCartRepository) Insert(newCart entities.Cart) (entities.Cart, error) {
	return entities.Cart{Total_Product: 0, Total_price: 0}, nil
}

func (mcr mockCartRepository) InsertProduct(cartID uint, newItem entities.Detail_cart) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1, ProductID: 1, Qty: 1}, nil
}

func (mcr mockCartRepository) UpdateProduct(cartID, productID uint, qty int) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1, ProductID: 1, Qty: 2}, nil
}

func (mcr mockCartRepository) DeleteProduct(cartID, productID uint) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1}, nil
}

type mockFalseCartRepository struct{}

func (mcr mockFalseCartRepository) Gets() ([]entities.Cart, error) {
	return []entities.Cart{}, errors.New("Bad Request")
}

func (mcr mockFalseCartRepository) Insert(newCart entities.Cart) (entities.Cart, error) {
	return entities.Cart{Total_Product: 0, Total_price: 0}, errors.New("Bad Request")
}

func (mcr mockFalseCartRepository) InsertProduct(cartID uint, newItem entities.Detail_cart) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1, ProductID: 1, Qty: 1}, errors.New("Bad Request")
}

func (mcr mockFalseCartRepository) UpdateProduct(cartID, productID uint, qty int) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1, ProductID: 1, Qty: 2}, errors.New("Bad Request")
}

func (mcr mockFalseCartRepository) DeleteProduct(cartID, productID uint) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1}, errors.New("Bad Request")
}
