package transactions

import "project-e-commerces/entities"

type TransactionInterface interface {
	Gets() ([]entities.Transaction, error)
	Get(userID int) (entities.Transaction, error)
	InsertT(newTransactions entities.Transaction) (entities.Transaction, error)
	InsertDT(newDetailTransactions entities.Detail_transaction) (entities.Detail_transaction, error)
	Update(updateTransactions entities.Transaction, trID int) (entities.Transaction, error)
	Delete(trID, userID int) (entities.Transaction, error)
	GetPaymentURL(trs entities.Transaction, userID uint) (string, error)
}
