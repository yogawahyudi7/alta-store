package transactions

import "project-e-commerces/entities"

type TransactionInterface interface {
	Gets(userID uint) ([]entities.Transaction, error)
	InsertT(newTransactions entities.Transaction) (entities.Transaction, error)
	InsertDT(newDetailTransactions entities.Detail_transaction) (entities.Detail_transaction, error)
	Update(updateStatus string, trID uint) (entities.Transaction, error)
	Delete(trID, userID uint) (entities.Transaction, error)
	GetsPaymentUrl(userID uint, totalPrice, totalQty int, invoiceID string) (string, error)
}
