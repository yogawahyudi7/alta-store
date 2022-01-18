package transactions

import "project-e-commerces/entities"

type TransactionInterface interface {
	Gets() ([]entities.Transaction, error)
	Get(userID int) (entities.Transaction, error)
	Insert(newTransactions entities.Transaction) (entities.Transaction, error)
	Update(updateTransactions entities.Transaction, trID int) (entities.Transaction, error)
	Delete(trID, userID int) (entities.Transaction, error)
}
