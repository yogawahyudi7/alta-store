package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"project-e-commerces/constants"
	"project-e-commerces/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

type GetUserResponseFormat struct {
	Data    []entities.User `json:"data"`
	Message string          `json:"message"`
}

type UpdateUserResponseFormat struct {
	Message string          `json:"message"`
	Data    []entities.User `json:"data"`
}
type GetUserResponse struct {
	Data    entities.User `json:"data"`
	Message string        `json:"message"`
	Token   string        `json:"token"`
}

type GetLogin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
type GetUserResponseLogin struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func TestRegister(t *testing.T) {
	t.Run("registerUser", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(
			map[string]string{
				"name":     "mock1",
				"email":    "mock1",
				"password": "mock1",
			},
		)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users/register")

		userController := NewUserController(mockUserRepository{})
		userController.Register(context)

		var response LoginResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "success register", response.Message)
	})

	t.Run("falseregisterUser", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(
			map[string]string{
				"name":     "mock1",
				"email":    "mock1",
				"password": "mock1",
			},
		)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users/register")

		userController := NewUserController(falseMockUserRepository{})
		userController.Register(context)

		var response string

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "masukan input", response)
	})
}

var jwtToken string

func TestUser(t *testing.T) {
	t.Run("LoginUser", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"email":    "mock1",
			"password": "mock1",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
		res := httptest.NewRecorder()
		// req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := NewUserController(mockUserRepository{})
		userController.Login(context)
		// fmt.Println(res)

		// fmt.Println(jwtToken)
		response := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		jwtToken = response.Token
		// fmt.Println(response)
		assert.Equal(t, "success login", response.Message)
	})
	t.Run("GetUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/users/profile")

		userController := NewUserController(mockUserRepository{})
		middleware.JWT([]byte(constants.JWT_SECRET_KEY))(userController.Get())(context)

		// fmt.Println(res)

		var response LoginResponseFormat
		// jwtToken = response.Token
		// fmt.Println(jwtToken)
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "success get data", response.Message)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		e := echo.New()
		reqBody, _ := json.Marshal(map[string]string{
			"name":     "yogaUpdate",
			"email":    "yogaUpdate",
			"password": "yogaUpdate",
		})
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/update")

		UserController := NewUserController(mockUserRepository{})

		err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(UserController.Update())(context)
		if err != nil {
			log.Fatal(err)
			return
		}

		responses := UpdateUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "success update", responses.Message)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/delete")

		UserController := NewUserController(mockUserRepository{})

		err := middleware.JWT([]byte(constants.JWT_SECRET_KEY))(UserController.Delete())(context)
		if err != nil {
			log.Fatal(err)
			return
		}
		responses := UpdateUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)
		assert.Equal(t, "success delete", responses.Message)
	})
}

var jwtToken2 string

func TestUserFalse(t *testing.T) {
	t.Run("LoginUserFalse", func(t *testing.T) {
		e := echo.New()

		bodyReq, _ := json.Marshal(map[string]int{
			"email":    1,
			"password": 1,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(bodyReq))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := NewUserController(mockUserRepository{})
		userController.Login(context)

		response := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		jwtToken2 = response.Token
		assert.Equal(t, "kesalahan input", response.Message)

	})
	t.Run("LoginUserFalse2", func(t *testing.T) {
		e := echo.New()

		bodyReq, _ := json.Marshal(map[string]string{
			"email":    "tidakAda",
			"password": "tidakAda",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(bodyReq))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		userController := NewUserController(falseMockUserRepository{})
		userController.Login(context)

		response := LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		jwtToken2 = response.Token
		assert.Equal(t, "email tidak ditemukan", response.Message)

	})
}

//MOCK
type mockUserRepository struct{}

func (m mockUserRepository) Register(user entities.User) (entities.User, error) {
	return entities.User{Name: "mock1", Password: "mock1", Email: "mock1"}, nil
}
func (m mockUserRepository) Login(email string) (entities.User, error) {
	hash, _ := Hashpwd("mock1")
	return entities.User{Email: "mock1", Password: hash}, nil
}
func (m mockUserRepository) Get(user int) (entities.User, error) {

	return entities.User{Name: "mock1", Password: "mock1", Email: "mock1", Role: "admin"}, nil
}
func (m mockUserRepository) Update(user entities.User) (entities.User, error) {
	hash, _ := Hashpwd("mock1")
	return entities.User{Email: "mock1", Password: hash}, nil
}
func (m mockUserRepository) Delete(userId int) error {
	return nil
}

type falseMockUserRepository struct{}

func (m falseMockUserRepository) Register(user entities.User) (entities.User, error) {
	return entities.User{}, errors.New("salah input")
}
func (m falseMockUserRepository) Login(email string) (entities.User, error) {
	// hash, _ := Hashpwd("mock1")
	return entities.User{}, errors.New("tidak temu iemail")
}
func (m falseMockUserRepository) Get(user int) (entities.User, error) {
	return entities.User{Name: "mock1", Password: "mock1", Email: "mock1"}, nil
}
func (m falseMockUserRepository) Update(user entities.User) (entities.User, error) {
	return entities.User{Name: user.Name, Email: user.Email, Password: user.Password}, nil
}
func (m falseMockUserRepository) Delete(userId int) error {
	return nil
}
