package utils

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {
	var connectionString string

	connectionString =
		fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			config.Database.Username,
			config.Database.Password,
			config.Database.Address,
			config.Database.DB_Port,
			config.Database.Name,
		)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {

	db.AutoMigrate(entities.Category{})
	db.AutoMigrate(entities.Product{})
	db.AutoMigrate(entities.Stock{})

	db.AutoMigrate(entities.Cart{})
	db.AutoMigrate(entities.Detail_cart{})
	db.AutoMigrate(entities.User{})

	db.AutoMigrate(entities.Payment{})
	db.AutoMigrate(entities.Transaction{})
	db.AutoMigrate(entities.Detail_transaction{})

}
