package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	StatusReject   string = "rejected"
	StatusApproved string = "approved"
)

type TransactionRepository interface {
	SaveTransaction(transaction Transaction, creditCard CreditCard) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if t.Amount+creditCard.Balance > creditCard.Limit {
		t.Status = StatusReject
	} else {
		t.Status = StatusApproved
		creditCard.Balance = creditCard.Balance + t.Amount
	}
}
