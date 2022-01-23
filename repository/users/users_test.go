package users

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	userRep := NewRepository(db)

	var dummyUser entities.User
	dummyUser.Name = "test1"
	dummyUser.Email = "test1"
	dummyUser.Password = "test1"
	dummyUser.CartID = 1

	// e := echo.New()

	t.Run("RegisterUser", func(t *testing.T) {
		res, err := userRep.Register(dummyUser)
		assert.Nil(t, err)
		assert.Equal(t, "test1", res.Name)

	})
	t.Run("LoginUser", func(t *testing.T) {
		res, err := userRep.Login("test1")
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("GetUser", func(t *testing.T) {
		res, err := userRep.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		var mockUpdateUser entities.User
		mockUpdateUser.Name = "test1"
		mockUpdateUser.Password = "test1"
		mockUpdateUser.CartID = 1

		res, err := userRep.Update(mockUpdateUser)
		fmt.Println(res)
		assert.Nil(t, err)
		assert.Equal(t, dummyUser.Email, res.Name)
	})
	t.Run("Delete User", func(t *testing.T) {
		err := userRep.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})
}
