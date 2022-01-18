package main

import (
	"fmt"
	"project-e-commerces/configs"
	controllers "project-e-commerces/delivery/controllers/users"
	"project-e-commerces/delivery/routes"
	repository "project-e-commerces/repository/users"
	"project-e-commerces/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	config := configs.GetConfig()
	db := utils.InitDB(config)
	repoUser := repository.NewRepository(db)
	controllerUser := controllers.NewUserController(repoUser)

	e := echo.New()
	routes.RegisterUserPath(e, *controllerUser)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
