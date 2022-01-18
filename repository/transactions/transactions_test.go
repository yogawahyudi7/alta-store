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
	db.AutoMigrate(&entities.Transaction{})

	transactionRepo := NewTransactionsRepo(db)

	t.Run("insert transaction", func(t *testing.T) {
		var mockTransaction entities.Transaction
		mockTransaction.ID = 1
		mockTransaction.Total = 1
		mockTransaction.Total_price = 10000
		mockTransaction.Total_qty = 1

		res, err := transactionRepo.Insert(mockTransaction)
		assert.Nil(t, err)
		assert.Equal(t, mockTransaction.Total, res.Total)
		assert.Equal(t, 1, int(res.ID))
	})
	t.Run("select * from transaction", func(t *testing.T) {
		res, err := transactionRepo.Gets()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

}
