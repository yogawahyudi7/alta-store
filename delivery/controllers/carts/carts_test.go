package carts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/configs"
	"project-e-commerces/constants"
	"project-e-commerces/delivery/controllers/users"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCarts(t *testing.T) {

	e := echo.New()
	jwtToken := ""
	t.Run("POST /users/login", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "Test@email.com",
			"password": "TestPassword1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCon := users.NewUsersControllers(mockUserRepository{})
		userCon.Login()(context)

		responses := users.LoginUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})

	t.Run("GET /carts", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts")

		cartCon := NewCartsControllers(mockCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)
	})

	t.Run("PUT /carts/additem", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
			Qty:       1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/additem")

		cartCon := NewCartsControllers(mockCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.PutItemIntoDetail_CartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})

	t.Run("DEL /carts/delitem", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/delitem")

		cartCon := NewCartsControllers(mockCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.DeleteItemFromDetail_CartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

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

	jwtToken := ""
	t.Run("POST /users/login", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "Test@email.com",
			"password": "TestPassword1",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userCon := users.NewUsersControllers(mockUserRepository{})
		userCon.Login()(context)

		responses := users.LoginUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, "Successful Operation", responses.Message)
		assert.Equal(t, 200, res.Code)

	})

	t.Run("GET /carts", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.Gets())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Internal Server Error", responses.Message)
		assert.Equal(t, 500, res.Code)
	})

	t.Run("PUT /carts/additem", func(t *testing.T) {
		reqBody, _ := json.Marshal(entities.Detail_cart{
			ProductID: 1,
			Qty:       1,
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/additem")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.PutItemIntoDetail_CartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := AddItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Internal Server Error", responses.Message)
		assert.Equal(t, 500, res.Code)

	})
	t.Run("PUT /carts/additem", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"product_id": "a",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/additem")

		cartCon := NewCartsControllers(mockCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.PutItemIntoDetail_CartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

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

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/delitem")

		cartCon := NewCartsControllers(mockFalseCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.DeleteItemFromDetail_CartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := DelItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Internal Server Error", responses.Message)
		assert.Equal(t, 500, res.Code)

	})
	t.Run("DEL /carts/delitem:id", func(t *testing.T) {
		reqBody, _ := json.Marshal(map[string]interface{}{
			"product_id": "b",
		})

		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/carts/delitem")

		cartCon := NewCartsControllers(mockCartRepository{})
		if err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(cartCon.DeleteItemFromDetail_CartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		responses := DelItemIntoDetail_CartResponsesFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		assert.Equal(t, "Bad Request", responses.Message)
		assert.Equal(t, 400, res.Code)

	})

}

type mockUserRepository struct{}

func (mur mockUserRepository) Login(name, password string) (entities.User, error) {
	return entities.User{ID: 1, Email: "Test@email.com", Password: "TestPassword1"}, nil
}

func (mur mockUserRepository) GetAll() ([]entities.User, error) {
	return []entities.User{
		{Name: "TestName1", Password: "TestPassword1"},
	}, nil
}
func (mur mockUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
}
func (mur mockUserRepository) Create(newUser entities.User) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
}
func (mur mockUserRepository) Update(updateUser entities.User, userId int) (entities.User, error) {
	return entities.User{Name: "TestName1", Password: "TestPassword1"}, nil
}
func (mur mockUserRepository) Delete(userId int) (entities.User, error) {
	return entities.User{ID: 1, Name: "TestName1", Password: "TestPassword1"}, nil
}

type mockCartRepository struct{}

func (mcr mockCartRepository) Get(cartID uint) ([]entities.Detail_cart, error) {
	return []entities.Detail_cart{{ID: 1}}, nil
}

func (mcr mockCartRepository) Insert(newCart entities.Cart) (entities.Cart, error) {
	return entities.Cart{Total_Product: 0, Total_price: 0}, nil
}

func (mcr mockCartRepository) InsertProduct(newItem entities.Detail_cart) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1, ProductID: 1, Qty: 1}, nil
}

func (mcr mockCartRepository) DeleteProduct(cartID, productID uint) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1}, nil
}

type mockFalseCartRepository struct{}

func (mcr mockFalseCartRepository) Get(cartID uint) ([]entities.Detail_cart, error) {
	return []entities.Detail_cart{}, errors.New("Bad Request")
}

func (mcr mockFalseCartRepository) Insert(newCart entities.Cart) (entities.Cart, error) {
	return entities.Cart{Total_Product: 0, Total_price: 0}, errors.New("Bad Request")
}

func (mcr mockFalseCartRepository) InsertProduct(newItem entities.Detail_cart) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1, ProductID: 1, Qty: 1}, errors.New("Bad Request")
}

func (mcr mockFalseCartRepository) DeleteProduct(cartID, productID uint) (entities.Detail_cart, error) {
	return entities.Detail_cart{ID: 1, CartID: 1}, errors.New("Bad Request")
}
