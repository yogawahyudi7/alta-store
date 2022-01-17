package main

import (
	"fmt"
	"log"
	"project-e-commerces/configs"
	"project-e-commerces/delivery/controllers/carts"
	"project-e-commerces/delivery/controllers/transactions"
	"project-e-commerces/delivery/routes"
	cartsRepo "project-e-commerces/repository/carts"
	transactionsRepo "project-e-commerces/repository/transactions"
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

	routes.RegisterPath(e, cartsCtrl, transactionsCtrl)

	address := fmt.Sprintf("localhost:%d", config.Port)
	log.Fatal(e.Start(address))
}
