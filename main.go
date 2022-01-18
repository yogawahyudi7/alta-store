package main

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/delivery/controllers/carts"
	"project-e-commerces/delivery/controllers/transactions"
	controllers "project-e-commerces/delivery/controllers/users"
	"project-e-commerces/delivery/routes"
	cartsRepo "project-e-commerces/repository/carts"
	transactionsRepo "project-e-commerces/repository/transactions"
	repository "project-e-commerces/repository/users"
	"project-e-commerces/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	e := echo.New()

	cartsRepo := cartsRepo.NewCartsRepo(db)
	cartsCtrl := carts.NewCartsControllers(cartsRepo)

	transactionsRepo := transactionsRepo.NewTransactionsRepo(db)
	transactionsCtrl := transactions.NewTransactionsControllers(transactionsRepo)

	repoUser := repository.NewRepository(db)
	controllerUser := controllers.NewUserController(repoUser)

	routes.RegisterUserPath(e, *controllerUser)
	routes.RegisterPath(e, cartsCtrl, transactionsCtrl)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
