package transactions

import "project-e-commerces/repository/transactions"

type TransactionsController struct {
	Repo transactions.TransactionInterface
}

func NewTransactionsControllers(tsrep transactions.TransactionInterface) *TransactionsController {
	return &TransactionsController{Repo: tsrep}
}
