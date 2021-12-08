package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

//Constructor
func NewTransaction() *Transaction {
	return &Transaction{}
}

//Transaction methods
func (t *Transaction) isValid() error {
	if t.Amount > 1000 {
		return errors.New("you dont have limit for this transaction")
	}

	if t.Amount < 1 {
		return errors.New("the amount must be greater than 1")
	}
	return nil
}

//TODO: test coverage
func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}
