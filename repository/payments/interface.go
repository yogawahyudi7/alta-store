package payments

import "project-e-commerces/entities"

type PaymentInterface interface {
	Gets() ([]entities.Payment, error)
	Get(pyID int) (entities.Payment, error)
	Insert(newPayments entities.Payment) (entities.Payment, error)
	Update(updatePayments entities.Payment, pyID int) (entities.Payment, error)
	Delete(pyID int) (entities.Payment, error)
}
