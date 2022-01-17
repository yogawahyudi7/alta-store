package carts

import (
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Cart{})
	db.AutoMigrate(&entities.Cart{})

	cartRepo := NewCartsRepo(db)

	t.Run("insert cart", func(t *testing.T) {
		var mockCart entities.Cart
		mockCart.ID = 1
		mockCart.Product_id = 1
		mockCart.Product_qty = 2
		mockCart.Total_price = 20000

		res, err := cartRepo.Insert(mockCart)
		assert.Nil(t, err)
		assert.Equal(t, mockCart.Product_id, res.Product_id)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("select * from cart", func(t *testing.T) {
		res, err := cartRepo.Gets()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
