package utils

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database", err)
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
}
