package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/entities"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

		var response GetUserResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "mock1", response.Data.Name)
	})
}
func TestLogin(t *testing.T) {
	t.Run("Login", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/login")

		userController := NewUserController(mockUserRepository{})
		userController.Register(context)

		var response GetUserResponse

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		jwtToken = response.Token
		assert.Equal(t, "success", response.Message)
	})
}

//MOCK
type mockUserRepository struct{}

func (m mockUserRepository) Register(entities.User) (entities.User, error) {

	return entities.User{Name: "mock1", Password: "mock1", Email: "mock1"}, nil

}
func (m mockUserRepository) Login(email, password string) (entities.User, error) {
	return entities.User{Email: email, Password: password}, nil
}

func (m mockUserRepository) Get(userId int) (entities.User, error) {
	return entities.User{Model: gorm.Model{ID: 1}, Name: "mock1", Password: "mock1", Email: "mock1"}, nil
}
func (m mockUserRepository) Update(user entities.User) (entities.User, error) {
	return entities.User{Name: user.Name, Email: user.Email, Password: user.Password}, nil
}

func (m mockUserRepository) Delete(userId int) error {
	return nil
}
