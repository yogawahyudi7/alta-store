package users

import (
	"net/http"
	"project-e-commerces/constants"
	"project-e-commerces/entities"
	repository "project-e-commerces/repository/users"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo repository.UserInterface
}

func NewUserController(user repository.UserInterface) *UserController {
	return &UserController{user}
}

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

	user, err := uc.Repo.GetLoginData(login.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not registered"))
	}

	hash, err := middlewares.Checkpwd(user.Password, login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "wrong password"))
	}

	var token string

	if hash {
		token, _ = middlewares.CreateToken(int(user.ID))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(token))
}

func (uc UserController) GetUser(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := uc.Repo.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not found"))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(user))
}

func (uc UserController) Delete(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	err := uc.Repo.Delete(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, err.Error()))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(err))
}

func (uc UserController) Update(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := uc.Repo.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not found"))
	}

	var tmpUser entities.User
	c.Bind((&tmpUser))
	user.Name = tmpUser.Name
	user.Email = tmpUser.Email

	hash, _ := middlewares.Hashpwd(user.Password)

	user.Password = hash

	userRes, err := uc.Repo.Update(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(userRes))
}

// Auth JWT
func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = int(userId)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET_KEY))
}

func ExtractTokenUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		return userId
	}
	return 0
}
