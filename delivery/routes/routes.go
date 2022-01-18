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
	auth.GET("/users/profile", uc.Get)
	auth.DELETE("/users/delete", uc.Delete)
	auth.PUT("/users/update", uc.Update)

}

func RegisterPath(e *echo.Echo, crCtrl *carts.CartsController, tsCtrl *transactions.TransactionsController) {

}
