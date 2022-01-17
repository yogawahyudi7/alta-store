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
func (tr *TransactionsRepository) Insert(newTransactions entities.Transaction) (entities.Transaction, error) {
	tr.db.Save(&newTransactions)
	return newTransactions, nil
}
