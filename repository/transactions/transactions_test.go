package transactions

import (
	"project-e-commerces/configs"
	"project-e-commerces/entities"
	"project-e-commerces/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Transaction{})
	db.Migrator().DropTable(&entities.Detail_transaction{})

	db.AutoMigrate(entities.Transaction{})
	db.AutoMigrate(entities.Detail_transaction{})

	transactionRepo := NewTransactionsRepo(db)

	t.Run("insert transaction", func(t *testing.T) {
		var mockTransaction entities.Transaction
		mockTransaction.Total_price = 10000
		mockTransaction.Total_qty = 1
		mockTransaction.User_id = 1
		mockTransaction.Status = "Pending"

		res, err := transactionRepo.InsertT(mockTransaction)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction.Total_price, res.Total_price)
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
