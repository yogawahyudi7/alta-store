package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type GetUserResponseFormat struct {
	Data    []entities.User `json:"data"`
	Message string          `json:"message"`
}

type GetUserResponse struct {
	Data    entities.User `json:"data"`
	Message string        `json:"message"`
	Token   string        `json:"token"`
}
type GetUserResponseLogin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var jwtToken string

func TestRegister(t *testing.T) {
	t.Run("register", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/register")

		userController := NewUserController(mockUserRepository{})
		userController.Register(context)

		var response entities.User

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "mock1", response.Name)
	})
}
func TestLogin(t *testing.T) {
	t.Run("Login", func(t *testing.T) {
		e := echo.New()

		resBody, _ := json.Marshal(map[string]interface{}{"email": "yoga1", "password": "yoga1"})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(resBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")
		fmt.Println(req)

		userController := NewUserController(mockUserRepository{})
		userController.Login(context)

		var response GetUserResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		fmt.Println(response.Message)
		// jwtToken = response.Data.(string)
		assert.Equal(t, "ALO", response)
	})
}

func TestGet(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		// context.Request().Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiYXV0aG9yaXplZCI6dHJ1ZSwiZXhwIjoxNjQyODY1NDQ1LCJ1c2VySWQiOjN9.KpzNONOjnUA8-6dXJv6yPN43iyNTX4WLOBr6nxPx24A")
		context.SetPath("/users/profile/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		userController := NewUserController(mockUserRepository{})
		userController.Get(context)

		var response entities.User

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		fmt.Println(response)
		assert.Equal(t, "mock1", response.Name)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users/profile")

		userController := NewUserController(mockUserRepository{})
		userController.Register(context)

		var response entities.User

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "mock1", response.Name)
	})
}

//MOCK
type mockUserRepository struct{}

func (m mockUserRepository) Register(entities.User) (entities.User, error) {

	return entities.User{Name: "mock1", Password: "mock1", Email: "mock1"}, nil

}

func (m mockUserRepository) Login(email string) (entities.User, error) {
	return entities.User{Email: "yoga1", Password: "yoga1"}, nil
}

// func (m mockUserRepository) FindUserByEmail(email string) (entities.User, error) {
// 	return entities.User{Email: "yoga1", Password: "yoga1"}, nil
// }
func (m mockUserRepository) Get(user int) (entities.User, error) {
	return entities.User{ID: 3, Name: "mock1", Password: "mock1", Email: "mock1", Role: "admin"}, nil
}
func (m mockUserRepository) Update(user entities.User) (entities.User, error) {
	return entities.User{Name: user.Name, Email: user.Email, Password: user.Password}, nil
}

func (m mockUserRepository) Delete(userId int) error {
	return nil
}
