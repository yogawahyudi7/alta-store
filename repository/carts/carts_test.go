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
	t.Run("insert some product into detail_cart", func(t *testing.T) {
		var mockDetailCart entities.Detail_cart
		mockDetailCart.ProductID = 1
		mockDetailCart.Qty = 1

		res, err := cartRepo.InsertProduct(1, mockDetailCart)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("update qty some product into detail_cart", func(t *testing.T) {
		res, err := cartRepo.UpdateProduct(1, 1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("delete some product into detail_cart", func(t *testing.T) {
		res, err := cartRepo.DeleteProduct(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
