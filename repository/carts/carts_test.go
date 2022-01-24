package carts

import (
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/repository/categorys"
	"project-e-commerces/repository/products"
	"project-e-commerces/repository/users"
	"project-e-commerces/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	catRepo := categorys.NewCategoryRepo(db)
	var newCategory entities.Category
	newCategory.Name = "besi"
	catRepo.CreateCategory(newCategory)

	prodRepo := products.NewProductRepo(db)
	var newProduct entities.Product
	newProduct.Name = "Product 1"
	newProduct.Category_id = 1

	prodRepo.CreateProduct(newProduct)

	userRepo := users.NewUsersRepo(db)
	var newUser entities.User
	newUser.Name = "TestNameCart1"
	newUser.Email = "TestCart1@email.com"
	newUser.Password = "TestPassword1Cart"
	userRepo.Create(newUser)

	cartRepo := NewCartsRepo(db)

	var newCart entities.Cart
	newCart.ID = 1
	newCart.User_id = 1
	newCart.Total_Product = 0
	newCart.Total_price = 0

	cartRepo.Insert(newCart)

	var newDCart entities.Detail_cart
	newDCart.CartID = 1
	newDCart.ProductID = 1

	cartRepo.InsertProduct(newDCart)

	t.Run("insert cart", func(t *testing.T) {
		var mockCart entities.Cart
		mockCart.Total_price = 2
		mockCart.Total_Product = 2

		res, err := cartRepo.Insert(mockCart)
		assert.Nil(t, err)
		assert.Equal(t, 0, int(res.ID))
	})
	t.Run("select * from cart", func(t *testing.T) {
		res, err := cartRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("insert some product into detail_cart", func(t *testing.T) {
		var mockDetailCart entities.Detail_cart
		mockDetailCart.ProductID = 1
		mockDetailCart.Qty = 1

		res, err := cartRepo.InsertProduct(mockDetailCart)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("delete some product into detail_cart", func(t *testing.T) {
		res, err := cartRepo.DeleteProduct(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
