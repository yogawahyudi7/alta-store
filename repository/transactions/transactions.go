package transactions

import (
	"project-e-commerces/entities"
	"strconv"
	"time"

	"gorm.io/gorm"

	midtrans "github.com/midtrans/midtrans-go"
	snap "github.com/midtrans/midtrans-go/snap"
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
func (tr *TransactionsRepository) InsertT(newTransactions entities.Transaction) (entities.Transaction, error) {
	tr.db.Save(&newTransactions)
	return newTransactions, nil
}
func (tr *TransactionsRepository) InsertDT(newDetailTransactions entities.Detail_transaction) (entities.Detail_transaction, error) {
	tr.db.Save(&newDetailTransactions)
	return newDetailTransactions, nil
}

func (tr *TransactionsRepository) Update(updateTransactions entities.Transaction, trID int) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	tr.db.Where("id=?", trID).Find(&transaction)
	transaction.Status = updateTransactions.Status

	tr.db.Save(&transaction)
	return transaction, nil
}

func (tr *TransactionsRepository) Delete(trID, userID int) (entities.Transaction, error) {
	transaction := entities.Transaction{}
	tr.db.Find(&transaction, "id=? AND user_id=?", trID, userID)
	tr.db.Delete(&transaction)
	return transaction, nil
}

func (tr *TransactionsRepository) GetPaymentURL(transaction entities.Transaction, userID uint) (string, error) {
	midtrans.ServerKey = "SB-Mid-server-WBQoXNegZ5veTRfQsX3WOGFq"
	midtrans.ClientKey = "SB-Mid-client-lbfJ_9e_8nsyvWWS"
	midtrans.Environment = midtrans.Sandbox
	tanggal := time.Now()
	y, m, d := tanggal.Date()
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "INV-" + strconv.Itoa(y) + strconv.Itoa(int(m)) + strconv.Itoa(d) + "/" + strconv.Itoa(int(userID)),
			GrossAmt: int64(transaction.Total_price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapTokenResp, _ := snap.CreateTransaction(req)
	return snapTokenResp.RedirectURL, nil
}

func (tr *TransactionsRepository) GetsPaymentUrl(userID uint, totalPrice, totalQty int) (string, error) {
	midtrans.ServerKey = "SB-Mid-server-WBQoXNegZ5veTRfQsX3WOGFq"
	midtrans.ClientKey = "SB-Mid-client-lbfJ_9e_8nsyvWWS"
	midtrans.Environment = midtrans.Sandbox
	tanggal := time.Now()
	y, m, d := tanggal.Date()
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "INV-" + strconv.Itoa(y) + strconv.Itoa(int(m)) + strconv.Itoa(d) + "/c/" + strconv.Itoa(int(userID)),
			GrossAmt: int64(totalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapTokenResp, _ := snap.CreateTransaction(req)
	return snapTokenResp.RedirectURL, nil
}
