package carts

import (
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Stock{})

	db.Migrator().DropTable(&entities.Cart{})
	db.Migrator().DropTable(&entities.Detail_cart{})
	db.Migrator().DropTable(&entities.User{})

	db.AutoMigrate(entities.Category{})
	db.AutoMigrate(entities.Product{})
	db.AutoMigrate(entities.Stock{})

	db.AutoMigrate(entities.Cart{})
	db.AutoMigrate(entities.Detail_cart{})
	db.AutoMigrate(entities.User{})

	cartRepo := NewCartsRepo(db)

	t.Run("Insert Cart", func(t *testing.T) {

		var newCart entities.Cart
		newCart.DateCheckout = time.Now()
		newCart.Total_Product = 0
		newCart.Total_price = 0

		res, err := cartRepo.Insert(newCart)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

}

func TestCartRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Cart{})
	db.AutoMigrate(&entities.Cart{})

	cartRepo := NewCartsRepo(db)

	t.Run("insert cart", func(t *testing.T) {
		var mockCart entities.Cart
		mockCart.ID = 1
		mockCart.Total_price = 0
		mockCart.Detail_cart_ID = []entities.Detail_cart{}

		res, err := cartRepo.Insert(mockCart)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("select * from cart", func(t *testing.T) {
		res, err := cartRepo.Gets()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
