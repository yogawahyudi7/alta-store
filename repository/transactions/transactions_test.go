package transactions

import (
	"fmt"
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/repository/carts"
	"project-e-commerces/repository/payments"
	"project-e-commerces/repository/users"
	"project-e-commerces/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Stock{})

	db.Migrator().DropTable(&entities.Cart{})
	db.Migrator().DropTable(&entities.Detail_cart{})
	db.Migrator().DropTable(&entities.User{})

	db.Migrator().DropTable(&entities.Payment{})
	db.Migrator().DropTable(&entities.Transaction{})
	db.Migrator().DropTable(&entities.Detail_transaction{})

	db.AutoMigrate(entities.Category{})
	db.AutoMigrate(entities.Product{})
	db.AutoMigrate(entities.Stock{})

	db.AutoMigrate(entities.Cart{})
	db.AutoMigrate(entities.Detail_cart{})
	db.AutoMigrate(entities.User{})

	db.AutoMigrate(entities.Payment{})
	db.AutoMigrate(entities.Transaction{})
	db.AutoMigrate(entities.Detail_transaction{})

	var newCart entities.Cart
	newCart.Total_price = 0
	newCart.DateCheckout = time.Now()
	newCart.Detail_cart_ID = []entities.Detail_cart{}

	var newUser entities.User
	newUser.Name = "TestName1"
	newUser.Password = "11223344"
	newUser.Email = "Test@email.com"
	newUser.Role = "user"
	newUser.Cart_id = 1

	var newPayment entities.Payment
	newPayment.Payment_type = "payment1"
	newPayment.Link = "ovo"

	var err error

	cartRep := carts.NewCartsRepo(db)
	_, err = cartRep.Insert(newCart)
	if err != nil {
		fmt.Println(err)
	}

	userRep := users.NewRepository(db)
	_, err = userRep.Register(newUser)
	if err != nil {
		fmt.Println(err)
	}

	paymentRep := payments.NewPaymentsRepo(db)
	_, err = paymentRep.Insert(newPayment)
	if err != nil {
		fmt.Println(err)
	}

	transactionRepo := NewTransactionsRepo(db)

	t.Run("insert transaction", func(t *testing.T) {
		var mockTransaction entities.Transaction
		mockTransaction.Total = 1
		mockTransaction.Total_price = 10000
		mockTransaction.Total_qty = 1
		mockTransaction.User_id = 1
		mockTransaction.PaymentID = 1
		mockTransaction.Status = "Pending"

		res, err := transactionRepo.Insert(mockTransaction)
		fmt.Println("si res", res)
		assert.Nil(t, err)

		assert.Equal(t, mockTransaction.Total, res.Total)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("select * from transaction", func(t *testing.T) {
		res, err := transactionRepo.Gets()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
	t.Run("select 1 from transaction", func(t *testing.T) {
		res, err := transactionRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, uint(1))
	})
	t.Run("update 1 from transaction", func(t *testing.T) {
		var mockUpdateTransaction entities.Transaction
		mockUpdateTransaction.Status = "SETTLEMENT"
		res, err := transactionRepo.Update(mockUpdateTransaction, 1)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, uint(1))
	})
	t.Run("delete 1 from transaction", func(t *testing.T) {
		res, err := transactionRepo.Delete(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, uint(1))
	})

}
