package transactions

import "project-e-commerces/entities"

type TransactionInterface interface {
	Gets() ([]entities.Transaction, error)
	Insert(newTransactions entities.Transaction) (entities.Transaction, error)
}
