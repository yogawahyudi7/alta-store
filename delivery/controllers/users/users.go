package controllers

import (
	"net/http"
	"project-e-commerces/entities"
	repository "project-e-commerces/repository/users"

	"project-e-commerces/delivery/common"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo repository.UserInterface
}

func NewUserController(user repository.UserInterface) *UserController {
	return &UserController{user}
}

//================
//REGISTER & LOGIN
//================
func (uc UserController) Register(c echo.Context) error {
	var user entities.User
	c.Bind(&user)
	hash, _ := Hashpwd(user.Password)
	user.Password = hash
	res, err := uc.Repo.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "email already exist"))
	}
	return c.JSON(http.StatusOK, common.SuccessResponse(res))
}

func (uc UserController) Login(c echo.Context) error {
	var login entities.User
	c.Bind(&login)
	user, err := uc.Repo.Login(login.Email, login.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not registered"))
	}

	hash, err := Checkpwd(user.Password, login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "wrong password"))
	}
	var token string
	if hash {
		token, _ = CreateToken(int(user.ID))
	}
	return c.JSON(http.StatusOK, common.SuccessResponse(token))
}

//================
//RUD USER
//================
func (uc UserController) Get(c echo.Context) error {
	userId := ExtractTokenUserId(c)
	user, err := uc.Repo.Get(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not found"))
	}
	return c.JSON(http.StatusOK, common.SuccessResponse(user))
}

func (uc UserController) Delete(c echo.Context) error {
	userId := ExtractTokenUserId(c)
	err := uc.Repo.Delete(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, err.Error()))
	}
	return c.JSON(http.StatusOK, common.SuccessResponse(err))
}

func (uc UserController) Update(c echo.Context) error {
	userId := ExtractTokenUserId(c)
	user, err := uc.Repo.Get(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not found"))
	}
	var tmpUser entities.User
	c.Bind((&tmpUser))
	user.Name = tmpUser.Name
	user.Email = tmpUser.Email
	hash, _ := Hashpwd(user.Password)
	user.Password = hash
	userRes, err := uc.Repo.Update(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, common.SuccessResponse(userRes))
}
