package routes

import (
	"project-e-commerces/delivery/controllers/carts"
	"project-e-commerces/delivery/controllers/transactions"

	"project-e-commerces/constants"
	controllers "project-e-commerces/delivery/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserPath(e *echo.Echo, uc controllers.UserController) {
	// e.Pre(middleware.RemoveTrailingSlash())
	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	//REGISTER & LOGIN
	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)

	//RUD USER
	e.GET("/users/profile", uc.Get, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.DELETE("/users/delete", uc.Delete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	e.PUT("/users/update", uc.Update, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

}

func RegisterPath(e *echo.Echo, crCtrl *carts.CartsController, tsCtrl *transactions.TransactionsController) {

}
