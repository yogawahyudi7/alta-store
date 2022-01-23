package controllers

import (
	"net/http"
	"project-e-commerces/entities"
	repository "project-e-commerces/repository/users"

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

	// var hash string
	hash, err := Hashpwd(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "input password salah")
	}
	user.Password = hash

	res, err := uc.Repo.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "masukan input")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success register",
		"data":    res,
	})
}

func (uc UserController) Login(c echo.Context) error {
	var login RegisterReqFormat
	if err := c.Bind(&login); err != nil {
		return c.JSON(http.StatusBadRequest, "kesalahan input")
	}

	user, err := uc.Repo.Login(login.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "email tidak ditemukan")
	}

	hash, err := Checkpwd(user.Password, login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "password salah")
	}

	var token string

	if hash {
		token, _ = CreateToken(int(user.ID), user.Role)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    user,
		"token":   token,
	})
}

//================
//RUD USER
//================
func (uc UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := ExtractTokenUserId(c)
		// userId, _ := strconv.Atoi(c.Param("id"))
		user, err := uc.Repo.Get(userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get data",
			"data":    user,
		})
	}
}

func (uc UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := ExtractTokenUserId(c)
		// userId, _ := strconv.Atoi(c.Param("id"))
		err := uc.Repo.Delete(userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete",
		})
	}
}

func (uc UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := ExtractTokenUserId(c)
		user, err := uc.Repo.Get(userId)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		var tmpUser entities.User
		c.Bind((&tmpUser))
		user.Name = tmpUser.Name
		user.Email = tmpUser.Email
		hash, _ := Hashpwd(tmpUser.Password)
		user.Password = hash
		userRes, err := uc.Repo.Update(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update",
			"data":    userRes,
		})
	}
}
