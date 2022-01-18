package payments

import (
	"project-e-commerces/entities"

	"gorm.io/gorm"
)

type PaymentsRepository struct {
	db *gorm.DB
}

func NewPaymentsRepo(db *gorm.DB) *PaymentsRepository {
	return &PaymentsRepository{db: db}
}
func (py *PaymentsRepository) Gets() ([]entities.Payment, error) {
	payments := []entities.Payment{}
	py.db.Find(&payments)
	return payments, nil
}
func (py *PaymentsRepository) Get(pyID int) (entities.Payment, error) {
	payments := entities.Payment{}
	py.db.Find(&payments)
	return payments, nil
}
func (py *PaymentsRepository) Insert(newPayments entities.Payment) (entities.Payment, error) {
	py.db.Save(&newPayments)
	return newPayments, nil
}
func (py *PaymentsRepository) Update(updatePayments entities.Payment, pyID int) (entities.Payment, error) {
	payment := entities.Payment{}
	py.db.Find(&updatePayments, "id=?", pyID)
	payment.Link = updatePayments.Link
	py.db.Save(&payment)
	return updatePayments, nil
}
func (py *PaymentsRepository) Delete(pyID int) (entities.Payment, error) {
	payment := entities.Payment{}
	py.db.Find(&payment, "id=?", pyID)
	py.db.Delete(&payment)
	return payment, nil
}
