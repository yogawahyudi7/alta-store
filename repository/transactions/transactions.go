package transactions

import (
	"project-e-commerces/entities"

	"gorm.io/gorm"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepo(db *gorm.DB) *TransactionsRepository {
	return &TransactionsRepository{db: db}
}
func (tr *TransactionsRepository) Gets() ([]entities.Transaction, error) {
	transactions := []entities.Transaction{}
	tr.db.Find(&transactions)
	return transactions, nil
}
func (tr *TransactionsRepository) Get(trID int) (entities.Transaction, error) {
	transactions := entities.Transaction{}
	tr.db.Find(&transactions)
	return transactions, nil
}
func (tr *TransactionsRepository) Insert(newTransactions entities.Transaction) (entities.Transaction, error) {
	tr.db.Save(&newTransactions)
	return newTransactions, nil
}
func (tr *TransactionsRepository) Update(updateTransactions entities.Transaction, trID int) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	tr.db.Find(&updateTransactions, "id=?", trID)

	transaction.Status = updateTransactions.Status
	tr.db.Save(&transaction)

	return updateTransactions, nil
}
func (tr *TransactionsRepository) Delete(trID, userID int) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	tr.db.Find(&transaction, "id=? AND user_id=?", trID, userID)
	tr.db.Delete(&transaction)
	return transaction, nil
}
